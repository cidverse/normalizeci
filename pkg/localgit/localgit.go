package localgit

import (
	"github.com/cidverse/normalizeci/pkg/common"
	"github.com/cidverse/normalizeci/pkg/vcsrepository"
	"github.com/gosimple/slug"
	"os"
	"path/filepath"
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
	return true
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
	data["NCI_SERVER_NAME"] = "local"
	data["NCI_SERVER_HOST"] = "localhost"
	data["NCI_SERVER_VERSION"] = ""

	// worker
	data["NCI_WORKER_ID"] = "local"
	data["NCI_WORKER_NAME"] = ""
	data["NCI_WORKER_VERSION"] = ""
	data["NCI_WORKER_ARCH"] = runtime.GOOS + "/" + runtime.GOARCH

	// pipeline
	data["NCI_PIPELINE_TRIGGER"] = "manual"
	data["NCI_PIPELINE_STAGE_NAME"] = ""
	data["NCI_PIPELINE_STAGE_SLUG"] = ""
	data["NCI_PIPELINE_JOB_NAME"] = ""
	data["NCI_PIPELINE_JOB_SLUG"] = ""

	// container registry
	data["NCI_CONTAINERREGISTRY_HOST"] = ""
	data["NCI_CONTAINERREGISTRY_REPOSITORY"] = slug.Make(common.GetDirectoryNameFromPath(filepath.Join(vcsrepository.FindRepositoryDirectory(common.GetWorkingDirectory())+string(os.PathSeparator), "file")))
	data["NCI_CONTAINERREGISTRY_USERNAME"] = ""
	data["NCI_CONTAINERREGISTRY_PASSWORD"] = ""

	// project
	data["NCI_PROJECT_ID"] = ""
	data["NCI_PROJECT_NAME"] = ""
	data["NCI_PROJECT_SLUG"] = ""
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
		name:    "Local Git Repository",
		slug:    "local-git",
	}

	return entity
}
