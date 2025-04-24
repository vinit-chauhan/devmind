package handlers

import (
	"context"
	"strings"

	"github.com/vinit-chauhan/devmind/config"
	"github.com/vinit-chauhan/devmind/internal/agent"
	"github.com/vinit-chauhan/devmind/internal/agent/types"
	"github.com/vinit-chauhan/devmind/internal/logger"
	"github.com/vinit-chauhan/devmind/internal/utils"
)

func Explain(ctx context.Context, msgs []types.Message) (string, error) {

	backend, err := agent.GetBackend(config.Config)
	if err != nil {
		msg := "Error getting backend: " + err.Error()
		logger.Error(msg)
		return "", err
	}

	resp, err := backend.Respond(ctx, msgs)
	if err != nil {
		msg := "Error getting response: " + err.Error()
		logger.Error(msg)
		return "", err
	}

	return resp.GetResponse(), nil
}

func GenerateExplainPrompt(content string) []types.Message {
	prompt := strings.Builder{}
	prompt.WriteString("Explain the following code snippet in detail:\n\n")
	prompt.WriteString("```go\n")
	prompt.WriteString(content)
	prompt.WriteString("```\n\n")
	prompt.WriteString("Only provide a detailed explanation of the code snippet above.")

	return []types.Message{
		{Role: "system", Content: utils.SystemPrompt},
		{Role: "user", Content: prompt.String()},
	}
}
