package gitlabci

import (
	"runtime"

	"github.com/cidverse/normalizeci/pkg/common"
	"github.com/cidverse/normalizeci/pkg/ncispec"
	"github.com/cidverse/normalizeci/pkg/nciutil"
	"github.com/cidverse/normalizeci/pkg/projectdetails"
	"github.com/cidverse/normalizeci/pkg/vcsrepository"
	"github.com/gosimple/slug"
)

// Normalize normalizes the environment variables into the common format
func (n Normalizer) Normalize(env map[string]string) map[string]string {
	var nci ncispec.NormalizeCISpec

	// common
	nci.Found = "true"
	nci.Version = n.version
	nci.ServiceName = n.name
	nci.ServiceSlug = n.slug

	// worker
	nci.WorkerId = env["CI_RUNNER_ID"]
	nci.WorkerName = env["CI_RUNNER_DESCRIPTION"]
	nci.WorkerType = "gitlab_hosted_vm"
	nci.WorkerOS = ""
	nci.WorkerVersion = env["CI_RUNNER_VERSION"]
	nci.WorkerArch = runtime.GOOS + "/" + runtime.GOARCH

	// pipeline
	nci.PipelineId = env["CI_PIPELINE_ID"]
	nci.PipelineTrigger = gitlabTriggerNormalize(env["CI_PIPELINE_SOURCE"])
	nci.PipelineStageName = env["CI_JOB_STAGE"]
	nci.PipelineStageSlug = slug.Make(env["CI_JOB_STAGE"])
	nci.PipelineJobId = env["CI_JOB_ID"]
	nci.PipelineJobName = env["CI_JOB_NAME"]
	nci.PipelineJobSlug = slug.Make(env["CI_JOB_NAME"])
	nci.PipelineJobStartedAt = env["CI_JOB_STARTED_AT"]
	nci.PipelineAttempt = "1"
	nci.PipelineConfigFile = "gitlab-ci.yml"
	nci.PipelineUrl = env["CI_JOB_URL"]

	// merge request
	if mergeRequestId, isMergeRequest := env["CI_MERGE_REQUEST_IID"]; isMergeRequest {
		nci.MergeRequestId = mergeRequestId
		nci.MergeRequestSourceBranchName = env["CI_MERGE_REQUEST_SOURCE_BRANCH_NAME"]
		nci.MergeRequestTargetBranchName = env["CI_MERGE_REQUEST_TARGET_BRANCH_NAME"]
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
	if len(env["CI_COMMIT_TAG"]) > 0 {
		nci.CommitRefType = "tag"
		nci.CommitRefName = env["CI_COMMIT_TAG"]
		nci.CommitRefPath = nci.CommitRefType + "/" + env["CI_COMMIT_TAG"]
		nci.CommitRefSlug = slug.Make(env["CI_COMMIT_TAG"])
		nci.CommitRefVcs = "refs/tags/" + env["CI_COMMIT_TAG"]
	} else {
		nci.CommitRefType = "branch"
		nci.CommitRefName = env["CI_COMMIT_REF_NAME"]
		nci.CommitRefPath = nci.CommitRefType + "/" + env["CI_COMMIT_REF_NAME"]
		nci.CommitRefSlug = slug.Make(env["CI_COMMIT_REF_NAME"])
		nci.CommitRefVcs = "refs/heads/" + env["CI_COMMIT_REF_NAME"]
	}
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
	nci.ProjectId = nciutil.FirstNonEmpty([]string{nciutil.GetValueFromMap(env, "CI_PROJECT_ID"), nciutil.GetValueFromMap(projectData, ncispec.NCI_PROJECT_ID)})
	nci.ProjectName = nciutil.FirstNonEmpty([]string{nciutil.GetValueFromMap(env, "CI_PROJECT_TITLE"), nciutil.GetValueFromMap(projectData, ncispec.NCI_PROJECT_NAME)})
	nci.ProjectPath = nciutil.FirstNonEmpty([]string{nciutil.GetValueFromMap(env, "CI_PROJECT_NAME"), nciutil.GetValueFromMap(projectData, ncispec.NCI_PROJECT_PATH)})
	nci.ProjectSlug = nciutil.FirstNonEmpty([]string{nciutil.GetValueFromMap(env, "CI_PROJECT_PATH_SLUG"), nciutil.GetValueFromMap(projectData, ncispec.NCI_PROJECT_SLUG)})
	nci.ProjectDescription = nciutil.FirstNonEmpty([]string{nciutil.GetValueFromMap(env, "CI_PROJECT_DESCRIPTION"), nciutil.GetValueFromMap(projectData, ncispec.NCI_PROJECT_DESCRIPTION)})
	nci.ProjectTopics = nciutil.FirstNonEmpty([]string{nciutil.GetValueFromMap(projectData, ncispec.NCI_PROJECT_TOPICS)})
	nci.ProjectIssueUrl = nciutil.FirstNonEmpty([]string{nciutil.GetValueFromMap(projectData, ncispec.NCI_PROJECT_ISSUE_URL)})
	nci.ProjectStargazers = nciutil.FirstNonEmpty([]string{nciutil.GetValueFromMap(projectData, ncispec.NCI_PROJECT_STARGAZERS)})
	nci.ProjectForks = nciutil.FirstNonEmpty([]string{nciutil.GetValueFromMap(projectData, ncispec.NCI_PROJECT_FORKS)})
	nci.ProjectDefaultBranch = nciutil.FirstNonEmpty([]string{nciutil.GetValueFromMap(env, "CI_DEFAULT_BRANCH"), nciutil.GetValueFromMap(projectData, ncispec.NCI_PROJECT_DEFAULT_BRANCH)})
	nci.ProjectUrl = nciutil.FirstNonEmpty([]string{nciutil.GetValueFromMap(env, "CI_PROJECT_URL"), nciutil.GetValueFromMap(projectData, ncispec.NCI_PROJECT_URL)})
	nci.ProjectDir = projectDir

	// container registry
	nci.ContainerRegistryHost = env["CI_REGISTRY"]
	nci.ContainerRegistryRepository = env["CI_REGISTRY_IMAGE"]
	if len(env["CI_DEPLOY_USER"]) > 0 {
		nci.ContainerRegistryUsername = env["CI_DEPLOY_USER"]
		nci.ContainerRegistryPassword = env["CI_DEPLOY_PASSWORD"]
	} else {
		nci.ContainerRegistryUsername = env["CI_REGISTRY_USER"]
		nci.ContainerRegistryPassword = env["CI_REGISTRY_PASSWORD"]
	}
	nci.ContainerRegistryTag = nci.CommitRefRelease

	// control
	if _, ok := env["CI_DEPLOY_FREEZE"]; ok {
		nci.DeployFreeze = env["CI_DEPLOY_FREEZE"]
	} else {
		nci.DeployFreeze = "false"
	}

	// custom input parameters
	variables, err := GetGitlabPipelineRun(env["CI_SERVER_URL"], env["CI_PROJECT_ID"], env["CI_PIPELINE_ID"], env["CI_JOB_TOKEN"])
	if err == nil {
		v := make(map[string]string)

		for _, variable := range variables {
			v[variable.Key] = variable.Value
		}

		nci.PipelineInput = v
	}

	return ncispec.ToMap(nci)
}

func gitlabTriggerNormalize(input string) string {
	if input == "merge_request_event" || input == "external_pull_request_event" {
		return ncispec.PipelineTriggerMergeRequest
	}
	if input == "schedule" {
		return ncispec.PipelineTriggerSchedule
	}

	return input
}
