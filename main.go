package main

import (
	"context"
	"os"
	"os/signal"

	"github.com/vinit-chauhan/devmind/cmd"
	"github.com/vinit-chauhan/devmind/config"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)

	// Initialize the configuration
	config.InitConfig()
	// Execute the root command
	go cmd.Execute(ctx, stop)

	<-ctx.Done()
}
