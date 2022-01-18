package gitlabci

import (
	"github.com/cidverse/normalizeci/pkg/common"
	"github.com/cidverse/normalizeci/pkg/projectdetails"
	"github.com/cidverse/normalizeci/pkg/vcsrepository"
	"github.com/gosimple/slug"
	"runtime"
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
	if env["GITLAB_CI"] == "true" {
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

	// worker
	data["NCI_WORKER_ID"] = env["CI_RUNNER_ID"]
	data["NCI_WORKER_NAME"] = env["CI_RUNNER_DESCRIPTION"]
	data["NCI_WORKER_VERSION"] = env["CI_RUNNER_VERSION"]
	data["NCI_WORKER_ARCH"] = runtime.GOOS + "/" + runtime.GOARCH

	// pipeline
	data["NCI_PIPELINE_TRIGGER"] = env["CI_PIPELINE_SOURCE"]
	if env["NCI_PIPELINE_TRIGGER"] == "pull_request" {
		// PR
		data["NCI_PIPELINE_PULL_REQUEST_ID"] = env["CI_MERGE_REQUEST_IID"]
	}

	data["NCI_PIPELINE_STAGE_NAME"] = env["CI_JOB_STAGE"]
	data["NCI_PIPELINE_STAGE_SLUG"] = slug.Make(env["CI_JOB_STAGE"])
	data["NCI_PIPELINE_JOB_NAME"] = env["CI_JOB_NAME"]
	data["NCI_PIPELINE_JOB_SLUG"] = slug.Make(env["CI_JOB_NAME"])

	// repository
	projectDir := vcsrepository.FindRepositoryDirectory(common.GetWorkingDirectory())
	addData, addDataErr := vcsrepository.GetVCSRepositoryInformation(projectDir)
	if addDataErr != nil {
		panic(addDataErr)
	}
	for addKey, addElement := range addData {
		data[addKey] = addElement
	}

	// project details
	projectData := projectdetails.GetProjectDetails(data["NCI_REPOSITORY_KIND"], data["NCI_REPOSITORY_REMOTE"])
	if projectData != nil {
		for addKey, addElement := range projectData {
			data[addKey] = addElement
		}
	}
	data["NCI_PROJECT_DIR"] = projectDir

	// container registry
	data["NCI_CONTAINERREGISTRY_HOST"] = env["CI_REGISTRY"]
	data["NCI_CONTAINERREGISTRY_REPOSITORY"] = env["CI_REGISTRY_IMAGE"]
	if len(env["CI_DEPLOY_USER"]) > 0 {
		data["NCI_CONTAINERREGISTRY_USERNAME"] = env["CI_DEPLOY_USER"]
		data["NCI_CONTAINERREGISTRY_PASSWORD"] = env["CI_DEPLOY_PASSWORD"]
	} else {
		data["NCI_CONTAINERREGISTRY_USERNAME"] = env["CI_REGISTRY_USER"]
		data["NCI_CONTAINERREGISTRY_PASSWORD"] = env["CI_REGISTRY_PASSWORD"]
	}
	data["NCI_CONTAINERREGISTRY_TAG"] = data["NCI_COMMIT_REF_RELEASE"]

	// control
	data["NCI_DEPLOY_FREEZE"] = env["CI_DEPLOY_FREEZE"]

	return data
}

func (n Normalizer) Denormalize(env map[string]string) map[string]string {
	data := make(map[string]string)

	data["GITLAB_CI"] = "true"

	return data
}

// NewNormalizer gets a instance of the normalizer
func NewNormalizer() Normalizer {
	entity := Normalizer{
		version: "0.3.0",
		name:    "GitLab CI",
		slug:    "gitlab-ci",
	}

	return entity
}
