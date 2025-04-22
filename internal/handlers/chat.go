package handlers

import (
	"fmt"
	"os"

	"github.com/vinit-chauhan/devmind/config"
	"github.com/vinit-chauhan/devmind/internal/agent"
	"github.com/vinit-chauhan/devmind/internal/logger"
)

func Chat(message string) string {
	backend, err := agent.GetBackend(config.Config)
	if err != nil {
		msg := "Error getting backend: " + err.Error()
		logger.Error(msg)
		fmt.Fprintln(os.Stderr, msg)
		os.Exit(1)
	}

	resp, err := backend.Respond(message)
	if err != nil {
		msg := "Error getting response: " + err.Error()
		logger.Error(msg)
		fmt.Fprintln(os.Stderr, msg)
		os.Exit(1)
	}

	return resp.GetResponse()
}
