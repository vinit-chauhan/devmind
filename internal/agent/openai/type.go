package openai

import (
	"github.com/vinit-chauhan/devmind/config"
	"github.com/vinit-chauhan/devmind/internal/agent/types"
)

var Backend *OpenAIBackend = nil

type OpenAIBackend struct {
	conf *config.OpenAIConfig
}

func NewOpenAIBackend(conf config.OpenAIConfig) types.Backend {
	if Backend != nil {
		return Backend
	}

	Backend = &OpenAIBackend{
		conf: &conf,
	}
	return Backend
}

func (b *OpenAIBackend) Respond(prompt string) (response string, err error) {
	return "OpenAI response", nil
}
