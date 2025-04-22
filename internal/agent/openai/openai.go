package openai

import (
	"github.com/vinit-chauhan/devmind/config"
	"github.com/vinit-chauhan/devmind/internal/agent/types"
)

func NewOpenAIBackend(conf config.OpenAIConfig) types.Backend {
	return &OpenAIBackend{
		conf: &conf,
	}
}

func (b *OpenAIBackend) Respond(prompt string) (response types.Response, err error) {
	return &OpenAIResponse{}, nil
}
