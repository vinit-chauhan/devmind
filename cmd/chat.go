package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/vinit-chauhan/devmind/config"
	"github.com/vinit-chauhan/devmind/internal/agent"
	"github.com/vinit-chauhan/devmind/internal/logger"
)

var chatCmd = &cobra.Command{
	Use:   "chat",
	Short: "Chat with the mind",
	Long:  `Chat with the mind. You can ask it anything and it will try to help you. It is a command line tool that can be used to generate, explain and fix code snippets, complete code, and more.`,
	Run: func(cmd *cobra.Command, args []string) {
		logger.Info("Chat command executed")

		backend, err := agent.GetBackend(config.Config.Backend)
		if err != nil {
			msg := "Error getting backend: " + err.Error()
			logger.Error(msg)
			fmt.Fprintln(os.Stderr, msg)
			return
		}

		message := strings.Builder{}
		for _, arg := range args {
			message.WriteString(arg + " ")
		}
		logger.Debug("Message: " + message.String())

		resp, err := backend.Respond(message.String())
		if err != nil {
			msg := "Error getting response: " + err.Error()
			logger.Error(msg)
			fmt.Fprintln(os.Stderr, msg)
			return
		}

		fmt.Printf("Response: %s\n", resp)
	},
}

func init() {
	rootCmd.AddCommand(chatCmd)
}
