package openai

import (
	"github.com/vinit-chauhan/devmind/config"
)

type OpenAIBackend struct {
	conf *config.OpenAIConfig
}

type OpenAIResponse struct {
	Model    string `json:"model"`
	Response string `json:"response"`
	Done     bool   `json:"done"`
}

func (r *OpenAIResponse) IsDone() bool {
	return r.Done
}

func (r *OpenAIResponse) GetResponse() string {
	return r.Response
}
