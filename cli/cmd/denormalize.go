package cmd

import (
	"fmt"
	"os"

	"github.com/cidverse/normalizeci/pkg/ncispec"
	"github.com/cidverse/normalizeci/pkg/normalizer"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(denormalizeCmd)
	denormalizeCmd.PersistentFlags().StringP("format", "f", normalizer.GetDefaultFormat(), "The format in which to store the normalized variables. (export, powershell, cmd)")
	denormalizeCmd.PersistentFlags().StringP("output", "o", "", "Write output to this file instead of writing it to stdout.")
	denormalizeCmd.PersistentFlags().Bool("strict", false, "Validate the generated variables against the spec and fail on errors?")
	denormalizeCmd.PersistentFlags().StringArrayP("target", "t", []string{}, "Additionally generates the environment for the target ci services")
}

var denormalizeCmd = &cobra.Command{
	Use:   "denormalize",
	Short: "denormalizes information about the current CI environment",
	Run: func(cmd *cobra.Command, args []string) {
		format, _ := cmd.Flags().GetString("format")
		outputFile, _ := cmd.Flags().GetString("output")
		strict, _ := cmd.Flags().GetBool("strict")
		targets, _ := cmd.Flags().GetStringArray("target")

		// run normalization
		var normalized = normalizer.Normalize()
		outputEnv := make(map[string]string)
		normalizer.SetProcessEnvironment(ncispec.ToMap(normalized))

		// targets
		if len(targets) > 0 {
			for _, target := range targets {
				denormalized := normalizer.Denormalize(target, normalized)
				for key, value := range denormalized {
					outputEnv[key] = value
				}
			}
		}

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
