package githubactions

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	"github.com/cidverse/normalizeci/pkg/ncispec"
	"github.com/cidverse/normalizeci/pkg/nciutil"
	"github.com/cidverse/normalizeci/pkg/normalizer/common"
	"github.com/cidverse/normalizeci/pkg/projectdetails"
	"github.com/cidverse/normalizeci/pkg/vcsrepository"
	"github.com/google/go-github/v52/github"
	"github.com/gosimple/slug"
	"github.com/rs/zerolog/log"
)

// Normalize normalizes the environment variables into the common format
func (n Normalizer) Normalize(env map[string]string) ncispec.NormalizeCISpec {
	var nci ncispec.NormalizeCISpec

	// common
	nci.Found = "true"
	nci.Version = n.version
	nci.ServiceName = n.name
	nci.ServiceSlug = n.slug

	// worker
	nci.WorkerId = env["RUNNER_TRACKING_ID"]
	nci.WorkerName = env["RUNNER_TRACKING_ID"]
	nci.WorkerType = "github_hosted_vm"
	nci.WorkerOS = env["ImageOS"] + ":" + env["ImageVersion"]
	nci.WorkerVersion = "latest"
	nci.WorkerArch = runtime.GOOS + "/" + runtime.GOARCH

	// pipeline
	nci.PipelineId = env["GITHUB_RUN_ID"]
	pipelineEvent := env["GITHUB_EVENT_NAME"]
	switch pipelineEvent {
	case "push":
		nci.PipelineTrigger = ncispec.PipelineTriggerPush
	case "pull_request":
		nci.PipelineTrigger = ncispec.PipelineTriggerMergeRequest
	default:
		nci.PipelineTrigger = ncispec.PipelineTriggerUnknown
	}

	nci.PipelineStageName = env["GITHUB_WORKFLOW"]
	nci.PipelineStageSlug = slug.Make(env["GITHUB_WORKFLOW"])
	nci.PipelineJobName = env["GITHUB_ACTION"]
	nci.PipelineJobSlug = slug.Make(env["GITHUB_ACTION"])
	nci.PipelineJobStartedAt = time.Now().UTC().Format(time.RFC3339)
	nci.PipelineAttempt = env["GITHUB_RUN_ATTEMPT"]
	nci.PipelineUrl = fmt.Sprintf("%s/%s/actions/runs/%s", env["GITHUB_SERVER_URL"], env["GITHUB_REPOSITORY"], env["GITHUB_RUN_ID"])

	// pull request (fallback in case there are issues with the event json)
	if nci.PipelineTrigger == ncispec.PipelineTriggerMergeRequest {
		splitRef := strings.Split(env["GITHUB_REF"], "/")
		nci.MergeRequestId = splitRef[2]
	}

	// repository
	projectDir := vcsrepository.FindRepositoryDirectory(common.GetWorkingDirectory())
	vcsData, addDataErr := vcsrepository.GetVCSRepositoryInformation(projectDir)
	if addDataErr != nil {
		panic(addDataErr)
	}
	nci.RepositoryKind = vcsData[ncispec.NCI_REPOSITORY_KIND]
	nci.RepositoryRemote = vcsData[ncispec.NCI_REPOSITORY_REMOTE]
	nci.RepositoryHostServer = vcsData[ncispec.NCI_REPOSITORY_HOST_SERVER]
	nci.RepositoryHostType = vcsData[ncispec.NCI_REPOSITORY_HOST_TYPE]
	nci.RepositoryStatus = vcsData[ncispec.NCI_REPOSITORY_STATUS]
	nci.CommitRefType = vcsData[ncispec.NCI_COMMIT_REF_TYPE]
	nci.CommitRefName = vcsData[ncispec.NCI_COMMIT_REF_NAME]
	nci.CommitRefPath = vcsData[ncispec.NCI_COMMIT_REF_PATH]
	nci.CommitRefSlug = vcsData[ncispec.NCI_COMMIT_REF_SLUG]
	nci.CommitRefVcs = vcsData[ncispec.NCI_COMMIT_REF_VCS]
	nci.CommitRefRelease = vcsData[ncispec.NCI_COMMIT_REF_RELEASE]
	nci.CommitSha = vcsData[ncispec.NCI_COMMIT_SHA]
	nci.CommitShaShort = vcsData[ncispec.NCI_COMMIT_SHA_SHORT]
	nci.CommitTitle = vcsData[ncispec.NCI_COMMIT_TITLE]
	nci.CommitDescription = vcsData[ncispec.NCI_COMMIT_DESCRIPTION]
	nci.CommitAuthorName = vcsData[ncispec.NCI_COMMIT_AUTHOR_NAME]
	nci.CommitAuthorEmail = vcsData[ncispec.NCI_COMMIT_AUTHOR_EMAIL]
	nci.CommitCommitterName = vcsData[ncispec.NCI_COMMIT_COMMITTER_NAME]
	nci.CommitCommitterEmail = vcsData[ncispec.NCI_COMMIT_COMMITTER_EMAIL]
	nci.CommitCount = vcsData[ncispec.NCI_COMMIT_COUNT]
	nci.LastreleaseRefName = vcsData[ncispec.NCI_LASTRELEASE_REF_NAME]
	nci.LastreleaseRefSlug = vcsData[ncispec.NCI_LASTRELEASE_REF_SLUG]
	nci.LastreleaseRefVcs = vcsData[ncispec.NCI_LASTRELEASE_REF_VCS]
	nci.LastreleaseCommitAfterCount = vcsData[ncispec.NCI_LASTRELEASE_COMMIT_AFTER_COUNT]

	// project details
	projectData := projectdetails.GetProjectDetails(nci.RepositoryKind, nci.RepositoryRemote, nci.RepositoryHostType, nci.RepositoryHostServer)
	nci.ProjectId = nciutil.GetValueFromMap(projectData, ncispec.NCI_PROJECT_ID)
	nci.ProjectName = nciutil.GetValueFromMap(projectData, ncispec.NCI_PROJECT_NAME)
	nci.ProjectPath = nciutil.GetValueFromMap(projectData, ncispec.NCI_PROJECT_PATH)
	nci.ProjectSlug = nciutil.GetValueFromMap(projectData, ncispec.NCI_PROJECT_SLUG)
	nci.ProjectDescription = nciutil.GetValueFromMap(projectData, ncispec.NCI_PROJECT_DESCRIPTION)
	nci.ProjectTopics = nciutil.GetValueFromMap(projectData, ncispec.NCI_PROJECT_TOPICS)
	nci.ProjectIssueUrl = nciutil.GetValueFromMap(projectData, ncispec.NCI_PROJECT_ISSUE_URL)
	nci.ProjectStargazers = nciutil.GetValueFromMap(projectData, ncispec.NCI_PROJECT_STARGAZERS)
	nci.ProjectForks = nciutil.GetValueFromMap(projectData, ncispec.NCI_PROJECT_FORKS)
	nci.ProjectDefaultBranch = nciutil.GetValueFromMap(projectData, ncispec.NCI_PROJECT_DEFAULT_BRANCH)
	nci.ProjectUrl = nciutil.GetValueFromMap(env, "GITHUB_SERVER_URL") + "/" + nciutil.GetValueFromMap(env, "GITHUB_REPOSITORY")
	nci.ProjectDir = projectDir

	// container registry
	nci.ContainerRegistryHost = ""
	nci.ContainerRegistryRepository = slug.Make(common.GetDirectoryNameFromPath(filepath.Join(vcsrepository.FindRepositoryDirectory(common.GetWorkingDirectory())+string(os.PathSeparator), "file")))
	nci.ContainerRegistryUsername = ""
	nci.ContainerRegistryPassword = ""
	nci.ContainerRegistryTag = nci.CommitRefRelease

	// control
	nci.DeployFreeze = "false"

	// query workflow and workflow run
	wfRun, wf, err := GetGithubWorkflowRun(env["GITHUB_REPOSITORY"], env["GITHUB_RUN_ID"])
	if err == nil {
		// pipeline
		nci.PipelineJobStartedAt = wfRun.GetRunStartedAt().UTC().Format(time.RFC3339)
		nci.PipelineConfigFile = wf.GetPath()
	}

	// parse event context
	githubEvent, err := ParseGithubEvent(os.Getenv("GITHUB_EVENT_NAME"), os.Getenv("GITHUB_EVENT_PATH"))
	if err == nil {
		variables := make(map[string]string)

		// pull request event
		if pullRequestEvent, ok := githubEvent.(*github.PullRequestEvent); ok {
			nci.MergeRequestId = fmt.Sprintf("%d", pullRequestEvent.PullRequest.GetNumber())
			nci.MergeRequestSourceBranchName = pullRequestEvent.PullRequest.Head.GetRef()
			nci.MergeRequestSourceHash = pullRequestEvent.PullRequest.Head.GetSHA()
			nci.MergeRequestTargetBranchName = pullRequestEvent.PullRequest.Base.GetRef()
			nci.MergeRequestTargetHash = pullRequestEvent.PullRequest.Base.GetSHA()
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

		nci.PipelineInput = variables
	}

	return nci
}
