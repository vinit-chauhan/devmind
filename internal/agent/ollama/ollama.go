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

func (b *OllamaBackend) Respond(ctx context.Context, msgs []types.Message) (types.Readable, error) {
	full := strings.Builder{}
	defer func() {
		close(errCh)
	}()

	// Produce the response in a goroutine
	go Produce(ctx, b, msgs, &full)

	// Print the response as it comes to stdout
	go consumer.Consume(ctx, tknCh, doneCh)

	// Wait for the response to finish
	<-doneCh

	parsed := &OllamaChatResponse{
		Response: api.Message{
			Role:    "assistant",
			Content: full.String(),
		},
		Done: true,
	}

	return parsed, nil
}

func Produce(ctx context.Context, b *OllamaBackend, msgs []types.Message, full *strings.Builder) {
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
				full.WriteString(cr.Message.Content)
				return nil
			}
		}

	}

	messages := make([]api.Message, len(msgs))
	for _, msg := range msgs {
		messages = append(messages, api.Message{
			Role:    msg.Role,
			Content: msg.Content,
		})
	}

	req := api.ChatRequest{
		Model:    b.conf.Model,
		Messages: messages,
	}

	err := b.client.Chat(ctx, &req, callbackFunc)
	if err != nil {
		logger.Error("Error in Ollama chat: " + err.Error())
		errCh <- err
	}
}
