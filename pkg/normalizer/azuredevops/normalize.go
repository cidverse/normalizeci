package azuredevops

import (
	"fmt"
	"runtime"
	"time"

	"github.com/cidverse/go-vcs/vcsutil"
	"github.com/cidverse/normalizeci/pkg/ncispec/common"
	v1 "github.com/cidverse/normalizeci/pkg/ncispec/v1"
	"github.com/cidverse/normalizeci/pkg/projectdetails"
	"github.com/cidverse/normalizeci/pkg/vcsrepository"
	"github.com/gosimple/slug"
)

// Normalize normalizes the environment variables into the common format
func (n Normalizer) Normalize(env map[string]string) v1.Spec {
	nci := v1.Create(n.name, n.slug)

	// worker
	nci.Worker = v1.Worker{
		Id:      env["AGENT_ID"],
		Name:    env["AGENT_MACHINENAME"],
		Type:    "azuredevops_hosted_vm",
		OS:      env["ImageOS"] + ":" + env["ImageVersion"],
		Version: env["AGENT_VERSION"],
		Arch:    runtime.GOOS + "/" + runtime.GOARCH,
	}

	// pipeline
	nci.Pipeline.Id = env["SYSTEM_PHASEID"]
	if env["BUILD_REASON"] == "Manual" {
		nci.Pipeline.Trigger = common.PipelineTriggerManual
	} else if env["BUILD_REASON"] == "IndividualCI" || env["BUILD_REASON"] == "BatchedCI" {
		nci.Pipeline.Trigger = common.PipelineTriggerPush
	} else if env["BUILD_REASON"] == "Schedule" {
		nci.Pipeline.Trigger = common.PipelineTriggerSchedule
	} else if env["BUILD_REASON"] == "PullRequest" {
		nci.Pipeline.Trigger = common.PipelineTriggerMergeRequest
	} else if env["BUILD_REASON"] == "BuildCompletion" {
		nci.Pipeline.Trigger = common.PipelineTriggerBuild
	} else {
		nci.Pipeline.Trigger = common.PipelineTriggerUnknown
	}
	nci.Pipeline.StageId = env["SYSTEM_STAGEID"]
	nci.Pipeline.StageName = env["SYSTEM_STAGENAME"] // SYSTEM_STAGEDISPLAYNAME
	nci.Pipeline.StageSlug = slug.Make(env["SYSTEM_STAGENAME"])
	nci.Pipeline.JobId = env["SYSTEM_JOBID"]
	nci.Pipeline.JobName = env["SYSTEM_JOBNAME"] // SYSTEM_JOBDISPLAYNAME
	nci.Pipeline.JobSlug = slug.Make(env["SYSTEM_JOBNAME"])
	nci.Pipeline.JobStartedAt = time.Now().Format(time.RFC3339)
	nci.Pipeline.Attempt = env["SYSTEM_JOBATTEMPT"]
	nci.Pipeline.Url = fmt.Sprintf("%s%s/_build/results?buildId=%s", env["SYSTEM_TEAMFOUNDATIONSERVERURI"], env["SYSTEM_TEAMPROJECT"], env["BUILD_BUILDID"])

	// repository
	projectDir, _ := vcsutil.FindProjectDirectory()
	vcsData, addDataErr := vcsrepository.GetVCSRepositoryInformation(projectDir)
	if addDataErr != nil {
		panic(addDataErr)
	}
	nci.Repository = vcsData.Repository
	nci.Commit = vcsData.Commit

	// project
	projectData, _ := projectdetails.GetProjectDetails(nci.Repository.Kind, nci.Repository.Remote, nci.Repository.HostType, nci.Repository.HostServer)
	nci.Project = projectData
	nci.Project.Url = env["BUILD_REPOSITORY_URI"]
	nci.Project.Dir = projectDir

	// flags
	nci.Flags.DeployFreeze = "false"

	return nci
}
