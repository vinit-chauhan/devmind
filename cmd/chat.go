package cmd

import (
	"strings"

	"github.com/spf13/cobra"
	"github.com/vinit-chauhan/devmind/cmd/ui"
	"github.com/vinit-chauhan/devmind/internal/handlers"
	"github.com/vinit-chauhan/devmind/internal/logger"
	"github.com/vinit-chauhan/devmind/internal/memory"
	"github.com/vinit-chauhan/devmind/internal/memory/chat"
	"github.com/vinit-chauhan/devmind/internal/utils"
)

var chatCmd = &cobra.Command{
	Use:   "chat",
	Short: "Chat with the mind",
	Long:  `Chat with the mind. You can ask it anything and it will try to help you. It is a command line tool that can be used to generate, explain and fix code snippets, complete code, and more.`,
	RunE:  runChat,
}

func init() {
	rootCmd.AddCommand(chatCmd)
	chatCmd.Flags().StringP("output", "o", "", "Output file to write the chat to")
}

func runChat(cmd *cobra.Command, args []string) error {
	ctx := cmd.Context()
	spinner := ctx.Value("spinner").(*ui.Spinner)

	message := strings.Join(args, " ")
	logger.Debug("Message: " + message)

	spinner.Start("Thinking...")
	msgs := handlers.GenerateChatPrompt(message)
	resp, err := handlers.Chat(ctx, msgs)
	if err != nil {
		return err
	}

	chats := []chat.Chat{
		chat.New("user", message),
		chat.New("assistant", strings.TrimSuffix(resp, "\n")),
	}
	memory.Brain.AddChatToMemory(chats)

	if file, _ := cmd.Flags().GetString("output"); file != "" {
		// write to file
		err := utils.WriteToFile(file, []byte(resp))
		if err != nil {
			return err
		}
	}

	return nil
}
