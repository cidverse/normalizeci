package cmd

import (
	"fmt"
	"github.com/cidverse/normalizeci/pkg/ncispec"
	"github.com/cidverse/normalizeci/pkg/normalizeci"
	"github.com/rs/zerolog/log"
	"os"
)

func normalizationCommand(format string, hostEnv bool, output string, strict bool) {
	// run normalization
	var normalizedEnv = normalizeci.RunDefaultNormalization()

	normalizeci.ConfigureProcessEnvironment(normalizedEnv)

	// set normalized variables in current session
	var nci = ncispec.OfMap(normalizedEnv)
	if hostEnv == false {
		nci.DATA = nil // exclude hostEnv from generation
	}
	content := normalizeci.FormatEnvironment(ncispec.ToMap(nci), format)

	// content?
	if len(content) == 0 {
		log.Error().Msg("unsupported format!")
		os.Exit(1)
	}

	// validate?
	if strict {
		errors := nci.Validate()
		if len(errors) > 0 {
			for _, line := range errors {
				fmt.Printf("%s: %s [%s]\n", line.Field, line.Description, line.Value)
			}
			os.Exit(1)
		}
	}

	// output
	if len(output) > 0 {
		fileOutput(output, content)
	} else {
		consoleOutput(content)
	}
}
