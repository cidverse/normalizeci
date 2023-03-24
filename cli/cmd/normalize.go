package cmd

import (
	"fmt"
	"os"

	"github.com/cidverse/normalizeci/pkg/ncispec"
	"github.com/cidverse/normalizeci/pkg/normalizeci"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(normalizeCmd)
	normalizeCmd.PersistentFlags().StringP("format", "f", normalizeci.GetDefaultFormat(), "The format in which to store the normalized variables. (export, powershell, cmd)")
	normalizeCmd.PersistentFlags().StringP("output", "o", "", "Write output to this file instead of writing it to stdout.")
	normalizeCmd.PersistentFlags().Bool("strict", false, "Validate the generated variables against the spec and fail on errors?")
	normalizeCmd.PersistentFlags().BoolP("version", "v", false, "all software has versions, this prints version information for normalizeci")
	normalizeCmd.PersistentFlags().StringArrayP("target", "t", []string{}, "Additionally generates the environment for the target ci services")
}

var normalizeCmd = &cobra.Command{
	Use:   "normalize",
	Short: "normalizes information about the current CI environment",
	Run: func(cmd *cobra.Command, args []string) {
		format, _ := cmd.Flags().GetString("format")
		outputFile, _ := cmd.Flags().GetString("output")
		strict, _ := cmd.Flags().GetBool("strict")
		targets, _ := cmd.Flags().GetStringArray("target")

		// run normalization
		var normalized = normalizeci.Normalize()
		outputEnv := ncispec.ToMap(normalized)
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

		// format content
		content, err := normalizeci.FormatEnvironment(outputEnv, format)
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
