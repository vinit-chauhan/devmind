package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/vinit-chauhan/devmind/internal/handlers"
	"github.com/vinit-chauhan/devmind/internal/logger"
)

var chatCmd = &cobra.Command{
	Use:   "chat",
	Short: "Chat with the mind",
	Long:  `Chat with the mind. You can ask it anything and it will try to help you. It is a command line tool that can be used to generate, explain and fix code snippets, complete code, and more.`,
	Run: func(cmd *cobra.Command, args []string) {
		logger.Info("Chat command executed")
		resp := handlers.Chat(args)
		fmt.Printf("Response: %s\n", resp)
	},
}

func init() {
	rootCmd.AddCommand(chatCmd)
}
