package agent

import (
	"fmt"

	"github.com/vinit-chauhan/devmind/config"
	"github.com/vinit-chauhan/devmind/internal/agent/ollama"
	"github.com/vinit-chauhan/devmind/internal/agent/openai"
	"github.com/vinit-chauhan/devmind/internal/agent/types"
)

func GetBackend(conf config.AppConfig) (types.Backend, error) {
	switch conf.Backend {
	case "ollama":
		return ollama.NewOllamaBackend(conf.Ollama), nil
	case "openai":
		return openai.NewOpenAIBackend(conf.OpenAI), nil
	default:
		return nil, fmt.Errorf("unsupported backend type: %s", conf.Backend)
	}
}
