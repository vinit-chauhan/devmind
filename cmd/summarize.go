package cmd

import (
	"errors"
	"strings"

	"github.com/spf13/cobra"
	"github.com/vinit-chauhan/devmind/cmd/ui"
	"github.com/vinit-chauhan/devmind/internal/handlers"
	"github.com/vinit-chauhan/devmind/internal/memory"
	"github.com/vinit-chauhan/devmind/internal/memory/chat"
	"github.com/vinit-chauhan/devmind/internal/utils"
)

var summarizeCmd = &cobra.Command{
	Use:   "summarize",
	Short: "Summarize code from a file or stdin",
	Long:  `Summarize the content provided in the command line arguments.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		var text string
		ctx := cmd.Context()
		spinner := ctx.Value("spinner").(*ui.Spinner)

		spinner.Start("Reading file...")
		path, _ := cmd.Flags().GetString("file")
		if path != "" {
			content, err := utils.ReadFileContent(path, utils.LineRange{})
			if err != nil {
				return err
			}
			text = string(content)
		} else {
			if len(args) == 0 {
				return errors.New("no text provided")
			}

			text = strings.Join(args, " ")
		}
		spinner.Stop()

		spinner.Start("Thinking...")
		msgs := handlers.GenerateSummarizePrompt(text)
		resp, err := handlers.Summarize(ctx, msgs)
		if err != nil {
			return err
		}

		chats := []chat.Chat{
			chat.New("user", "summarize the content provided by user"),
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
	},
}

func init() {
	rootCmd.AddCommand(summarizeCmd)
	summarizeCmd.Flags().StringP("file", "f", "", "File to summarize")
	summarizeCmd.Flags().StringP("output", "o", "", "Output file to write the summary to")
}
