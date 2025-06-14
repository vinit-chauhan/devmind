package handlers

import (
	"context"

	"github.com/vinit-chauhan/devmind/config"
	"github.com/vinit-chauhan/devmind/internal/agent"
	"github.com/vinit-chauhan/devmind/internal/agent/types"
	"github.com/vinit-chauhan/devmind/internal/logger"
	"github.com/vinit-chauhan/devmind/internal/memory"
	"github.com/vinit-chauhan/devmind/internal/utils"
)

func Chat(ctx context.Context, msgs []types.Message) (string, error) {
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

func GenerateChatPrompt(prompt string) []types.Message {
	systemPrompt := utils.SystemPrompt
	return []types.Message{
		{Role: "system", Content: systemPrompt},
		{Role: "assistant", Content: memory.Brain.GetMemoryPrompt()},
		{Role: "user", Content: prompt},
	}
}
