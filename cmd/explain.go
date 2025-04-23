package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/vinit-chauhan/devmind/cmd/ui"
	"github.com/vinit-chauhan/devmind/internal/handlers"
	"github.com/vinit-chauhan/devmind/internal/utils"
)

var explainCmd = &cobra.Command{
	Use:   "explain",
	Short: "Explain a code snippet",
	Long:  `Explain a code snippet. You can provide a code snippet and it will explain it to you in detail. It is a command line tool that can be used to generate, explain and fix code snippets, complete code, and more.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		var prompt string
		ctx := cmd.Context()
		spinner := ctx.Value("spinner").(*ui.Spinner)

		l, _ := cmd.Flags().GetString("lines")
		path, _ := cmd.Flags().GetString("file")

		spinner.Start("Reading file...")
		if path == "" {
			fmt.Println("TODO: read from stdin")
			spinner.Stop()
			return fmt.Errorf("file path is required")
		} else {
			lines, err := utils.ParseLineRange(l)
			if err != nil {
				fmt.Println("Error parsing line range:", err)
				return err
			}
			content, err := handlers.ReadFileContent(path, lines)
			if err != nil {
				fmt.Println("Error reading file:", err)
				return err
			}

			prompt = handlers.GeneratePrompt(content)
			spinner.Stop()
		}

		spinner.Start("Thinking...")
		_, err := handlers.Explain(ctx, prompt)
		if err != nil {
			return err
		}

		return nil
	},
}

func init() {
	explainCmd.Flags().StringP("lines", "l", "", "line range to explain (eg. 1-10)")
	explainCmd.Flags().StringP("file", "f", "", "file to explain")
	explainCmd.MarkFlagRequired("file")

	rootCmd.AddCommand(explainCmd)
}
