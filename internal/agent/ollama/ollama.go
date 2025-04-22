package ollama

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"github.com/vinit-chauhan/devmind/config"
	"github.com/vinit-chauhan/devmind/internal/agent/types"
	"github.com/vinit-chauhan/devmind/internal/logger"
)

func NewOllamaBackend(conf config.OllamaConfig) types.Backend {
	return &OllamaBackend{
		conf: &conf,
	}
}

func (b *OllamaBackend) Respond(prompt string) (response types.Response, err error) {
	response = types.EmptyResponse{}
	req := NewOllamaRequest(b.conf).WithPrompt(prompt)

	reqBody, err := json.Marshal(req)
	if err != nil {
		return
	}

	endpoint := b.conf.Host
	if endpoint[len(endpoint)-1] == '/' {
		endpoint = endpoint[:len(endpoint)-1]
	}
	endpoint += "/api/generate"

	resp, err := http.Post(endpoint, "application/json", bytes.NewReader(reqBody))
	if err != nil {
		logger.Error("Error making request to Ollama: " + err.Error())
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		respBody, _ := io.ReadAll(resp.Body)
		logger.Error("Error response from Ollama: " + resp.Status + " : " + string(respBody))
		return
	}

	var parsed OllamaResponse
	if err = json.NewDecoder(resp.Body).Decode(&parsed); err != nil {
		logger.Error("Error decoding response from Ollama: " + err.Error())
		return
	}

	return &parsed, nil
}
