package cmd

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/vinit-chauhan/devmind/cmd/ui"
)

var rootCmd = &cobra.Command{
	Use:   "devmind",
	Short: "A mind that helps you with all things development",
	Long:  `A mind that helps you with all things development. It can help you with code generation, code completion, and more. It is a command line tool that can be used to generate, explain and fix code snippets, complete code, and more.`,
}

func Execute(ctx context.Context, stop context.CancelFunc) {
	defer stop()
	spinner := ui.NewSpinner(ctx)
	ctx = context.WithValue(ctx, "spinner", spinner)

	if err := rootCmd.ExecuteContext(ctx); err != nil {
		fmt.Println(err)
		return
	}
}
