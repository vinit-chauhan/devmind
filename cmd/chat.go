package cmd

import (
	"strings"

	"github.com/spf13/cobra"
	"github.com/vinit-chauhan/devmind/cmd/ui"
	"github.com/vinit-chauhan/devmind/internal/handlers"
	"github.com/vinit-chauhan/devmind/internal/logger"
)

var chatCmd = &cobra.Command{
	Use:   "chat",
	Short: "Chat with the mind",
	Long:  `Chat with the mind. You can ask it anything and it will try to help you. It is a command line tool that can be used to generate, explain and fix code snippets, complete code, and more.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		message := strings.Builder{}
		ctx := cmd.Context()
		spinner := ctx.Value("spinner").(*ui.Spinner)

		for _, arg := range args {
			message.WriteString(arg + " ")
		}
		logger.Debug("Message: " + message.String())

		spinner.Start("Thinking...")
		_, err := handlers.Chat(ctx, message.String())
		if err != nil {
			return err
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(chatCmd)
}
