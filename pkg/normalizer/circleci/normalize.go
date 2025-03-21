package circleci

import (
	"fmt"
	"runtime"
	"time"

	"github.com/cidverse/go-vcs/vcsutil"
	"github.com/cidverse/normalizeci/pkg/ncispec/common"
	v1 "github.com/cidverse/normalizeci/pkg/ncispec/v1"
	"github.com/cidverse/normalizeci/pkg/normalizer/api"
	"github.com/cidverse/normalizeci/pkg/projectdetails"
	"github.com/cidverse/normalizeci/pkg/vcsrepository"
	"github.com/gosimple/slug"
)

// Normalize normalizes the environment variables into the common format
func (n Normalizer) Normalize(env map[string]string) (v1.Spec, error) {
	nci := v1.Create(n.name, n.slug)

	// worker
	nci.Worker = v1.Worker{
		Id:      env["CIRCLE_NODE_INDEX"],
		Name:    env["CIRCLE_NODE_INDEX"],
		Type:    "circleci_hosted_vm",
		OS:      "unknown",
		Version: "latest",
		Arch:    runtime.GOOS + "/" + runtime.GOARCH,
	}

	// pipeline
	nci.Pipeline.Id = env["CIRCLE_PIPELINE_ID"]
	nci.Pipeline.Trigger = common.PipelineTriggerUnknown
	nci.Pipeline.StageId = env["CIRCLE_WORKFLOW_ID"]
	nci.Pipeline.StageName = "default"
	nci.Pipeline.StageSlug = slug.Make("default")
	nci.Pipeline.JobId = env["CIRCLE_WORKFLOW_JOB_ID"]
	nci.Pipeline.JobName = env["CIRCLE_JOB"]
	nci.Pipeline.JobSlug = slug.Make(env["CIRCLE_JOB"])
	nci.Pipeline.JobStartedAt = time.Now().Format(time.RFC3339)
	nci.Pipeline.Attempt = "0"
	nci.Pipeline.Url = env["CIRCLE_BUILD_URL"]

	// repository
	projectDir, err := vcsutil.FindProjectDirectoryFromWorkDir()
	if err != nil {
		return nci, fmt.Errorf("failed to find project directory: %v", err)
	}
	vcsData, err := vcsrepository.GetVCSRepositoryInformation(projectDir)
	if err != nil {
		return nci, fmt.Errorf("failed to get repository details: %v", err)
	}
	nci.Repository = vcsData.Repository
	nci.Commit = vcsData.Commit

	// project
	projectData, err := projectdetails.GetProjectDetails(nci.Repository.Kind, nci.Repository.Remote, nci.Repository.HostType, nci.Repository.HostServer)
	if err != nil {
		return nci, fmt.Errorf("failed to get project details: %v", err)
	}
	nci.Project = projectData
	nci.Project.Dir = projectDir
	nci.Project.UID = api.GetProjectUID(nci.Repository, nci.Project)

	// flags
	nci.Flags.DeployFreeze = "false"

	return nci, nil
}
