package cmd

import (
	"fmt"
	"os"
	"runtime"

	"github.com/spf13/cobra"
)

// Version will be set at build time
var Version string

// RepositoryStatus will be set at build time
var RepositoryStatus string

// CommitHash will be set at build time
var CommitHash string

// BuildAt will be set at build time
var BuildAt string

func init() {
	rootCmd.AddCommand(versionCmd)
	versionCmd.PersistentFlags().Bool("short", false, "only prints the plain version number without any other information")
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "print the version",
	Run: func(cmd *cobra.Command, args []string) {
		short, _ := cmd.Flags().GetBool("short")
		if short {
			fmt.Fprintf(os.Stdout, Version)
		} else {
			fmt.Fprintf(os.Stdout, "GitVersion:    %s\n", Version)
			fmt.Fprintf(os.Stdout, "GitCommit:     %s\n", CommitHash)
			fmt.Fprintf(os.Stdout, "GitTreeState:  %s\n", RepositoryStatus)
			fmt.Fprintf(os.Stdout, "BuildDate:     %s\n", BuildAt)
			fmt.Fprintf(os.Stdout, "GoVersion:     %s\n", runtime.Version())
			fmt.Fprintf(os.Stdout, "Compiler:      %s\n", runtime.Compiler)
			fmt.Fprintf(os.Stdout, "Platform:      %s\n", runtime.GOOS+"/"+runtime.GOARCH)
		}
	},
}
