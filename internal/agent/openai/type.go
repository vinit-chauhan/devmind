package openai

import (
	"github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
	"github.com/vinit-chauhan/devmind/config"
	"github.com/vinit-chauhan/devmind/internal/agent/types"
	"github.com/vinit-chauhan/devmind/internal/logger"
)

type OpenAIBackend struct {
	conf   *config.OpenAIConfig
	client *openai.Client
}

type OpenAIResponse struct {
	Response types.Message `json:"message"`
	Done     bool          `json:"done"`
}

func NewOpenAIBackend(conf config.OpenAIConfig) types.Backend {
	if conf.ApiKey == "" {
		logger.Error("ApiKey not provided in the config")
		return nil
	}

	client := openai.NewClient(option.WithAPIKey(conf.ApiKey))

	return &OpenAIBackend{
		conf:   &conf,
		client: &client,
	}
}

func (r *OpenAIResponse) IsDone() bool {
	return r.Done
}

func (r *OpenAIResponse) GetResponse() string {
	return r.Response.Content
}
