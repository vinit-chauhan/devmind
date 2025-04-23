package ollama

import (
	"context"
	"strings"

	"github.com/ollama/ollama/api"
	"github.com/vinit-chauhan/devmind/internal/agent/types"
	"github.com/vinit-chauhan/devmind/internal/consumer"
	"github.com/vinit-chauhan/devmind/internal/logger"
)

var (
	tknCh  = make(chan string, 16)
	doneCh = make(chan struct{}, 1)
	errCh  = make(chan error, 1)
)

func (b *OllamaBackend) Respond(ctx context.Context, prompt string) (types.Readable, error) {
	var parsed OllamaChatResponse
	full := strings.Builder{}
	defer func() {
		close(errCh)
	}()

	// Produce the response in a goroutine
	go Produce(ctx, b, prompt)

	// Print the response as it comes to stdout
	go consumer.Consume(ctx, tknCh, doneCh, &full)

	// Wait for the response to finish
	<-doneCh

	parsed.Response = api.Message{
		Role:    "assistant",
		Content: full.String(),
	}
	parsed.Done = true

	logger.Debug("Ollama response: " + full.String())

	return &parsed, nil
}

func Produce(ctx context.Context, b *OllamaBackend, prompt string) {
	defer func() {
		logger.Debug("Closing Ollama Producer")
		close(tknCh)
	}()

	callbackFunc := func(cr api.ChatResponse) error {
		for {
			select {
			case <-ctx.Done():
				logger.Debug("Context done in Ollama chat")
				return context.Canceled
			case tknCh <- cr.Message.Content:
				return nil
			}
		}

	}

	req := api.ChatRequest{
		Model: b.conf.Model,
		Messages: []api.Message{
			{Role: "system", Content: SystemPrompt},
			{Role: "user", Content: prompt},
		},
	}

	err := b.client.Chat(ctx, &req, callbackFunc)
	if err != nil {
		logger.Error("Error in Ollama chat: " + err.Error())
		errCh <- err
	}
}
