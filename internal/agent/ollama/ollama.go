package ollama

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/vinit-chauhan/devmind/internal/agent/types"
)

func (b *OllamaBackend) Respond(prompt string) (types.Readable, error) {

	req := OllamaChatRequest{
		Model:  b.conf.Model,
		Stream: false,
		Messages: []Message{
			{Role: "system", Content: SystemPrompt},
			{Role: "user", Content: prompt},
		},
	}

	reqBody, err := json.Marshal(req)
	if err != nil {
		return types.EmptyResponse, fmt.Errorf("failed to marshal chat request: %w", err)
	}

	url := strings.TrimSuffix(b.conf.Host, "/") + "/api/chat"
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		return types.EmptyResponse, fmt.Errorf("failed to call Ollama: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return types.EmptyResponse, fmt.Errorf("ollama error [%d]: %s", resp.StatusCode, string(body))
	}

	var parsed OllamaChatResponse
	if err := json.NewDecoder(resp.Body).Decode(&parsed); err != nil {
		return types.EmptyResponse, fmt.Errorf("failed to decode chat response: %w", err)
	}

	return &parsed, nil
}
