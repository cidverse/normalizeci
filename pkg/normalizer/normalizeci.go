package normalizer

import (
	"errors"
	"fmt"
	"os"
	"runtime"
	"strings"

	v1 "github.com/cidverse/normalizeci/pkg/ncispec/v1"
	"github.com/cidverse/normalizeci/pkg/normalizer/api"
	"github.com/cidverse/normalizeci/pkg/normalizer/azuredevops"
	"github.com/cidverse/normalizeci/pkg/normalizer/githubactions"
	"github.com/cidverse/normalizeci/pkg/normalizer/gitlabci"
	"github.com/cidverse/normalizeci/pkg/normalizer/localgit"
	"github.com/rs/zerolog/log"
)

// holds all known normalizers
var normalizers []api.Normalizer

func init() {
	normalizers = append(normalizers, azuredevops.NewNormalizer())
	normalizers = append(normalizers, githubactions.NewNormalizer())
	normalizers = append(normalizers, gitlabci.NewNormalizer())
	normalizers = append(normalizers, localgit.NewNormalizer())
}

func Normalize() v1.Spec {
	env := api.GetMachineEnvironment()
	return NormalizeEnv(env)
}

// NormalizeEnv executes the ci normalization for all supported services
func NormalizeEnv(env map[string]string) v1.Spec {
	// normalize (iterate over all supported systems and normalize variables if possible)
	var normalized v1.Spec
	for _, normalizer := range normalizers {
		if normalizer.Check(env) {
			log.Debug().Msg("Matched " + normalizer.GetName() + ", not checking for any other matches.")
			normalized = normalizer.Normalize(env)
			break
		} else {
			log.Debug().Msg("Didn't match in " + normalizer.GetName())
		}
	}

	return normalized
}

// Denormalize will generate ci variables for the target service
func Denormalize(target string, env v1.Spec) map[string]string {
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
func FormatEnvironment(normalized map[string]string, format string) (string, error) {
	if format == "export" {
		return setNormalizedEnvironmentExport(normalized), nil
	} else if format == "powershell" {
		return setNormalizedEnvironmentPowershell(normalized), nil
	} else if format == "cmd" {
		return setNormalizedEnvironmentCmd(normalized), nil
	}

	return "", errors.New("unsupported format: " + format)
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
