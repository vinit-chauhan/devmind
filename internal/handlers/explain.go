package handlers

import (
	"context"
	"strings"

	"github.com/vinit-chauhan/devmind/config"
	"github.com/vinit-chauhan/devmind/internal/agent"
	"github.com/vinit-chauhan/devmind/internal/logger"
)

func Explain(ctx context.Context, prompt string) (string, error) {

	backend, err := agent.GetBackend(config.Config)
	if err != nil {
		msg := "Error getting backend: " + err.Error()
		logger.Error(msg)
		return "", err
	}

	resp, err := backend.Respond(ctx, prompt)
	if err != nil {
		msg := "Error getting response: " + err.Error()
		logger.Error(msg)
		return "", err
	}

	return resp.GetResponse(), nil
}

func GeneratePrompt(content []byte) string {
	prompt := strings.Builder{}
	prompt.WriteString("Explain the following code snippet in detail:\n\n")
	prompt.WriteString("```go\n")
	prompt.Write(content)
	prompt.WriteString("```\n\n")
	prompt.WriteString("Only provide a detailed explanation of the code snippet above.")

	return prompt.String()
}
