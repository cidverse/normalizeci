package azuredevops

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"time"

	"github.com/cidverse/normalizeci/pkg/common"
	"github.com/cidverse/normalizeci/pkg/ncispec"
	"github.com/cidverse/normalizeci/pkg/nciutil"
	"github.com/cidverse/normalizeci/pkg/projectdetails"
	"github.com/cidverse/normalizeci/pkg/vcsrepository"
	"github.com/gosimple/slug"
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
	nci.WorkerId = env["AGENT_ID"]
	nci.WorkerName = env["AGENT_MACHINENAME"]
	nci.WorkerType = "azuredevops_hosted_vm"
	nci.WorkerOS = env["ImageOS"] + ":" + env["ImageVersion"]
	nci.WorkerVersion = env["AGENT_VERSION"]
	nci.WorkerArch = runtime.GOOS + "/" + runtime.GOARCH

	// pipeline
	nci.PipelineId = env["SYSTEM_PHASEID"]
	if env["BUILD_REASON"] == "Manual" {
		nci.PipelineTrigger = ncispec.PipelineTriggerManual
	} else if env["BUILD_REASON"] == "IndividualCI" || env["BUILD_REASON"] == "BatchedCI" {
		nci.PipelineTrigger = ncispec.PipelineTriggerPush
	} else if env["BUILD_REASON"] == "Schedule" {
		nci.PipelineTrigger = ncispec.PipelineTriggerSchedule
	} else if env["BUILD_REASON"] == "PullRequest" {
		nci.PipelineTrigger = ncispec.PipelineTriggerMergeRequest
	} else if env["BUILD_REASON"] == "BuildCompletion" {
		nci.PipelineTrigger = ncispec.PipelineTriggerBuild
	} else {
		nci.PipelineTrigger = ncispec.PipelineTriggerUnknown
	}
	nci.PipelineStageId = env["SYSTEM_STAGEID"]
	nci.PipelineStageName = env["SYSTEM_STAGENAME"] // SYSTEM_STAGEDISPLAYNAME
	nci.PipelineStageSlug = slug.Make(env["SYSTEM_STAGENAME"])
	nci.PipelineJobId = env["SYSTEM_JOBID"]
	nci.PipelineJobName = env["SYSTEM_JOBNAME"] // SYSTEM_JOBDISPLAYNAME
	nci.PipelineJobSlug = slug.Make(env["SYSTEM_JOBNAME"])
	nci.PipelineJobStartedAt = time.Now().Format(time.RFC3339)
	nci.PipelineAttempt = env["SYSTEM_JOBATTEMPT"]
	nci.PipelineUrl = fmt.Sprintf("%s%s/_build/results?buildId=%s", env["SYSTEM_TEAMFOUNDATIONSERVERURI"], env["SYSTEM_TEAMPROJECT"], env["BUILD_BUILDID"])

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
	nci.ProjectUrl = env["BUILD_REPOSITORY_URI"]
	nci.ProjectDir = projectDir

	// container registry
	nci.ContainerRegistryHost = ""
	nci.ContainerRegistryRepository = slug.Make(common.GetDirectoryNameFromPath(filepath.Join(vcsrepository.FindRepositoryDirectory(common.GetWorkingDirectory())+string(os.PathSeparator), "file")))
	nci.ContainerRegistryUsername = ""
	nci.ContainerRegistryPassword = ""
	nci.ContainerRegistryTag = nci.CommitRefRelease

	// control
	nci.DeployFreeze = "false"

	return nci
}
