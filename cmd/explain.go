package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
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
		path, _ := cmd.Flags().GetString("path")
		fmt.Printf("File subcommand executed with path: %s\n", path)
	},
}

func init() {
	fileSubCmd.Flags().StringP("path", "p", "", "Path of the file to explain")
	fileSubCmd.MarkFlagRequired("path")
	explainCmd.AddCommand(fileSubCmd)

	rootCmd.AddCommand(explainCmd)
}
