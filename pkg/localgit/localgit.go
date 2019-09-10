package localgit

import (
	"runtime"
	"strings"
	"os"

	"github.com/PhilippHeuer/normalize-ci/pkg/common"
)

// Normalizer is the implementation of the normalizer
type Normalizer struct {
	version string
	name    string
	slug    string
}

// GetName returns the name of the normalizer
func (n Normalizer) GetName() string {
	return n.name
}

// Check if this package can handle the current environment
func (n Normalizer) Check(env []string) bool {
	return true
}

// Normalize normalizes the environment variables into the common format
func (n Normalizer) Normalize(env []string) []string {
	var normalized []string

	// common
	normalized = append(normalized, "NCI=true")
	normalized = append(normalized, "NCI_VERSION="+n.version)
	normalized = append(normalized, "NCI_SERVICE_NAME="+n.name)
	normalized = append(normalized, "NCI_SERVICE_SLUG="+n.slug)

	// server
	normalized = append(normalized, "NCI_SERVER_NAME=local")
	normalized = append(normalized, "NCI_SERVER_HOST=localhost")
	normalized = append(normalized, "NCI_SERVER_VERSION=")

	// worker
	normalized = append(normalized, "NCI_WORKER_ID=local")
	normalized = append(normalized, "NCI_WORKER_NAME=")
	normalized = append(normalized, "NCI_WORKER_VERSION=")
	normalized = append(normalized, "NCI_WORKER_ARCH="+runtime.GOOS+"/"+runtime.GOARCH)

	// pipeline
	normalized = append(normalized, "NCI_PIPELINE_TRIGGER=manual")
	normalized = append(normalized, "NCI_PIPELINE_STAGE_NAME=")
	normalized = append(normalized, "NCI_PIPELINE_STAGE_SLUG=")
	normalized = append(normalized, "NCI_PIPELINE_JOB_NAME=")
	normalized = append(normalized, "NCI_PIPELINE_JOB_SLUG=")

	// container registry
	normalized = append(normalized, "NCI_CONTAINERREGISTRY_HOST="+common.GetEnvironment(env, "NCI_CONTAINERREGISTRY_HOST"))
	normalized = append(normalized, "NCI_CONTAINERREGISTRY_REPOSITORY="+common.GetEnvironmentOrDefault(env, "NCI_CONTAINERREGISTRY_REPOSITORY", strings.ToLower(common.GetDirectoryNameFromPath(common.GetGitDirectory()+string(os.PathSeparator)+".git"))))
	normalized = append(normalized, "NCI_CONTAINERREGISTRY_USERNAME="+common.GetEnvironment(env, "NCI_CONTAINERREGISTRY_USERNAME"))
	normalized = append(normalized, "NCI_CONTAINERREGISTRY_PASSWORD="+common.GetEnvironment(env, "NCI_CONTAINERREGISTRY_PASSWORD"))

	// project
	normalized = append(normalized, "NCI_PROJECT_ID=")
	normalized = append(normalized, "NCI_PROJECT_NAME=")
	normalized = append(normalized, "NCI_PROJECT_SLUG=")
	normalized = append(normalized, "NCI_PROJECT_DIR="+common.GetGitDirectory())

	// repository
	normalized = append(normalized, common.GetSCMArguments(common.GetGitDirectory())...)

	return normalized
}

// NewNormalizer gets a instance of the normalizer
func NewNormalizer() Normalizer {
	entity := Normalizer{
		version: "0.1.0",
		name:    "Local Git Repository",
		slug:    "local-git",
	}

	return entity
}
