package ollama

import (
	"context"

	"github.com/ollama/ollama/api"
	"github.com/vinit-chauhan/devmind/internal/agent/types"
	"github.com/vinit-chauhan/devmind/internal/logger"
)

func (b *OllamaBackend) Respond(ctx context.Context, prompt string) (types.Readable, error) {

	req := api.ChatRequest{
		Model:  b.conf.Model,
		Stream: &b.conf.Stream,
		Messages: []api.Message{
			{Role: "system", Content: SystemPrompt},
			{Role: "user", Content: prompt},
		},
	}

	var parsed OllamaChatResponse
	var responseFunc api.ChatResponseFunc = func(cr api.ChatResponse) error {
		parsed = OllamaChatResponse{
			Response: cr.Message,
			Done:     cr.Done,
		}
		logger.Debug("Generated response: " + cr.Message.Content)
		return nil
	}

	if err := b.client.Chat(ctx, &req, responseFunc); err != nil {
		return types.EmptyResponse, err
	}

	return &parsed, nil
}
