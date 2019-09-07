package azuredevops

import (
	"runtime"

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
	if common.IsEnvironmentSetTo(env, "TF_BUILD", "true") {
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
	normalized = append(normalized, "NCI_SERVER_NAME="+common.GetEnvironment(env, "BUILD_REPOSITORY_PROVIDER"))
	normalized = append(normalized, "NCI_SERVER_HOST="+common.GetHostFromURL(common.GetEnvironment(env, "BUILD_REPOSITORY_URI")))
	normalized = append(normalized, "NCI_SERVER_VERSION=")

	// worker
	normalized = append(normalized, "NCI_WORKER_ID="+common.GetEnvironment(env, "AGENT_ID"))
	normalized = append(normalized, "NCI_WORKER_NAME="+common.GetEnvironment(env, "AGENT_MACHINENAME"))
	normalized = append(normalized, "NCI_WORKER_VERSION="+common.GetEnvironment(env, "AGENT_VERSION"))
	normalized = append(normalized, "NCI_WORKER_ARCH="+runtime.GOOS+"/"+runtime.GOARCH)

	// pipeline
	if common.GetEnvironment(env, "BUILD_REASON") == "Manual" {
		normalized = append(normalized, "NCI_PIPELINE_TRIGGER=manual")
	} else if common.GetEnvironment(env, "BUILD_REASON") == "IndividualCI" || common.GetEnvironment(env, "BUILD_REASON") == "BatchedCI" {
		normalized = append(normalized, "NCI_PIPELINE_TRIGGER=push")
	} else if common.GetEnvironment(env, "BUILD_REASON") == "Schedule" {
		normalized = append(normalized, "NCI_PIPELINE_TRIGGER=schedule")
	} else if common.GetEnvironment(env, "BUILD_REASON") == "PullRequest" {
		normalized = append(normalized, "NCI_PIPELINE_TRIGGER=pull_request")
	} else if common.GetEnvironment(env, "BUILD_REASON") == "BuildCompletion" {
		normalized = append(normalized, "NCI_PIPELINE_TRIGGER=build")
	} else {
		normalized = append(normalized, "NCI_PIPELINE_TRIGGER=unknown")
	}
	normalized = append(normalized, "NCI_PIPELINE_STAGE_NAME="+common.GetEnvironment(env, "SYSTEM_STAGENAME")) // SYSTEM_STAGEDISPLAYNAME
	normalized = append(normalized, "NCI_PIPELINE_STAGE_SLUG="+common.GetSlug(common.GetEnvironment(env, "SYSTEM_STAGENAME")))
	normalized = append(normalized, "NCI_PIPELINE_JOB_NAME="+common.GetEnvironment(env, "SYSTEM_JOBNAME")) // SYSTEM_JOBDISPLAYNAME
	normalized = append(normalized, "NCI_PIPELINE_JOB_SLUG="+common.GetSlug(common.GetEnvironment(env, "SYSTEM_JOBNAME")))

	// project
	normalized = append(normalized, "NCI_PROJECT_ID="+common.GetEnvironment(env, "SYSTEM_TEAMPROJECTID"))
	normalized = append(normalized, "NCI_PROJECT_NAME="+common.GetEnvironment(env, "SYSTEM_TEAMPROJECT"))
	normalized = append(normalized, "NCI_PROJECT_SLUG="+common.GetSlug(common.GetEnvironment(env, "SYSTEM_TEAMPROJECT")))
	normalized = append(normalized, "NCI_PROJECT_DIR="+common.GetGitDirectory())

	// repository
	normalized = append(normalized, common.GetSCMArguments(common.GetGitDirectory())...)

	return normalized
}

// NewNormalizer gets a instance of the normalizer
func NewNormalizer() Normalizer {
	entity := Normalizer{
		version: "0.1.0",
		name:    "Azure DevOps Pipeline",
		slug:    "azure-devops",
	}

	return entity
}
