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
		spinner := ui.NewSpinner(cmd.Context())
		l, _ := cmd.Flags().GetString("lines")
		path, _ := cmd.Flags().GetString("file")
		var prompt string

		if path == "" {
			fmt.Println("Error: file path is required")
			return fmt.Errorf("file path is required")
		} else {
			spinner.Start("Reading file...")
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
		resp, err := handlers.Explain(cmd.Context(), prompt)
		if err != nil {
			return err
		}

		spinner.Stop()
		fmt.Printf("Response: \n%s", resp)
		return nil
	},
}

func init() {
	explainCmd.Flags().StringP("lines", "l", "", "line range to explain (eg. 1-10)")
	explainCmd.Flags().StringP("file", "f", "", "file to explain")
	explainCmd.MarkFlagRequired("file")

	rootCmd.AddCommand(explainCmd)
}
