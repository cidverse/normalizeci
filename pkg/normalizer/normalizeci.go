package normalizer

import (
	"errors"
	"fmt"
	"log/slog"
	"os"
	"runtime"
	"strings"

	"github.com/cidverse/go-vcs/vcsutil"
	v1 "github.com/cidverse/normalizeci/pkg/ncispec/v1"
	"github.com/cidverse/normalizeci/pkg/normalizer/api"
	"github.com/cidverse/normalizeci/pkg/normalizer/appveyor"
	"github.com/cidverse/normalizeci/pkg/normalizer/azuredevops"
	"github.com/cidverse/normalizeci/pkg/normalizer/circleci"
	"github.com/cidverse/normalizeci/pkg/normalizer/githubactions"
	"github.com/cidverse/normalizeci/pkg/normalizer/gitlabci"
	"github.com/cidverse/normalizeci/pkg/normalizer/localgit"
)

type Options struct {
	Env        map[string]string
	ProjectDir string
}

func GetNormalizers(opts Options) ([]api.Normalizer, error) {
	// try to find the project directory if not set
	if opts.ProjectDir != "" {
		projectDir, err := vcsutil.FindProjectDirectoryFromWorkDir()
		if err != nil {
			return nil, fmt.Errorf("failed to find project directory: %v", err)
		}
		slog.With("dir", projectDir).Debug("Automatically detected and set project directory")

		opts.ProjectDir = projectDir
	}

	return []api.Normalizer{
		appveyor.NewNormalizer(),
		azuredevops.NewNormalizer(),
		circleci.NewNormalizer(),
		githubactions.NewNormalizer(),
		gitlabci.NewNormalizer(),
		localgit.NewNormalizer(opts.ProjectDir),
	}, nil
}

func Normalize() (v1.Spec, error) {
	return NormalizeEnv(Options{})
}

// NormalizeEnv executes the ci normalization for all supported services
func NormalizeEnv(opts Options) (v1.Spec, error) {
	normalizers, err := GetNormalizers(opts)
	if err != nil {
		return v1.Spec{}, err
	}

	// default to machine environment
	if opts.Env == nil {
		opts.Env = api.GetMachineEnvironment()
	}

	// normalize (iterate over all supported systems and normalize variables if possible)
	for _, normalizer := range normalizers {
		if normalizer.Check(opts.Env) {
			slog.With("normalizer", normalizer.GetName()).With("dir", opts.ProjectDir).Debug("Matched normalizer, not checking for any other matches.")
			return normalizer.Normalize(opts.Env)
		} else {
			slog.With("normalizer", normalizer.GetName()).With("dir", opts.ProjectDir).Debug("Skipping normalizer, check failed")
		}
	}

	return v1.Spec{}, errors.New("no matching normalizer found")
}

// Denormalize will generate ci variables for the target service
func Denormalize(opts Options, target string, env v1.Spec) (map[string]string, error) {
	normalizers, err := GetNormalizers(opts)
	if err != nil {
		return nil, err
	}

	// default to machine environment
	if opts.Env == nil {
		opts.Env = api.GetMachineEnvironment()
	}

	// denormalize
	for _, normalizer := range normalizers {
		if target == normalizer.GetSlug() {
			slog.With("normalizer", normalizer.GetName()).With("dir", opts.ProjectDir).Debug("Matched normalizer, not checking for any other matches.")
			return normalizer.Denormalize(env)
		} else {
			slog.With("normalizer", normalizer.GetName()).With("dir", opts.ProjectDir).Debug("Skipping normalizer, check failed")
		}
	}

	return nil, errors.New("no matching denormalizer found")
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
			slog.With("key", key).With("value", element).Error("failed to set env property")
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
