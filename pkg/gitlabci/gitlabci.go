package gitlabci

import (
	"runtime"
	"strings"
	
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
	if common.IsEnvironmentSetTo(env, "GITLAB_CI", "true") {
		return true
	}

	return false
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
	normalized = append(normalized, "NCI_SERVER_NAME="+common.GetEnvironment(env, "CI_SERVER_NAME"))
	normalized = append(normalized, "NCI_SERVER_HOST="+common.GetEnvironment(env, "CI_SERVER_HOST"))
	normalized = append(normalized, "NCI_SERVER_VERSION="+common.GetEnvironment(env, "CI_SERVER_VERSION"))

	// worker
	normalized = append(normalized, "NCI_WORKER_ID="+common.GetEnvironment(env, "CI_RUNNER_ID"))
	normalized = append(normalized, "NCI_WORKER_NAME="+common.GetEnvironment(env, "CI_RUNNER_DESCRIPTION"))
	normalized = append(normalized, "NCI_WORKER_VERSION="+common.GetEnvironment(env, "CI_RUNNER_VERSION"))
	normalized = append(normalized, "NCI_WORKER_ARCH="+runtime.GOOS+"/"+runtime.GOARCH)

	// pipeline
	normalized = append(normalized, "NCI_PIPELINE_TRIGGER="+common.GetEnvironment(env, "CI_PIPELINE_SOURCE"))
	if common.GetEnvironment(normalized, "NCI_PIPELINE_TRIGGER") == "pull_request" {
		// PR
		normalized = append(normalized, "NCI_PIPELINE_PULL_REQUEST_ID="+common.GetEnvironment(env, "CI_MERGE_REQUEST_IID"))
	}
	normalized = append(normalized, "NCI_PIPELINE_STAGE_NAME="+common.GetEnvironment(env, "CI_JOB_STAGE"))
	normalized = append(normalized, "NCI_PIPELINE_STAGE_SLUG="+common.GetSlug(common.GetEnvironment(env, "CI_JOB_STAGE")))
	normalized = append(normalized, "NCI_PIPELINE_JOB_NAME="+common.GetEnvironment(env, "CI_JOB_NAME"))
	normalized = append(normalized, "NCI_PIPELINE_JOB_SLUG="+common.GetSlug(common.GetEnvironment(env, "CI_JOB_NAME")))

	// container registry
	normalized = append(normalized, "NCI_CONTAINERREGISTRY_HOST="+common.GetEnvironmentOrDefault(env, "NCI_CONTAINERREGISTRY_HOST", common.GetEnvironment(env, "CI_REGISTRY")))
	normalized = append(normalized, "NCI_CONTAINERREGISTRY_REPOSITORY="+common.GetEnvironmentOrDefault(env, "NCI_CONTAINERREGISTRY_REPOSITORY", common.GetEnvironment(env, "CI_REGISTRY_IMAGE")))
	if common.HasEnvironment(env, "CI_DEPLOY_USER") {
		normalized = append(normalized, "NCI_CONTAINERREGISTRY_USERNAME="+common.GetEnvironmentOrDefault(env, "NCI_CONTAINERREGISTRY_USERNAME", common.GetEnvironment(env, "CI_DEPLOY_USER")))
		normalized = append(normalized, "NCI_CONTAINERREGISTRY_PASSWORD="+common.GetEnvironmentOrDefault(env, "NCI_CONTAINERREGISTRY_PASSWORD", common.GetEnvironment(env, "CI_DEPLOY_PASSWORD")))
	} else {
		normalized = append(normalized, "NCI_CONTAINERREGISTRY_USERNAME="+common.GetEnvironmentOrDefault(env, "NCI_CONTAINERREGISTRY_USERNAME", common.GetEnvironment(env, "CI_REGISTRY_USER")))
		normalized = append(normalized, "NCI_CONTAINERREGISTRY_PASSWORD="+common.GetEnvironmentOrDefault(env, "NCI_CONTAINERREGISTRY_PASSWORD", common.GetEnvironment(env, "CI_REGISTRY_PASSWORD")))
	}

	// project
	normalized = append(normalized, "NCI_PROJECT_ID="+common.GetEnvironment(env, "CI_PROJECT_ID"))
	normalized = append(normalized, "NCI_PROJECT_NAME="+common.GetEnvironment(env, "CI_PROJECT_NAME"))
	normalized = append(normalized, "NCI_PROJECT_SLUG="+common.GetSlug(common.GetEnvironment(env, "CI_PROJECT_PATH")))
	normalized = append(normalized, "NCI_PROJECT_DIR="+common.GetGitDirectory())

	// repository
	if common.HasEnvironment(env, "CI_COMMIT_TAG") {
		normalized = append(normalized, "NCI_COMMIT_REF_TYPE=tag")
	} else {
		normalized = append(normalized, "NCI_COMMIT_REF_TYPE=branch")
	}
	normalized = append(normalized, "NCI_COMMIT_REF_NAME="+common.GetEnvironment(env, "CI_COMMIT_REF_NAME"))
	normalized = append(normalized, "NCI_COMMIT_REF_SLUG="+common.GetSlug(common.GetEnvironment(env, "CI_COMMIT_REF_NAME")))
	normalized = append(normalized, "NCI_COMMIT_REF_RELEASE="+strings.TrimLeft(common.GetSlug(common.GetEnvironment(env, "CI_COMMIT_REF_NAME")), "v"))
	normalized = append(normalized, "NCI_COMMIT_TITLE="+common.GetEnvironment(env, "CI_COMMIT_TITLE"))
	normalized = append(normalized, "NCI_COMMIT_DESCRIPTION="+common.GetEnvironment(env, "CI_COMMIT_DESCRIPTION"))
	normalized = append(normalized, "NCI_COMMIT_SHA="+common.GetEnvironment(env, "CI_COMMIT_SHA"))
	normalized = append(normalized, "NCI_COMMIT_SHA_SHORT="+common.GetEnvironment(env, "CI_COMMIT_SHORT_SHA"))

	return normalized
}

// NewNormalizer gets a instance of the normalizer
func NewNormalizer() Normalizer {
	entity := Normalizer{
		version: "0.1.0",
		name:    "GitLab CI",
		slug:    "gitlab-ci",
	}

	return entity
}
