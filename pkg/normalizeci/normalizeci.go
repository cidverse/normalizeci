package normalizeci

import (
	"fmt"
	"github.com/qubid/normalizeci/pkg/common"
	"github.com/rs/zerolog/log"
	"io"
	"os"
	"runtime"
	"strings"

	"github.com/qubid/normalizeci/pkg/azuredevops"
	"github.com/qubid/normalizeci/pkg/githubactions"
	"github.com/qubid/normalizeci/pkg/gitlabci"
	"github.com/qubid/normalizeci/pkg/localgit"
)

// RunNormalization executes the ci normalization for all supported services
func RunNormalization(env []string) []string {
	// initialize normalizers
	var normalizers []common.Normalizer
	normalizers = append(normalizers, azuredevops.NewNormalizer())
	normalizers = append(normalizers, githubactions.NewNormalizer())
	normalizers = append(normalizers, gitlabci.NewNormalizer())
	normalizers = append(normalizers, localgit.NewNormalizer())

	// normalize (iterate over all supported systems and normalize variables if possible)
	var normalized []string
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
func SetNormalizedEnvironment(normalized []string) {
	if runtime.GOOS == "linux" {
		setNormalizedEnvironmentLinux(normalized)
	} else if runtime.GOOS == "windows" {
		setNormalizedEnvironmentWindows(normalized)
	}
}

func setNormalizedEnvironmentLinux(normalized []string) {
	for _, entry := range normalized {
		entrySplit := strings.SplitN(entry, "=", 2)

		err := os.Setenv(entrySplit[0], entrySplit[1])
		common.CheckForError(err)

		// print via stdout and escape values
		s := fmt.Sprintf("export %s=\"%s\"\n", entrySplit[0], strings.ReplaceAll(entrySplit[1], "\"", "\\\""))
		io.WriteString(os.Stdout, s) // Ignoring error for simplicity.
	}
}

func setNormalizedEnvironmentWindows(normalized []string) {
	for _, entry := range normalized {
		entrySplit := strings.SplitN(entry, "=", 2)

		err := os.Setenv(entrySplit[0], entrySplit[1])
		common.CheckForError(err)

		// print via stdout and escape values
		s := fmt.Sprintf("Set-Variable -Name %s -Value \"%s\";\n", entrySplit[0], strings.ReplaceAll(entrySplit[1], "\"", "\\\""))
		io.WriteString(os.Stdout, s) // Ignoring error for simplicity.
	}
}