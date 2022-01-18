package githubactions

import (
	"github.com/cidverse/normalizeci/pkg/projectdetails"
	"github.com/cidverse/normalizeci/pkg/vcsrepository"
	"github.com/gosimple/slug"
	"os"
	"path/filepath"
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
	if env["GITHUB_ACTIONS"] == "true" {
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
	data["NCI_WORKER_ID"] = env["RUNNER_TRACKING_ID"]
	data["NCI_WORKER_NAME"] = env["RUNNER_TRACKING_ID"]
	data["NCI_WORKER_VERSION"] = env["ImageVersion"]
	data["NCI_WORKER_ARCH"] = runtime.GOOS + "/" + runtime.GOARCH

	// pipeline
	pipelineEvent := env["GITHUB_EVENT_NAME"]
	switch pipelineEvent {
	case "push":
		data["NCI_PIPELINE_TRIGGER"] = "push"
	case "pull_request":
		data["NCI_PIPELINE_TRIGGER"] = "pull_request"
	default:
		data["NCI_PIPELINE_TRIGGER"] = "unknown"
	}
	if env["NCI_PIPELINE_TRIGGER"] == "pull_request" {
		// PR
		data["NCI_PIPELINE_PULL_REQUEST_ID"] = "unknown" // not supported by GH yet.
	}
	data["NCI_PIPELINE_STAGE_NAME"] = env["GITHUB_WORKFLOW"]
	data["NCI_PIPELINE_STAGE_SLUG"] = slug.Make(env["GITHUB_WORKFLOW"])
	data["NCI_PIPELINE_JOB_NAME"] = env["GITHUB_ACTION"]
	data["NCI_PIPELINE_JOB_SLUG"] = slug.Make(env["GITHUB_ACTION"])

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
	data["NCI_CONTAINERREGISTRY_HOST"] = ""
	data["NCI_CONTAINERREGISTRY_REPOSITORY"] = slug.Make(common.GetDirectoryNameFromPath(filepath.Join(vcsrepository.FindRepositoryDirectory(common.GetWorkingDirectory())+string(os.PathSeparator), "file")))
	data["NCI_CONTAINERREGISTRY_USERNAME"] = ""
	data["NCI_CONTAINERREGISTRY_PASSWORD"] = ""
	data["NCI_CONTAINERREGISTRY_TAG"] = data["NCI_COMMIT_REF_RELEASE"]

	return data
}

func (n Normalizer) Denormalize(env map[string]string) map[string]string {
	return make(map[string]string)
}

// NewNormalizer gets a instance of the normalizer
func NewNormalizer() Normalizer {
	entity := Normalizer{
		version: "0.3.0",
		name:    "GitHub Actions",
		slug:    "github-actions",
	}

	return entity
}
