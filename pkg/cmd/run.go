package cmd

import (
	"fmt"
	"github.com/cidverse/normalizeci/pkg/ncispec"
	"github.com/cidverse/normalizeci/pkg/normalizeci"
	"github.com/rs/zerolog/log"
	"os"
)

func normalizationCommand(format string, hostEnv bool, output string, strict bool, targets []string) {
	// run normalization
	var normalizedEnv = normalizeci.RunDefaultNormalization()

	// set normalized variables in current session
	var nci = ncispec.OfMap(normalizedEnv)
	if hostEnv == false {
		nci.DATA = nil // exclude hostEnv from generation
	}
	outputEnv := ncispec.ToMap(nci)

	// set process env
	normalizeci.SetProcessEnvironment(normalizedEnv)

	// targets
	if len(targets) > 0 {
		for _, target := range targets {
			denormalized := normalizeci.RunDenormalization(target, normalizedEnv)
			for key, value := range denormalized {
				outputEnv[key] = value
			}
		}
	}

	// content?
	content := normalizeci.FormatEnvironment(outputEnv, format)
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
