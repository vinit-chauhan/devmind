package handlers

import (
	"fmt"
	"os"
	"strings"

	"github.com/vinit-chauhan/devmind/config"
	"github.com/vinit-chauhan/devmind/internal/agent"
	"github.com/vinit-chauhan/devmind/internal/logger"
)

func Chat(args []string) string {
	message := strings.Builder{}
	for _, arg := range args {
		message.WriteString(arg + " ")
	}
	logger.Debug("Message: " + message.String())

	backend, err := agent.GetBackend(config.Config.Backend)
	if err != nil {
		msg := "Error getting backend: " + err.Error()
		logger.Error(msg)
		fmt.Fprintln(os.Stderr, msg)
		os.Exit(1)
	}

	resp, err := backend.Respond(message.String())
	if err != nil {
		msg := "Error getting response: " + err.Error()
		logger.Error(msg)
		fmt.Fprintln(os.Stderr, msg)
		os.Exit(1)
	}

	return resp
}
