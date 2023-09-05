package githubactions

import (
	"encoding/json"
	"fmt"
	"os"
	"runtime"
	"strings"
	"time"

	"github.com/cidverse/go-vcs/vcsutil"
	"github.com/cidverse/normalizeci/pkg/ncispec/common"
	v1 "github.com/cidverse/normalizeci/pkg/ncispec/v1"
	"github.com/cidverse/normalizeci/pkg/nciutil"
	"github.com/cidverse/normalizeci/pkg/projectdetails"
	"github.com/cidverse/normalizeci/pkg/vcsrepository"
	"github.com/google/go-github/v55/github"
	"github.com/gosimple/slug"
	"github.com/rs/zerolog/log"
)

// Normalize normalizes the environment variables into the common format
func (n Normalizer) Normalize(env map[string]string) (v1.Spec, error) {
	nci := v1.Create(n.name, n.slug)

	// worker
	nci.Worker = v1.Worker{
		Id:      env["RUNNER_TRACKING_ID"],
		Name:    env["RUNNER_TRACKING_ID"],
		Type:    "github_hosted_vm",
		OS:      env["ImageOS"] + ":" + env["ImageVersion"],
		Version: "latest",
		Arch:    runtime.GOOS + "/" + runtime.GOARCH,
	}

	// pipeline
	nci.Pipeline.Id = env["GITHUB_RUN_ID"]
	pipelineEvent := env["GITHUB_EVENT_NAME"]
	switch pipelineEvent {
	case "push":
		nci.Pipeline.Trigger = common.PipelineTriggerPush
	case "pull_request":
		nci.Pipeline.Trigger = common.PipelineTriggerMergeRequest
	default:
		nci.Pipeline.Trigger = common.PipelineTriggerUnknown
	}

	nci.Pipeline.StageName = env["GITHUB_WORKFLOW"]
	nci.Pipeline.StageSlug = slug.Make(env["GITHUB_WORKFLOW"])
	nci.Pipeline.JobName = env["GITHUB_ACTION"]
	nci.Pipeline.JobSlug = slug.Make(env["GITHUB_ACTION"])
	nci.Pipeline.JobStartedAt = time.Now().UTC().Format(time.RFC3339)
	nci.Pipeline.Attempt = env["GITHUB_RUN_ATTEMPT"]
	nci.Pipeline.Url = fmt.Sprintf("%s/%s/actions/runs/%s", env["GITHUB_SERVER_URL"], env["GITHUB_REPOSITORY"], env["GITHUB_RUN_ID"])

	// pull request (fallback in case there are issues with the event json)
	if nci.Pipeline.Trigger == common.PipelineTriggerMergeRequest {
		splitRef := strings.Split(env["GITHUB_REF"], "/")
		nci.MergeRequest.Id = splitRef[2]
	}

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
	nci.Project.Url = nciutil.GetValueFromMap(env, "GITHUB_SERVER_URL") + "/" + nciutil.GetValueFromMap(env, "GITHUB_REPOSITORY")
	nci.Project.Dir = projectDir

	// flags
	nci.Flags.DeployFreeze = "false"

	// query workflow and workflow run
	wfRun, wf, err := GetGithubWorkflowRun(env["GITHUB_REPOSITORY"], env["GITHUB_RUN_ID"])
	if err == nil {
		// pipeline
		nci.Pipeline.JobStartedAt = wfRun.GetRunStartedAt().UTC().Format(time.RFC3339)
		nci.Pipeline.ConfigFile = wf.GetPath()
	}

	// parse event context
	githubEvent, err := ParseGithubEvent(os.Getenv("GITHUB_EVENT_NAME"), os.Getenv("GITHUB_EVENT_PATH"))
	if err == nil {
		variables := make(map[string]string)

		// pull request event
		if pullRequestEvent, ok := githubEvent.(*github.PullRequestEvent); ok {
			nci.MergeRequest = v1.MergeRequest{
				Id:               fmt.Sprintf("%d", pullRequestEvent.PullRequest.GetNumber()),
				Title:            pullRequestEvent.PullRequest.GetTitle(),
				SourceBranchName: pullRequestEvent.PullRequest.Head.GetRef(),
				SourceHash:       pullRequestEvent.PullRequest.Head.GetSHA(),
				TargetBranchName: pullRequestEvent.PullRequest.Base.GetRef(),
				TargetHash:       pullRequestEvent.PullRequest.Base.GetSHA(),
			}
		}

		// workflow dispatch event can have custom input parameters
		if dispatchEvent, ok := githubEvent.(*github.WorkflowDispatchEvent); ok {
			if dispatchEvent.Inputs != nil {
				var inputs map[string]interface{}
				err := json.Unmarshal(dispatchEvent.Inputs, &inputs)
				if err != nil {
					log.Error().Err(err).Msg("failed to parse inputs in github workflow dispatch event")
				}

				for key, value := range inputs {
					variables[key] = fmt.Sprintf("%v", value)
				}
			}
		}

		nci.Pipeline.Input = variables
	}

	return nci, nil
}
