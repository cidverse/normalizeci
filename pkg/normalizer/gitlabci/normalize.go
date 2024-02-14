package gitlabci

import (
	"fmt"
	"runtime"

	"github.com/cidverse/go-vcs/vcsutil"
	"github.com/cidverse/normalizeci/pkg/ncispec/common"
	v1 "github.com/cidverse/normalizeci/pkg/ncispec/v1"
	"github.com/cidverse/normalizeci/pkg/nciutil"
	"github.com/cidverse/normalizeci/pkg/projectdetails"
	"github.com/cidverse/normalizeci/pkg/vcsrepository"
	"github.com/gosimple/slug"
	"github.com/rs/zerolog/log"
)

// Normalize normalizes the environment variables into the common format
func (n Normalizer) Normalize(env map[string]string) (v1.Spec, error) {
	nci := v1.Create(n.name, n.slug)

	// worker
	nci.Worker = v1.Worker{
		Id:      env["CI_RUNNER_ID"],
		Name:    env["CI_RUNNER_DESCRIPTION"],
		Type:    "gitlab_hosted_vm",
		OS:      "",
		Version: env["CI_RUNNER_VERSION"],
		Arch:    runtime.GOOS + "/" + runtime.GOARCH,
	}

	// pipeline
	nci.Pipeline.Id = env["CI_PIPELINE_ID"]
	nci.Pipeline.Trigger = gitlabTriggerNormalize(env["CI_PIPELINE_SOURCE"])
	nci.Pipeline.StageName = env["CI_JOB_STAGE"]
	nci.Pipeline.StageSlug = slug.Make(env["CI_JOB_STAGE"])
	nci.Pipeline.JobId = env["CI_JOB_ID"]
	nci.Pipeline.JobName = env["CI_JOB_NAME"]
	nci.Pipeline.JobSlug = slug.Make(env["CI_JOB_NAME"])
	nci.Pipeline.JobStartedAt = env["CI_JOB_STARTED_AT"]
	nci.Pipeline.Attempt = "1"
	nci.Pipeline.ConfigFile = "gitlab-ci.yml"
	nci.Pipeline.Url = env["CI_JOB_URL"]

	// merge request
	if mergeRequestId, isMergeRequest := env["CI_MERGE_REQUEST_IID"]; isMergeRequest {
		nci.MergeRequest.Id = mergeRequestId
		nci.MergeRequest.SourceBranchName = env["CI_MERGE_REQUEST_SOURCE_BRANCH_NAME"]
		nci.MergeRequest.SourceHash = env["CI_MERGE_REQUEST_SOURCE_BRANCH_SHA"]
		nci.MergeRequest.TargetBranchName = env["CI_MERGE_REQUEST_TARGET_BRANCH_NAME"]
		nci.MergeRequest.TargetHash = env["CI_MERGE_REQUEST_SOURCE_BRANCH_SHA"]
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
	if len(env["CI_COMMIT_TAG"]) > 0 {
		nci.Commit.RefType = "tag"
		nci.Commit.RefName = env["CI_COMMIT_TAG"]
		nci.Commit.RefPath = nci.Commit.RefType + "/" + env["CI_COMMIT_TAG"]
		nci.Commit.RefSlug = slug.Make(env["CI_COMMIT_TAG"])
		nci.Commit.RefVCS = "refs/tags/" + env["CI_COMMIT_TAG"]
	} else {
		nci.Commit.RefType = "branch"
		nci.Commit.RefName = env["CI_COMMIT_REF_NAME"]
		nci.Commit.RefPath = nci.Commit.RefType + "/" + env["CI_COMMIT_REF_NAME"]
		nci.Commit.RefSlug = slug.Make(env["CI_COMMIT_REF_NAME"])
		nci.Commit.RefVCS = "refs/heads/" + env["CI_COMMIT_REF_NAME"]
	}

	// project details
	projectData, err := projectdetails.GetProjectDetails(nci.Repository.Kind, nci.Repository.Remote, nci.Repository.HostType, nci.Repository.HostServer)
	if err != nil {
		// CI_JOB_TOKEN read_project access is pending for 6 years (https://gitlab.com/gitlab-org/gitlab/-/issues/17511)
		log.Debug().Err(err).Msg("failed to get project details")
	}
	nci.Project.Id = nciutil.FirstNonEmpty([]string{nciutil.GetValueFromMap(env, "CI_PROJECT_ID"), projectData.Id})
	nci.Project.Name = nciutil.FirstNonEmpty([]string{nciutil.GetValueFromMap(env, "CI_PROJECT_TITLE"), projectData.Name})
	nci.Project.Path = nciutil.FirstNonEmpty([]string{nciutil.GetValueFromMap(env, "CI_PROJECT_NAME"), projectData.Path})
	nci.Project.Slug = nciutil.FirstNonEmpty([]string{nciutil.GetValueFromMap(env, "CI_PROJECT_PATH_SLUG"), projectData.Slug})
	nci.Project.Description = nciutil.FirstNonEmpty([]string{nciutil.GetValueFromMap(env, "CI_PROJECT_DESCRIPTION"), projectData.Description})
	nci.Project.Topics = nciutil.FirstNonEmpty([]string{projectData.Topics})
	nci.Project.IssueUrl = nciutil.FirstNonEmpty([]string{projectData.IssueUrl})
	nci.Project.Stargazers = nciutil.FirstNonEmpty([]string{projectData.Stargazers})
	nci.Project.Forks = nciutil.FirstNonEmpty([]string{projectData.Forks})
	nci.Project.DefaultBranch = nciutil.FirstNonEmpty([]string{nciutil.GetValueFromMap(env, "CI_DEFAULT_BRANCH"), projectData.DefaultBranch})
	nci.Project.Url = nciutil.FirstNonEmpty([]string{nciutil.GetValueFromMap(env, "CI_PROJECT_URL"), projectData.Url})
	nci.Project.Dir = projectDir

	// flags
	if _, ok := env["CI_DEPLOY_FREEZE"]; ok {
		nci.Flags.DeployFreeze = env["CI_DEPLOY_FREEZE"]
	} else {
		nci.Flags.DeployFreeze = "false"
	}

	// custom input parameters
	variables, err := GetGitlabPipelineRun(env["CI_SERVER_URL"], env["CI_PROJECT_ID"], env["CI_PIPELINE_ID"], env["CI_JOB_TOKEN"])
	if err == nil {
		v := make(map[string]string)

		for _, variable := range variables {
			v[variable.Key] = variable.Value
		}

		nci.Pipeline.Input = v
	}

	return nci, nil
}

func gitlabTriggerNormalize(input string) string {
	if input == "merge_request_event" || input == "external_pull_request_event" {
		return common.PipelineTriggerMergeRequest
	}
	if input == "schedule" {
		return common.PipelineTriggerSchedule
	}

	return input
}
