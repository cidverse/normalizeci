package normalizeci

import (
	"fmt"
	"github.com/cidverse/normalizeci/pkg/common"
	"github.com/rs/zerolog/log"
	"os"
	"runtime"
	"strings"

	"github.com/cidverse/normalizeci/pkg/azuredevops"
	"github.com/cidverse/normalizeci/pkg/githubactions"
	"github.com/cidverse/normalizeci/pkg/gitlabci"
	"github.com/cidverse/normalizeci/pkg/localgit"
)

// holds all known normalizers
var normalizers []common.Normalizer

func init() {
	normalizers = append(normalizers, azuredevops.NewNormalizer())
	normalizers = append(normalizers, githubactions.NewNormalizer())
	normalizers = append(normalizers, gitlabci.NewNormalizer())
	normalizers = append(normalizers, localgit.NewNormalizer())
}

func RunDefaultNormalization() map[string]string {
	env := common.GetMachineEnvironment()
	return RunNormalization(env)
}

// RunNormalization executes the ci normalization for all supported services
func RunNormalization(env map[string]string) map[string]string {
	// normalize (iterate over all supported systems and normalize variables if possible)
	var normalized map[string]string
	for _, normalizer := range normalizers {
		if normalizer.Check(env) == true {
			log.Debug().Msg("Matched " + normalizer.GetName() + ", not checking for any other matches.")
			normalized = normalizer.Normalize(env)
			break
		} else {
			log.Debug().Msg("Didn't match in " + normalizer.GetName())
		}
	}

	return normalized
}

// RunDenormalization will generate ci variables for the target service
func RunDenormalization(target string, env map[string]string) map[string]string {
	// denormalize
	var normalized map[string]string
	for _, normalizer := range normalizers {
		if target == normalizer.GetSlug() {
			log.Debug().Msg("Matched " + normalizer.GetName() + ", not checking for any other matches.")
			normalized = normalizer.Denormalize(env)
			break
		} else {
			log.Debug().Msg("Didn't match in " + normalizer.GetName())
		}
	}

	return normalized
}

// FormatEnvironment makes the normalized environment available in the current session
func FormatEnvironment(normalized map[string]string, format string) string {
	if format == "export" {
		return setNormalizedEnvironmentExport(normalized)
	} else if format == "powershell" {
		return setNormalizedEnvironmentPowershell(normalized)
	} else if format == "cmd" {
		return setNormalizedEnvironmentCmd(normalized)
	}

	return ""
}

func GetDefaultFormat() string {
	if runtime.GOOS == "linux" {
		return "export"
	} else if runtime.GOOS == "windows" {
		return "powershell"
	}

	return ""
}

func SetProcessEnvironment(normalized map[string]string) {
	for key, element := range normalized {
		err := os.Setenv(key, element)
		if err != nil {
			log.Err(err).Str("key", key).Str("value", element).Msg("failed to set env property")
		}
	}
}

func setNormalizedEnvironmentExport(normalized map[string]string) string {
	var sb strings.Builder

	for key, element := range normalized {
		// print via stdout and escape values
		sb.WriteString(fmt.Sprintf("export %s=\"%s\"\n", key, strings.ReplaceAll(element, "\"", "\\\"")))
	}

	return sb.String()
}

func setNormalizedEnvironmentPowershell(normalized map[string]string) string {
	var sb strings.Builder

	for key, element := range normalized {
		// print via stdout and escape values
		sb.WriteString(fmt.Sprintf("$env:%s=\"%s\";\n", key, strings.ReplaceAll(element, "\"", "\\\"")))
	}

	return sb.String()
}

func setNormalizedEnvironmentCmd(normalized map[string]string) string {
	var sb strings.Builder

	for key, element := range normalized {
		// print via stdout and escape values
		sb.WriteString(fmt.Sprintf("set %s=%s\n", key, element))
	}

	return sb.String()
}
