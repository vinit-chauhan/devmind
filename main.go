package main

import (
	"context"
	"os"
	"os/signal"

	"github.com/vinit-chauhan/devmind/cmd"
	"github.com/vinit-chauhan/devmind/cmd/ui"
	"github.com/vinit-chauhan/devmind/config"
	"github.com/vinit-chauhan/devmind/internal/logger"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	done := make(chan struct{})
	ctx = context.WithValue(ctx, "done", done)
	ctx = context.WithValue(ctx, "spinner", ui.NewSpinner(ctx))

	// Initialize the configuration
	config.InitConfig()
	// Execute the root command
	go cmd.Execute(ctx, stop)

	<-ctx.Done()
	<-done
	logger.Debug("Received shutdown signal, shutting down...")
}
