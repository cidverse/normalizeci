package api

import (
	"fmt"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	v1 "github.com/cidverse/normalizeci/pkg/ncispec/v1"
)

// Normalizer is a common interface to work with all normalizers
type Normalizer interface {
	GetName() string
	GetSlug() string
	Check(env map[string]string) bool
	Normalize(env map[string]string) (v1.Spec, error)
	Denormalize(spec v1.Spec) (map[string]string, error)
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

// GetEnvironmentMerge returns a map with all environment variables contained in env
func GetEnvironmentMerge(env []string, overwrite []string) map[string]string {
	data := make(map[string]string)

	for _, entry := range env {
		z := strings.SplitN(entry, "=", 2)
		data[z[0]] = z[1]
	}

	for _, entry := range overwrite {
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
func GetHostFromURL(addr string) (string, error) {
	u, err := url.Parse(addr)
	if err != nil {
		return "", fmt.Errorf("failed to parse URL: %w", err)
	}

	return u.Host, nil
}

func ToEnvName(input string) string {
	return strings.Replace(strings.ToUpper(input), ".", "_", -1)
}

func GetProjectUID(repository v1.Repository, project v1.Project) string {
	if os.Getenv("NCI_PROJECT_UID_TEMPLATE") != "" {
		input := os.Getenv("NCI_PROJECT_UID_TEMPLATE")

		input = strings.ReplaceAll(input, "{{NCI_REPOSITORY_HOST_SERVER}}", repository.HostServer)
		input = strings.ReplaceAll(input, "{{NCI_REPOSITORY_HOST_SERVER_SLUG}}", repository.HostServerSlug)
		input = strings.ReplaceAll(input, "{{NCI_PROJECT_ID}}", project.ID)
		input = strings.ReplaceAll(input, "{{NCI_PROJECT_SLUG}}", project.Slug)

		return input
	}

	if repository.HostServerSlug == "" || project.ID == "" {
		return ""
	}

	return fmt.Sprintf("%s-%s", repository.HostServerSlug, project.ID)
}
