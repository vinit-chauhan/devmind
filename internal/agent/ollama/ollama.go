package ollama

import (
	"context"
	"strings"

	"github.com/ollama/ollama/api"
	"github.com/vinit-chauhan/devmind/internal/agent/types"
	"github.com/vinit-chauhan/devmind/internal/logger"
)

func (b *OllamaBackend) Respond(ctx context.Context, prompt string) (types.Readable, error) {

	req := api.ChatRequest{
		Model: b.conf.Model,
		Messages: []api.Message{
			{Role: "system", Content: SystemPrompt},
			{Role: "user", Content: prompt},
		},
	}

	full := strings.Builder{}
	tknCh := make(chan string, 16)
	doneCh := make(chan struct{}, 1)
	errCh := make(chan error, 1)
	defer close(errCh)
	var parsed OllamaChatResponse

	go Produce(ctx, b, &req, tknCh, doneCh, errCh)

	go Consume(ctx, tknCh, doneCh, &full)

	<-doneCh

	parsed.Response = api.Message{
		Role:    "assistant",
		Content: full.String(),
	}
	parsed.Done = true

	logger.Debug("Ollama response: " + full.String())

	return &parsed, nil
}

func Produce(ctx context.Context, b *OllamaBackend, req *api.ChatRequest, tknCh chan string, doneCh chan struct{}, errCh chan error) {
	callbackFunc := func(cr api.ChatResponse) error {
		tknCh <- cr.Message.Content
		return nil
	}

	err := b.client.Chat(ctx, req, callbackFunc)
	if err != nil {
		logger.Error("Error in Ollama chat: " + err.Error())
		errCh <- err
	}
	close(tknCh)
}
