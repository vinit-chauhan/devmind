package agent

import "fmt"

func GetBackend(backendType string) (Backend, error) {
	switch backendType {

	case "ollama":
		return NewOllamaBackend(), nil
	case "openai":
		return NewOpenAIBackend(), nil
	default:
		return nil, fmt.Errorf("unsupported backend type: %s", backendType)
	}
}
