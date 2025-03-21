package appveyor

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
		Id:      "0",
		Name:    "unknown",
		Type:    "appveyor_hosted_vm",
		OS:      env["APPVEYOR_BUILD_WORKER_IMAGE"],
		Version: "latest",
		Arch:    runtime.GOOS + "/" + runtime.GOARCH,
	}

	// pipeline
	nci.Pipeline.Id = env["APPVEYOR_BUILD_ID"]
	nci.Pipeline.Trigger = common.PipelineTriggerPush
	nci.Pipeline.StageId = env["APPVEYOR_BUILD_ID"]
	nci.Pipeline.StageName = "default"
	nci.Pipeline.StageSlug = slug.Make("default")
	nci.Pipeline.JobId = env["APPVEYOR_JOB_ID"]
	nci.Pipeline.JobName = env["APPVEYOR_JOB_NAME"]
	nci.Pipeline.JobSlug = slug.Make(env["APPVEYOR_JOB_NAME"])
	nci.Pipeline.JobStartedAt = time.Now().Format(time.RFC3339)
	nci.Pipeline.Attempt = env["APPVEYOR_JOB_NUMBER"]
	nci.Pipeline.Url = fmt.Sprintf("%s/project/%s/%s/builds/%s", env["APPVEYOR_URL"], env["APPVEYOR_ACCOUNT_NAME"], env["APPVEYOR_PROJECT_SLUG"], env["APPVEYOR_BUILD_ID"])

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
	nci.Commit.Hash = env["APPVEYOR_REPO_COMMIT"]
	nci.Commit.HashShort = env["APPVEYOR_REPO_COMMIT"][:7]

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
