package handlers

import (
	"context"

	"github.com/vinit-chauhan/devmind/config"
	"github.com/vinit-chauhan/devmind/internal/agent"
	"github.com/vinit-chauhan/devmind/internal/agent/types"
	"github.com/vinit-chauhan/devmind/internal/logger"
	"github.com/vinit-chauhan/devmind/internal/utils"
)

func Summarize(ctx context.Context, msgs []types.Message) (string, error) {
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

func GenerateSummarizePrompt(content string) []types.Message {
	prompt := "Summarize the following text or code in 3 to 4 sentences:\n\n"
	prompt += "\n" + content + "\n\n"
	prompt += "Only provide a summary of the code snippet above."

	return []types.Message{
		{Role: "system", Content: utils.SystemPromptSummarize},
		{Role: "user", Content: prompt},
	}
}
