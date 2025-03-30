package cmd

import (
	"fmt"
	"os"

	"github.com/cidverse/normalizeci/pkg/normalizer"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

func denormalizeCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "denormalize",
		Short: "denormalizes information about the current CI environment",
		Run: func(cmd *cobra.Command, args []string) {
			format, _ := cmd.Flags().GetString("format")
			outputFile, _ := cmd.Flags().GetString("output")
			strict, _ := cmd.Flags().GetBool("strict")
			targets, _ := cmd.Flags().GetStringArray("target")

			// run normalization
			var normalized, err = normalizer.Normalize()
			if err != nil {
				log.Fatal().Err(err).Msg("normalization failed")
			}

			// output
			outputEnv := make(map[string]string)

			// targets
			if len(targets) > 0 {
				for _, target := range targets {
					denormalized, err := normalizer.Denormalize(normalizer.Options{}, target, normalized)
					if err != nil {
						log.Fatal().Err(err).Str("target", target).Msg("denormalization failed")
					}

					for key, value := range denormalized {
						outputEnv[key] = value
					}
				}
			}

			// set process env
			normalizer.SetProcessEnvironment(outputEnv)

			// content?
			content, err := normalizer.FormatEnvironment(outputEnv, format)
			if err != nil {
				log.Fatal().Str("format", format).Str("supported", "export,powershell,cmd").Msg("unsupported format!")
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
			if len(outputFile) > 0 {
				fileOutput(outputFile, content)
			} else {
				consoleOutput(content)
			}
		},
	}

	cmd.PersistentFlags().StringP("format", "f", normalizer.GetDefaultFormat(), "The format in which to store the normalized variables. (export, powershell, cmd)")
	cmd.PersistentFlags().StringP("output", "o", "", "Write output to this file instead of writing it to stdout.")
	cmd.PersistentFlags().Bool("strict", false, "Validate the generated variables against the spec and fail on errors?")
	cmd.PersistentFlags().StringArrayP("target", "t", []string{}, "Additionally generates the environment for the target ci services")

	return cmd
}
