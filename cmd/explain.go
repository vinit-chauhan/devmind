package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/vinit-chauhan/devmind/internal/handlers"
	"github.com/vinit-chauhan/devmind/internal/utils"
)

var explainCmd = &cobra.Command{
	Use:   "explain",
	Short: "Explain a code snippet",
	Long:  `Explain a code snippet. You can provide a code snippet and it will explain it to you in detail. It is a command line tool that can be used to generate, explain and fix code snippets, complete code, and more.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Explain command executed")
	},
	Args:      cobra.ExactArgs(1),
	ValidArgs: []string{"file"},
}

var fileSubCmd = &cobra.Command{
	Use:   "file",
	Short: "File to explain",
	Long:  `File to explain. You can provide a file path and it will explain the code in the file to you in detail.`,
	Run: func(cmd *cobra.Command, args []string) {
		path := args[0]
		l, _ := cmd.Flags().GetString("lines")
		lines, err := utils.ParseLineRange(l)
		if err != nil {
			fmt.Println("Error parsing line range:", err)
			return
		}

		handlers.Explain(path, lines)
	},
}

func init() {
	fileSubCmd.Flags().StringP("lines", "l", "", "line range to explain (eg. 1-10)")
	explainCmd.AddCommand(fileSubCmd)

	rootCmd.AddCommand(explainCmd)
}
