package normalizeci

import (
	"fmt"
	"github.com/cidverse/normalizeci/pkg/common"
	"github.com/rs/zerolog/log"
	"io"
	"os"
	"runtime"
	"strings"

	"github.com/cidverse/normalizeci/pkg/azuredevops"
	"github.com/cidverse/normalizeci/pkg/githubactions"
	"github.com/cidverse/normalizeci/pkg/gitlabci"
	"github.com/cidverse/normalizeci/pkg/localgit"
)

func RunDefaultNormalization() map[string]string {
	env := common.GetMachineEnvironment()
	return RunNormalization(env)
}

// RunNormalization executes the ci normalization for all supported services
func RunNormalization(env map[string]string) map[string]string {
	// initialize normalizers
	var normalizers []common.Normalizer
	normalizers = append(normalizers, azuredevops.NewNormalizer())
	normalizers = append(normalizers, githubactions.NewNormalizer())
	normalizers = append(normalizers, gitlabci.NewNormalizer())
	normalizers = append(normalizers, localgit.NewNormalizer())

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

// SetNormalizedEnvironment makes the normalized environment available in the current session
func SetNormalizedEnvironment(normalized map[string]string) {
	if runtime.GOOS == "linux" {
		setNormalizedEnvironmentLinux(normalized)
	} else if runtime.GOOS == "windows" {
		setNormalizedEnvironmentWindows(normalized)
	}
}

func setNormalizedEnvironmentLinux(normalized map[string]string) {
	for key, element := range normalized {
		err := os.Setenv(key, element)
		common.CheckForError(err)

		// print via stdout and escape values
		s := fmt.Sprintf("export %s=\"%s\"\n", key, strings.ReplaceAll(element, "\"", "\\\""))
		io.WriteString(os.Stdout, s) // Ignoring error for simplicity.
	}
}

func setNormalizedEnvironmentWindows(normalized map[string]string) {
	for key, element := range normalized {
		err := os.Setenv(key, element)
		common.CheckForError(err)

		// print via stdout and escape values
		s := fmt.Sprintf("Set-Variable -Name %s -Value \"%s\";\n", key, strings.ReplaceAll(element, "\"", "\\\""))
		io.WriteString(os.Stdout, s) // Ignoring error for simplicity.
	}
}
