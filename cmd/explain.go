package cmd

import (
	"github.com/spf13/cobra"
	"github.com/vinit-chauhan/devmind/cmd/ui"
	"github.com/vinit-chauhan/devmind/internal/handlers"
	"github.com/vinit-chauhan/devmind/internal/logger"
	"github.com/vinit-chauhan/devmind/internal/utils"
)

var explainCmd = &cobra.Command{
	Use:     "explain",
	Short:   "Explain code from a file or stdin",
	Long:    `Explain code from a file or stdin. You can specify the line range to explain using the -l flag. If no file is specified, it will read from stdin.`,
	Example: `devmind explain -f <file> -l <line range>`,
	RunE: func(cmd *cobra.Command, args []string) error {
		var prompt string
		ctx := cmd.Context()
		spinner := ctx.Value("spinner").(*ui.Spinner)

		lines, _ := cmd.Flags().GetString("lines")
		path, _ := cmd.Flags().GetString("file")

		spinner.Start("Reading file...")
		content, err := readContent(path, lines)
		if err != nil {
			return err
		}
		logger.Debug("Generating prompt...")
		prompt = handlers.GeneratePrompt(content)
		spinner.Stop()

		logger.Debug("Explaining the content...")
		spinner.Start("Thinking...")
		_, err = handlers.Explain(ctx, prompt)
		if err != nil {
			return err
		}

		return nil
	},
}

func init() {
	explainCmd.Flags().StringP("lines", "l", "", "line range to explain (eg. 1-10)")
	explainCmd.Flags().StringP("file", "f", "", "file to explain")

	rootCmd.AddCommand(explainCmd)
}

func readContent(path, lines string) ([]byte, error) {
	var content []byte
	var err error

	if path == "" {
		//read from stdin
		logger.Debug("Reading from stdin")
		content, err = utils.ReadStdin()
		if err != nil {
			logger.Error("Error reading from stdin:" + err.Error())
			return nil, err
		}
	} else {
		//read from file
		lr, err := utils.ParseLineRange(lines)
		if err != nil {
			logger.Error("Error parsing line range:" + err.Error())
			return nil, err
		}

		content, err = utils.ReadFileContent(path, lr)
		if err != nil {
			logger.Error("Error reading file:" + err.Error())
			return nil, err
		}
	}

	return content, nil
}
