package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"runtime"
)

// Version will be set at build time
var Version string

// CommitHash will be set at build time
var CommitHash string

// BuildAt will be set at build time
var BuildAt string

var rootCmd = &cobra.Command{
	Use:   `normalizeci`,
	Short: `normalizeci provides a foundation for platform-agnostic CICD processes.`,
	Long:  `normalizeci provides a foundation for platform-agnostic CICD processes.`,
	PreRunE: func(cmd *cobra.Command, args []string) error {
		version, _ := cmd.Flags().GetBool("version")
		if version {
			printVersion()
			os.Exit(0)
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		format, _ := cmd.Flags().GetString("format")
		format = defaultFormat(format)
		output, _ := cmd.Flags().GetString("output")
		hostEnv, _ := cmd.Flags().GetBool("hostEnv")
		strict, _ := cmd.Flags().GetBool("strict")
		targets, _ := cmd.Flags().GetStringArray("target")

		normalizationCommand(format, hostEnv, output, strict, targets)
	},
}

func init() {
	rootCmd.PersistentFlags().StringP("format", "f", "systemdefault", "The format in which to store the normalized variables. (export, powershell, cmd)")
	rootCmd.PersistentFlags().StringP("output", "o", "", "Write output to this file instead of writing it to stdout.")
	rootCmd.PersistentFlags().Bool("hostenv", false, "Should include os env along with normalized variables into the target?")
	rootCmd.PersistentFlags().Bool("strict", false, "Validate the generated variables against the spec and fail on errors?")
	rootCmd.PersistentFlags().BoolP("version", "v", false, "all software has versions, this prints version information for normalizeci")
	rootCmd.PersistentFlags().StringArrayP("target", "t", []string{}, "Additionally generates the environment for the target ci services")
}

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

func printVersion() {
	fmt.Println("normalizeci " + Version + "-" + CommitHash + " " + runtime.GOOS + "/" + runtime.GOARCH + " BuildDate=" + BuildAt)
}
