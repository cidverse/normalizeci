package cmd

import (
	"os"
	"strings"

	"github.com/cidverse/cidverseutils/zerologconfig"
	"github.com/spf13/cobra"
)

var cfg zerologconfig.LogConfig

func RootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   `normalizeci`,
		Short: `normalizeci provides a foundation for platform-agnostic CI-CD processes.`,
		Long:  `normalizeci provides a foundation for platform-agnostic CI-CD processes.`,
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			if cfg.LogLevel == "" {
				if os.Getenv("NCI_LOG_LEVEL") != "" {
					cfg.LogLevel = os.Getenv("NCI_LOG_LEVEL")
				} else if os.Getenv("NCI_DEBUG") == "true" { // legacy 0.x toggle for debug mode
					cfg.LogLevel = "trace"
				}
			}

			zerologconfig.Configure(cfg)
		},
		Run: func(cmd *cobra.Command, args []string) {
			_ = cmd.Help()
			os.Exit(0)
		},
	}

	cmd.PersistentFlags().StringVar(&cfg.LogLevel, "log-level", "info", "log level - allowed: "+strings.Join(zerologconfig.ValidLogLevels, ","))
	cmd.PersistentFlags().StringVar(&cfg.LogFormat, "log-format", "color", "log format - allowed: "+strings.Join(zerologconfig.ValidLogFormats, ","))
	cmd.PersistentFlags().BoolVar(&cfg.LogCaller, "log-caller", false, "include caller in log functions")

	cmd.AddCommand(versionCmd())
	cmd.AddCommand(normalizeCmd())
	cmd.AddCommand(denormalizeCmd())

	return cmd
}
