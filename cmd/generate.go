package cmd

import (
	"strings"

	"github.com/spf13/cobra"
	"github.com/vinit-chauhan/devmind/cmd/ui"
	"github.com/vinit-chauhan/devmind/internal/handlers"
	"github.com/vinit-chauhan/devmind/internal/memory"
	"github.com/vinit-chauhan/devmind/internal/memory/chat"
	"github.com/vinit-chauhan/devmind/internal/utils"
)

func init() {
	generateCmd.Flags().StringP("prompt", "p", "", "Prompt to generate code from")
	generateCmd.Flags().StringP("output", "o", "", "Output file to write the generated code to")
	generateCmd.MarkFlagRequired("prompt")
	rootCmd.AddCommand(generateCmd)
}

var generateCmd = &cobra.Command{
	Use:     "generate",
	Short:   "Generate code from a prompt (Experimental)",
	Long:    `Generate code from prompt. You can specify the output file using the -o flag. If no file is specified, it will print to stdout.`,
	Example: `devmind generate -p <prompt> -o <file>`,
	RunE: func(cmd *cobra.Command, args []string) error {
		prompt, _ := cmd.Flags().GetString("prompt")
		path, _ := cmd.Flags().GetString("output")

		ctx := cmd.Context()
		spinner := ctx.Value("spinner").(*ui.Spinner)
		spinner.Start("Generating...")

		msgs := handlers.GenerateCodePrompt(prompt, path != "")

		resp, err := handlers.GenerateCode(ctx, msgs)
		if err != nil {
			return err
		}

		chats := []chat.Chat{
			chat.New("user", prompt),
			chat.New("assistant", strings.TrimSuffix(resp, "\n")),
		}
		memory.Brain.AddChatToMemory(chats)

		if path != "" {
			// write to file
			err = utils.WriteToFile(path, []byte(resp))
			if err != nil {
				return err
			}
		}
		return nil
	},
}
