package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   `normalizeci`,
	Short: `normalizeci provides a foundation for platform-agnostic CI-CD processes.`,
	Long:  `normalizeci provides a foundation for platform-agnostic CI-CD processes.`,
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Help()
		os.Exit(0)
	},
}

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}
