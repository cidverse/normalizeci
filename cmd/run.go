package cmd

import (
	"fmt"
	"os"

	"github.com/cidverse/normalizeci/pkg/ncispec"
	"github.com/cidverse/normalizeci/pkg/normalizeci"
	"github.com/rs/zerolog/log"
)

func normalizationCommand(format string, output string, strict bool, targets []string) {
	// run normalization
	var normalized = normalizeci.Normalize()

	// set normalized variables in current session
	outputEnv := ncispec.ToMap(normalized)

	// set process env
	normalizeci.SetProcessEnvironment(ncispec.ToMap(normalized))

	// targets
	if len(targets) > 0 {
		for _, target := range targets {
			denormalized := normalizeci.Denormalize(target, normalized)
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
		errors := normalized.Validate()
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
