package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "cli-service",
	Short: "CLI application",
	Long:  `cli-service is a command-line tool that allows users to interact with a system for posting, liking messages, and viewing feeds.`,
}

// Execute runs the root command and initializes the CLI application.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
