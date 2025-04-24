package ollama

import (
	"net/http"
	"net/url"

	"github.com/ollama/ollama/api"
	"github.com/vinit-chauhan/devmind/config"
	"github.com/vinit-chauhan/devmind/internal/agent/types"
	"github.com/vinit-chauhan/devmind/internal/logger"
)

type OllamaBackend struct {
	conf   *config.OllamaConfig
	client *api.Client
}

func NewOllamaBackend(conf config.OllamaConfig) types.Backend {
	host, err := url.Parse(conf.Host)
	if err != nil {
		logger.Error("Invalid Ollama host URL: " + conf.Host)
		return nil
	}

	return &OllamaBackend{
		conf:   &conf,
		client: api.NewClient(host, http.DefaultClient),
	}
}

type OllamaChatResponse struct {
	Response api.Message `json:"message"`
	Done     bool        `json:"done"`
}

func (r *OllamaChatResponse) IsDone() bool {
	return r.Done
}

func (r *OllamaChatResponse) GetResponse() string {
	return r.Response.Content
}
