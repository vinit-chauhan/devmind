package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "devmind",
	Short: "A mind that helps you with all things development",
	Long:  `A mind that helps you with all things development. It can help you with code generation, code completion, and more. It is a command line tool that can be used to generate, explain and fix code snippets, complete code, and more.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		return
	}
}
