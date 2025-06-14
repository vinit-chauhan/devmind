package cmd

import (
	"context"

	"github.com/spf13/cobra"
	"github.com/vinit-chauhan/devmind/internal/logger"
)

var rootCmd = &cobra.Command{
	Use:   "devmind",
	Short: "A mind that helps you with all things development",
	Long:  `A mind that helps you with all things development. It can help you with code generation, code completion, and more. It is a command line tool that can be used to generate, explain and fix code snippets, complete code, and more. to store the response redirect the output to a file using the redirect operation flag.`,
}

func Execute(ctx context.Context, stop context.CancelFunc) {
	done := ctx.Value("done").(chan struct{})
	defer func() {
		logger.Debug("Closing root command")
		stop()
		done <- struct{}{}
	}()

	if err := rootCmd.ExecuteContext(ctx); err != nil {
		logger.Error("Error executing command: " + err.Error())
		return
	}

	logger.Debug("Command executed successfully")
}
