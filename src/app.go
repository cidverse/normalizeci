package main

import (
	"github.com/PhilippHeuer/normalize-ci/pkg/common"
	log "github.com/sirupsen/logrus"
	"os"
	"strings"

	"github.com/PhilippHeuer/normalize-ci/pkg/normalizeci"
)

// Version will be set at build time
var Version string

// CommitHash will be set at build time
var CommitHash string

// Init Hook
func init() {
	// Output to Stderr to not pollute stdout redirects with logs
	log.SetOutput(os.Stderr)

	// Only log the warning severity or above.
	log.SetLevel(log.TraceLevel)

	// detect debug mode
	debugValue, debugIsSet := os.LookupEnv("NCI_DEBUG")
	if debugIsSet && strings.ToLower(debugValue) == "true" {
		log.SetLevel(log.TraceLevel)
	}
}

// CLI Main Entrypoint
func main() {
	// get all environment variables
	env := common.GetFullEnv()

	// run normalization
	var normalized = normalizeci.RunNormalization(env)

	// set normalized variables in current session
	normalizeci.SetNormalizedEnvironment(normalized)
}
