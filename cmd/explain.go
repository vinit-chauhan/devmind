package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/vinit-chauhan/devmind/cmd/ui"
	"github.com/vinit-chauhan/devmind/internal/handlers"
	"github.com/vinit-chauhan/devmind/internal/logger"
	"github.com/vinit-chauhan/devmind/internal/memory"
	"github.com/vinit-chauhan/devmind/internal/memory/chat"
	"github.com/vinit-chauhan/devmind/internal/utils"
)

var explainCmd = &cobra.Command{
	Use:     "explain",
	Short:   "Explain code from a file or stdin",
	Long:    `Explain code from a file or stdin. You can specify the line range to explain using the -l flag. If no file is specified, it will read from stdin.`,
	Example: `devmind explain -f <file> -l <line range>`,
	RunE:    runExplain,
}

func init() {
	explainCmd.Flags().StringP("lines", "l", "", "line range to explain (eg. 1-10)")
	explainCmd.Flags().StringP("file", "f", "", "file to explain")
	explainCmd.Flags().StringP("output", "o", "", "Output file to write the explanation to")

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

		fmt.Println("contend:", content)
		if len(content) == 0 {
			logger.Error("No content read from stdin")
			return nil, fmt.Errorf("no content read from stdin")
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

func runExplain(cmd *cobra.Command, args []string) error {
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
	msgs := handlers.GenerateExplainPrompt(string(content))
	spinner.Stop()

	logger.Debug("Explaining the content...")
	spinner.Start("Thinking...")
	resp, err := handlers.Explain(ctx, msgs)
	if err != nil {
		return err
	}

	chats := []chat.Chat{
		chat.New("user", "Explain the content of the file given by user"),
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
