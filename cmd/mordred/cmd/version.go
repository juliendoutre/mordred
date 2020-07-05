package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var hash string

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print Mordred's version",
	Long:  "Print the commit hash of this Mordred's release",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Mordred: %s\n", hash)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
