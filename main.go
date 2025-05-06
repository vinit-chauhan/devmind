package main

import (
	"context"
	"os"
	"os/signal"

	"github.com/vinit-chauhan/devmind/cmd"
	"github.com/vinit-chauhan/devmind/cmd/ui"
	"github.com/vinit-chauhan/devmind/config"
	"github.com/vinit-chauhan/devmind/internal/constants"
	"github.com/vinit-chauhan/devmind/internal/logger"
	"github.com/vinit-chauhan/devmind/internal/memory"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	done := make(chan struct{})
	ctx = context.WithValue(ctx, "done", done)
	ctx = context.WithValue(ctx, "spinner", ui.NewSpinner(ctx))

	// Set os specific variables
	constants.Init()

	// Initialize the configuration
	logger.Init()
	config.Init()
	memory.Init()

	// Execute the root command
	go cmd.Execute(ctx, stop)

	<-ctx.Done()
	<-done
	logger.Debug("Received shutdown signal, shutting down...")
}
