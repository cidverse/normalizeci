package main

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
	"strings"

	"github.com/cidverse/normalizeci/pkg/normalizeci"
)

// Version will be set at build time
var Version string

// CommitHash will be set at build time
var CommitHash string

// Init Hook
func init() {
	// Output to Stderr to not pollute stdout redirects with logs
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	// Only log the warning severity or above.
	zerolog.SetGlobalLevel(zerolog.WarnLevel)

	// detect debug mode
	debugValue, debugIsSet := os.LookupEnv("NCI_DEBUG")
	if debugIsSet && strings.ToLower(debugValue) == "true" {
		zerolog.SetGlobalLevel(zerolog.TraceLevel)
	}
}

// CLI Main Entrypoint
func main() {
	// run normalization
	var normalized = normalizeci.RunDefaultNormalization()

	// set normalized variables in current session
	normalizeci.SetNormalizedEnvironment(normalized)
}
