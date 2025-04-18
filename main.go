package main

import (
	"github.com/vinit-chauhan/devmind/cmd"
	"github.com/vinit-chauhan/devmind/config"
)

func main() {
	// Initialize the configuration
	config.InitConfig()

	// Execute the root command
	cmd.Execute()
}
