package common

import (
	"github.com/rs/zerolog/log"
	"net/url"
	"os"
	"path/filepath"
	"strings"
)

// Normalizer is a common interface to work with all normalizers
type Normalizer interface {
	GetName() string
	GetSlug() string
	Check(env map[string]string) bool
	Normalize(env map[string]string) map[string]string
	Denormalize(env map[string]string) map[string]string
}

// GetMachineEnvironment returns a map with all environment variables set on the machine
func GetMachineEnvironment() map[string]string {
	data := make(map[string]string)

	for _, entry := range os.Environ() {
		z := strings.SplitN(entry, "=", 2)

		if len(z[0]) > 0 {
			data[z[0]] = z[1]
		}
	}

	return data
}

// GetEnvironmentFrom returns a map with all environment variables contained in env
func GetEnvironmentFrom(env []string) map[string]string {
	data := make(map[string]string)

	for _, entry := range env {
		z := strings.SplitN(entry, "=", 2)
		data[z[0]] = z[1]
	}

	return data
}

// GetWorkingDirectory returns the current working directory
func GetWorkingDirectory() string {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	return dir
}

// GetDirectoryNameFromPath gets the directory name from a path
func GetDirectoryNameFromPath(path string) string {
	dir := filepath.Dir(path)
	parent := filepath.Base(dir)

	return parent
}

// GetHostFromURL gets the host from a url
func GetHostFromURL(addr string) string {
	u, err := url.Parse(addr)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to get host part from url")
	}

	return u.Host
}
