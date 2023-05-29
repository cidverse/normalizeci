package localgit

import (
	"runtime"
	"time"

	"github.com/cidverse/go-vcs/vcsutil"
	"github.com/cidverse/normalizeci/pkg/ncispec/common"
	v1 "github.com/cidverse/normalizeci/pkg/ncispec/v1"
	"github.com/cidverse/normalizeci/pkg/nciutil"
	"github.com/cidverse/normalizeci/pkg/projectdetails"
	"github.com/cidverse/normalizeci/pkg/vcsrepository"
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

// GetSlug returns the slug of the normalizer
func (n Normalizer) GetSlug() string {
	return n.slug
}

// Check if this package can handle the current environment
func (n Normalizer) Check(env map[string]string) bool {
	return true
}

// Normalize normalizes the environment variables into the common format
func (n Normalizer) Normalize(env map[string]string) v1.Spec {
	nci := v1.Create(n.name, n.slug)

	// worker
	nci.Worker.Id = "local"
	nci.Worker.Name = "localhost"
	nci.Worker.Type = "local"
	nci.Worker.OS = runtime.GOOS
	nci.Worker.Version = "1.0.0"
	nci.Worker.Arch = runtime.GOOS + "/" + runtime.GOARCH

	// pipeline
	nci.Pipeline.Id = nciutil.GenerateSnowflakeId()
	nci.Pipeline.Trigger = common.PipelineTriggerCLI
	nci.Pipeline.StageId = nciutil.GenerateSnowflakeId()
	nci.Pipeline.StageName = common.PipelineStageDefault
	nci.Pipeline.StageSlug = common.PipelineStageDefault
	nci.Pipeline.JobId = nciutil.GenerateSnowflakeId()
	nci.Pipeline.JobName = common.PipelineJobDefault
	nci.Pipeline.JobSlug = common.PipelineJobDefault
	nci.Pipeline.JobStartedAt = time.Now().Format(time.RFC3339)
	nci.Pipeline.Attempt = "1"

	// repository
	projectDir, _ := vcsutil.FindProjectDirectory()
	vcsData, addDataErr := vcsrepository.GetVCSRepositoryInformation(projectDir)
	if addDataErr != nil {
		panic(addDataErr)
	}
	nci.Repository = vcsData.Repository
	nci.Commit = vcsData.Commit

	// project details
	projectData, _ := projectdetails.GetProjectDetails(nci.Repository.Kind, nci.Repository.Remote, nci.Repository.HostType, nci.Repository.HostServer)
	nci.Project = projectData
	nci.Project.Dir = projectDir

	// flags
	nci.Flags.DeployFreeze = "false"

	return nci
}

func (n Normalizer) Denormalize(spec v1.Spec) map[string]string {
	return make(map[string]string)
}

// NewNormalizer gets an instance of the normalizer
func NewNormalizer() Normalizer {
	entity := Normalizer{
		version: "0.3.0",
		name:    "Local Git Repository",
		slug:    "local-git",
	}

	return entity
}
