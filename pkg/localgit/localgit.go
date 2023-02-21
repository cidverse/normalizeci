package localgit

import (
	"os"
	"path/filepath"
	"runtime"

	"github.com/cidverse/normalizeci/pkg/common"
	"github.com/cidverse/normalizeci/pkg/ncispec"
	"github.com/cidverse/normalizeci/pkg/nciutil"
	"github.com/cidverse/normalizeci/pkg/projectdetails"
	"github.com/cidverse/normalizeci/pkg/vcsrepository"
	"github.com/gosimple/slug"
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
func (n Normalizer) Normalize(env map[string]string) map[string]string {
	var nci ncispec.NormalizeCISpec

	// common
	nci.Found = "true"
	nci.Version = n.version
	nci.ServiceName = n.name
	nci.ServiceSlug = n.slug

	// worker
	nci.WorkerId = "local"
	nci.WorkerName = "localhost"
	nci.WorkerVersion = "1.0.0"
	nci.WorkerArch = runtime.GOOS + "/" + runtime.GOARCH

	// pipeline
	nci.PipelineId = nciutil.GenerateSnowflakeId()
	nci.PipelineTrigger = ncispec.PipelineTriggerCLI
	nci.PipelineStageId = nciutil.GenerateSnowflakeId()
	nci.PipelineStageName = ncispec.PipelineStageDefault
	nci.PipelineStageSlug = ncispec.PipelineStageDefault
	nci.PipelineJobId = nciutil.GenerateSnowflakeId()
	nci.PipelineJobName = ncispec.PipelineJobDefault
	nci.PipelineJobSlug = ncispec.PipelineJobDefault

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
	nci.ProjectUrl = nciutil.GetValueFromMap(projectData, ncispec.NCI_PROJECT_URL)
	nci.ProjectDir = projectDir

	// container registry
	nci.ContainerregistryHost = ""
	nci.ContainerregistryUsername = ""
	nci.ContainerregistryPassword = ""
	if len(nci.ProjectPath) > 0 {
		nci.ContainerregistryRepository = nci.ProjectPath
	} else {
		nci.ContainerregistryRepository = slug.Make(common.GetDirectoryNameFromPath(filepath.Join(vcsrepository.FindRepositoryDirectory(common.GetWorkingDirectory())+string(os.PathSeparator), "file")))
	}
	nci.ContainerregistryTag = nci.CommitRefRelease

	nci.DeployFreeze = "false"

	return ncispec.ToMap(nci)
}

func (n Normalizer) Denormalize(env map[string]string) map[string]string {
	return env
}

// NewNormalizer gets a instance of the normalizer
func NewNormalizer() Normalizer {
	entity := Normalizer{
		version: "0.3.0",
		name:    "Local Git Repository",
		slug:    "local-git",
	}

	return entity
}
