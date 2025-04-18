package ollama

import (
	"github.com/vinit-chauhan/devmind/config"
	"github.com/vinit-chauhan/devmind/internal/agent/types"
)

type OllamaBackend struct {
	conf *config.OllamaConfig
}

func NewOllamaBackend(conf config.OllamaConfig) types.Backend {
	return &OllamaBackend{
		conf: &conf,
	}
}

func (b *OllamaBackend) Respond(prompt string) (response string, err error) {
	return "Ollama response", nil
}
