package cmd

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use:   "mordred",
	Short: "A Git repository analysis tool",
	Long:  "Mordred is a CLI tool that parses blobs and commits in a Git repository to recover sensitive or personal data such as secrets, API keys, emails etc.",
}

// Execute Mordred.
func Execute() error {
	return rootCmd.Execute()
}
