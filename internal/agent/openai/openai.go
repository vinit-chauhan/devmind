package openai

import (
	"context"

	"github.com/vinit-chauhan/devmind/config"
	"github.com/vinit-chauhan/devmind/internal/agent/types"
)

func NewOpenAIBackend(conf config.OpenAIConfig) types.Backend {
	return &OpenAIBackend{
		conf: &conf,
	}
}

func (b *OpenAIBackend) Respond(ctx context.Context, msgs []types.Message) (response types.Readable, err error) {
	return &OpenAIResponse{}, nil
}
