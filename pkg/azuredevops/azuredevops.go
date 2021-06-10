package azuredevops

import (
	"github.com/cidverse/normalizeci/pkg/vcsrepository"
	"github.com/gosimple/slug"
	"runtime"

	"github.com/cidverse/normalizeci/pkg/common"
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
func (n Normalizer) Check(env map[string]string) bool {
	if env["TF_BUILD"] == "True" {
		return true
	}

	return false
}

// Normalize normalizes the environment variables into the common format
func (n Normalizer) Normalize(env map[string]string) map[string]string {
	data := make(map[string]string)

	// common
	data["NCI"] = "true"
	data["NCI_VERSION"] = n.version
	data["NCI_SERVICE_NAME"] = n.name
	data["NCI_SERVICE_SLUG"] = n.slug

	// server
	data["NCI_SERVER_NAME"] = env["BUILD_REPOSITORY_PROVIDER"]
	data["NCI_SERVER_HOST"] = common.GetHostFromURL(env["BUILD_REPOSITORY_URI"])
	data["NCI_SERVER_VERSION"] = ""

	// worker
	data["NCI_WORKER_ID"] = env["AGENT_ID"]
	data["NCI_WORKER_NAME"] = env["AGENT_MACHINENAME"]
	data["NCI_WORKER_VERSION"] = env["AGENT_VERSION"]
	data["NCI_WORKER_ARCH"] = runtime.GOOS + "/" + runtime.GOARCH

	// pipeline
	if env["BUILD_REASON"] == "Manual" {
		data["NCI_PIPELINE_TRIGGER"] = "manual"
	} else if env["BUILD_REASON"] == "IndividualCI" || env["BUILD_REASON"] == "BatchedCI" {
		data["NCI_PIPELINE_TRIGGER"] = "push"
	} else if env["BUILD_REASON"] == "Schedule" {
		data["NCI_PIPELINE_TRIGGER"] = "schedule"
	} else if env["BUILD_REASON"] == "PullRequest" {
		data["NCI_PIPELINE_TRIGGER"] = "pull_request"
	} else if env["BUILD_REASON"] == "BuildCompletion" {
		data["NCI_PIPELINE_TRIGGER"] = "build"
	} else {
		data["NCI_PIPELINE_TRIGGER"] = "unknown"
	}
	data["NCI_PIPELINE_STAGE_NAME"] = env["SYSTEM_STAGENAME"] // SYSTEM_STAGEDISPLAYNAME
	data["NCI_PIPELINE_STAGE_SLUG"] = slug.Make(env["SYSTEM_STAGENAME"])
	data["NCI_PIPELINE_JOB_NAME"] = env["SYSTEM_JOBNAME"] // SYSTEM_JOBDISPLAYNAME
	data["NCI_PIPELINE_JOB_SLUG"] = slug.Make(env["SYSTEM_JOBNAME"])

	// project
	data["NCI_PROJECT_ID"] = env["SYSTEM_TEAMPROJECTID"]
	data["NCI_PROJECT_NAME"] = env["SYSTEM_TEAMPROJECT"]
	data["NCI_PROJECT_SLUG"] = slug.Make(env["SYSTEM_TEAMPROJECT"])
	data["NCI_PROJECT_DIR"] = vcsrepository.FindRepositoryDirectory(common.GetWorkingDirectory())

	// repository
	addData, addDataErr := vcsrepository.GetVCSRepositoryInformation(data["NCI_PROJECT_DIR"])
	if addDataErr != nil {
		panic(addDataErr)
	}
	for addKey, addElement := range addData {
		data[addKey] = addElement
	}

	return data
}

// NewNormalizer gets a instance of the normalizer
func NewNormalizer() Normalizer {
	entity := Normalizer{
		version: "0.2.0",
		name:    "Azure DevOps Pipeline",
		slug:    "azure-devops",
	}

	return entity
}
