package cmd

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/cidverse/normalizeci/pkg/normalizer"
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
				slog.With("err", err).Error("failed to normalize CI environment")
				os.Exit(1)
			}

			// output
			outputEnv := make(map[string]string)

			// targets
			if len(targets) > 0 {
				for _, target := range targets {
					denormalized, err := normalizer.Denormalize(normalizer.Options{}, target, normalized)
					if err != nil {
						slog.With("err", err).With("target", target).Error("failed to denormalize CI environment")
						os.Exit(1)
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
				slog.With("format", format).With("supported-formats", "export,powershell,cmd").Error("unsupported format!")
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
