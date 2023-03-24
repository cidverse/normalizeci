package main

import (
	"os"
	"strings"

	"github.com/cidverse/normalizeci/cli/cmd"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
	status  = "clean"
)

// Init Hook
func init() {
	// Output to Stderr to not pollute stdout redirects with logs
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	// Only log the warning severity or above.
	zerolog.SetGlobalLevel(zerolog.WarnLevel)

	// detect debug mode
	debugValue, debugIsSet := os.LookupEnv("NCI_DEBUG")
	if debugIsSet && strings.EqualFold(debugValue, "true") {
		zerolog.SetGlobalLevel(zerolog.TraceLevel)
	}

	// Set Version Information
	cmd.Version = version
	cmd.CommitHash = commit
	cmd.BuildAt = date
	cmd.RepositoryStatus = status
}

// CLI Main Entrypoint
func main() {
	cmdErr := cmd.Execute()
	if cmdErr != nil {
		log.Fatal().Err(cmdErr).Msg("internal cli library error")
	}
}
