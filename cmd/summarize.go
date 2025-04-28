package cmd

import (
	"strings"

	"github.com/spf13/cobra"
	"github.com/vinit-chauhan/devmind/cmd/ui"
	"github.com/vinit-chauhan/devmind/internal/handlers"
	"github.com/vinit-chauhan/devmind/internal/logger"
)

var summarize = &cobra.Command{
	Use:   "summarize",
	Short: "Summarize code from a file or stdin",
	Long:  `Summarize the content provided in the command line arguments.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := cmd.Context()
		spinner := ctx.Value("spinner").(*ui.Spinner)

		text := strings.Join(args, " ")
		logger.Debug("Message: " + text)

		spinner.Start("Thinking...")
		msgs := handlers.GenerateSummarizePrompt(text)
		_, err := handlers.Summarize(ctx, msgs)
		if err != nil {
			return err
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(summarize)
}
