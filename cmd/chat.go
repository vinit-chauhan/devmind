package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var chatCmd = &cobra.Command{
	Use:   "chat",
	Short: "Chat with the mind",
	Long:  `Chat with the mind. You can ask it anything and it will try to help you. It is a command line tool that can be used to generate, explain and fix code snippets, complete code, and more.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Chat command executed")

		message := strings.Builder{}
		for _, arg := range args {
			message.WriteString(arg + " ")
		}
		fmt.Printf("Message: %s\n", message.String())
	},
}

func init() {
	rootCmd.AddCommand(chatCmd)
}
