package ncispec

import (
	"strings"
)

const (
	NCI                                  = "NCI"
	NCI_VERSION                          = "NCI_VERSION"
	NCI_SERVICE_NAME                     = "NCI_SERVICE_NAME"
	NCI_SERVICE_SLUG                     = "NCI_SERVICE_SLUG"
	NCI_WORKER_ID                        = "NCI_WORKER_ID"
	NCI_WORKER_NAME                      = "NCI_WORKER_NAME"
	NCI_WORKER_TYPE                      = "NCI_WORKER_TYPE"
	NCI_WORKER_OS                        = "NCI_WORKER_OS"
	NCI_WORKER_VERSION                   = "NCI_WORKER_VERSION"
	NCI_WORKER_ARCH                      = "NCI_WORKER_ARCH"
	NCI_PIPELINE_ID                      = "NCI_PIPELINE_ID"
	NCI_PIPELINE_TRIGGER                 = "NCI_PIPELINE_TRIGGER"
	NCI_PIPELINE_STAGE_ID                = "NCI_PIPELINE_STAGE_ID"
	NCI_PIPELINE_STAGE_NAME              = "NCI_PIPELINE_STAGE_NAME"
	NCI_PIPELINE_STAGE_SLUG              = "NCI_PIPELINE_STAGE_SLUG"
	NCI_PIPELINE_JOB_ID                  = "NCI_PIPELINE_JOB_ID"
	NCI_PIPELINE_JOB_NAME                = "NCI_PIPELINE_JOB_NAME"
	NCI_PIPELINE_JOB_SLUG                = "NCI_PIPELINE_JOB_SLUG"
	NCI_PIPELINE_JOB_STARTED_AT          = "NCI_PIPELINE_JOB_STARTED_AT"
	NCI_PIPELINE_ATTEMPT                 = "NCI_PIPELINE_ATTEMPT"
	NCI_PIPELINE_CONFIG_FILE             = "NCI_PIPELINE_CONFIG_FILE"
	NCI_PIPELINE_URL                     = "NCI_PIPELINE_URL"
	NCI_PIPELINE_INPUT                   = "NCI_PIPELINE_INPUT"
	NCI_MERGE_REQUEST_ID                 = "NCI_MERGE_REQUEST_ID"
	NCI_MERGE_REQUEST_SOURCE_BRANCH_NAME = "NCI_MERGE_REQUEST_SOURCE_BRANCH_NAME"
	NCI_MERGE_REQUEST_TARGET_BRANCH_NAME = "NCI_MERGE_REQUEST_TARGET_BRANCH_NAME"
	NCI_PROJECT_ID                       = "NCI_PROJECT_ID"
	NCI_PROJECT_NAME                     = "NCI_PROJECT_NAME"
	NCI_PROJECT_PATH                     = "NCI_PROJECT_PATH"
	NCI_PROJECT_SLUG                     = "NCI_PROJECT_SLUG"
	NCI_PROJECT_DESCRIPTION              = "NCI_PROJECT_DESCRIPTION"
	NCI_PROJECT_TOPICS                   = "NCI_PROJECT_TOPICS"
	NCI_PROJECT_ISSUE_URL                = "NCI_PROJECT_ISSUE_URL"
	NCI_PROJECT_STARGAZERS               = "NCI_PROJECT_STARGAZERS"
	NCI_PROJECT_FORKS                    = "NCI_PROJECT_FORKS"
	NCI_PROJECT_DIR                      = "NCI_PROJECT_DIR"
	NCI_PROJECT_DEFAULT_BRANCH           = "NCI_PROJECT_DEFAULT_BRANCH"
	NCI_PROJECT_URL                      = "NCI_PROJECT_URL"
	NCI_REPOSITORY_KIND                  = "NCI_REPOSITORY_KIND"
	NCI_REPOSITORY_REMOTE                = "NCI_REPOSITORY_REMOTE"
	NCI_REPOSITORY_HOST_SERVER           = "NCI_REPOSITORY_HOST_SERVER"
	NCI_REPOSITORY_HOST_TYPE             = "NCI_REPOSITORY_HOST_TYPE"
	NCI_REPOSITORY_STATUS                = "NCI_REPOSITORY_STATUS"
	NCI_COMMIT_REF_TYPE                  = "NCI_COMMIT_REF_TYPE"
	NCI_COMMIT_REF_NAME                  = "NCI_COMMIT_REF_NAME"
	NCI_COMMIT_REF_PATH                  = "NCI_COMMIT_REF_PATH"
	NCI_COMMIT_REF_SLUG                  = "NCI_COMMIT_REF_SLUG"
	NCI_COMMIT_REF_VCS                   = "NCI_COMMIT_REF_VCS"
	NCI_COMMIT_REF_RELEASE               = "NCI_COMMIT_REF_RELEASE"
	NCI_COMMIT_SHA                       = "NCI_COMMIT_SHA"
	NCI_COMMIT_SHA_SHORT                 = "NCI_COMMIT_SHA_SHORT"
	NCI_COMMIT_AUTHOR_NAME               = "NCI_COMMIT_AUTHOR_NAME"
	NCI_COMMIT_AUTHOR_EMAIL              = "NCI_COMMIT_AUTHOR_EMAIL"
	NCI_COMMIT_COMMITTER_NAME            = "NCI_COMMIT_COMMITTER_NAME"
	NCI_COMMIT_COMMITTER_EMAIL           = "NCI_COMMIT_COMMITTER_EMAIL"
	NCI_COMMIT_TITLE                     = "NCI_COMMIT_TITLE"
	NCI_COMMIT_DESCRIPTION               = "NCI_COMMIT_DESCRIPTION"
	NCI_COMMIT_COUNT                     = "NCI_COMMIT_COUNT"
	NCI_CONTAINERREGISTRY_HOST           = "NCI_CONTAINERREGISTRY_HOST"
	NCI_CONTAINERREGISTRY_USERNAME       = "NCI_CONTAINERREGISTRY_USERNAME"
	NCI_CONTAINERREGISTRY_PASSWORD       = "NCI_CONTAINERREGISTRY_PASSWORD"
	NCI_CONTAINERREGISTRY_REPOSITORY     = "NCI_CONTAINERREGISTRY_REPOSITORY"
	NCI_CONTAINERREGISTRY_TAG            = "NCI_CONTAINERREGISTRY_TAG"
	NCI_LASTRELEASE_REF_NAME             = "NCI_LASTRELEASE_REF_NAME"
	NCI_LASTRELEASE_REF_SLUG             = "NCI_LASTRELEASE_REF_SLUG"
	NCI_LASTRELEASE_REF_VCS              = "NCI_LASTRELEASE_REF_VCS"
	NCI_LASTRELEASE_COMMIT_AFTER_COUNT   = "NCI_LASTRELEASE_COMMIT_AFTER_COUNT"
	NCI_DEPLOY_FREEZE                    = "NCI_DEPLOY_FREEZE"
)

type NormalizeCISpec struct {
	Found       string `validate:"required"`         // Will be set the true, if the variables have been normalized. (this script)
	Version     string `validate:"required"`         // The revision of nci that was used to generate the normalized variables.
	ServiceName string `validate:"required"`         // The commercial name of the used ci service. (e.g. GitLab CI, Travis CI, CircleCI, Jenkins)
	ServiceSlug string `validate:"required,is-slug"` // The commercial name normalized as slug for use in scripts, will not be changed.

	WorkerId      string `validate:"required"` // A unique id of the ci worker.
	WorkerName    string `validate:"required"` // The human readable name of the ci worker.
	WorkerType    string `validate:"required"`
	WorkerOS      string // Worker OS or OS Image
	WorkerVersion string `validate:"required"`         // The version of the ci worker.
	WorkerArch    string `validate:"required,is-arch"` // The arch of the ci worker. (ie. linux/amd64)

	PipelineId           string `validate:"required"`
	PipelineTrigger      string `validate:"required,oneof=cli manual push trigger api schedule pull_request build"` // What triggered the pipeline. (ie. manual/push/trigger/api/schedule/pull_request/build)
	PipelineStageId      string
	PipelineStageName    string `validate:"required"`         // Human readable name of the current stage.
	PipelineStageSlug    string `validate:"required,is-slug"` // Slug of the current stage.
	PipelineJobId        string
	PipelineJobName      string `validate:"required"`         // Human readable name of the current job.
	PipelineJobSlug      string `validate:"required,is-slug"` // Slug of the current job.
	PipelineJobStartedAt string `validate:"required"`
	PipelineAttempt      string `validate:"number"`
	PipelineConfigFile   string // Pipeline Config File
	PipelineUrl          string // Pipeline URL
	PipelineInput        map[string]string

	MergeRequestId               string `validate:"required_if=PipelineTrigger pull_request"` // The number of the pull request, is only present if `PipelineTrigger` = pull_request.
	MergeRequestSourceBranchName string
	MergeRequestTargetBranchName string

	ProjectId            string // Unique project id, can be used in deployments.
	ProjectName          string // Unique project id, can be used in deployments.
	ProjectPath          string // Path of the Namespace and the project
	ProjectSlug          string `validate:"required,is-slug"` // Project slug, that can be used in deployments.
	ProjectDescription   string // The project description.
	ProjectTopics        string // The topics / tags of the project.
	ProjectIssueUrl      string // A template for links to issues, contains a `{ID}` placeholder.
	ProjectStargazers    string `validate:"number"` // The number of people who `follow` / `bookmarked` the project.
	ProjectForks         string `validate:"number"` // The number of forks of the project.
	ProjectDir           string // Project directory on the local filesystem.
	ProjectUrl           string
	ProjectDefaultBranch string `` // The default branch

	ContainerRegistryHost       string // The hostname of the container registry.
	ContainerRegistryUsername   string // The username used for container registry authentication.
	ContainerRegistryPassword   string // The password used for container registry authentication.
	ContainerRegistryRepository string `validate:"required"` // The repository, that should be used for the current project.
	ContainerRegistryTag        string `validate:"required"` // The tag that should be build.

	RepositoryKind       string `validate:"required"` //  The used version control system. (git)
	RepositoryRemote     string `validate:"required"` // The remote url pointing at the repository. (git remote url or `local` if no remote was found)
	RepositoryHostServer string `validate:"required"` // Host of the git repository server, for example github.com
	RepositoryHostType   string `validate:"required"` // Type of the git repository server (github, gitlab, ...)
	RepositoryStatus     string `validate:"required"` // The repository status (dirty, clean)
	CommitRefType        string `validate:"required"` // The reference type. (branch / tag)
	CommitRefName        string `validate:"required"` // Human-readable name of the current repository reference.
	CommitRefPath        string `validate:"required"` // Combination of the ref type and ref name. (tag/v1.0.0 or branch/main)
	CommitRefSlug        string `validate:"required"` // Slug of the current repository reference.
	CommitRefVcs         string `validate:"required"` // Holds the vcs specific absolute reference name. (ex: `refs/heads/main`)// Release version of the artifact, without leading `v` or `/` - should be in format `x.y.z` or `feature-abc`.
	CommitRefRelease     string `validate:"required"` // Release version of the artifact, without leading `v` or `/` - should be in format `x.y.z` or `feature-abc`.
	CommitSha            string `validate:"required"` // A unique hash, that each commit gets.
	CommitShaShort       string `validate:"required"` // A short form of the unique commit hash. (8 chars)
	CommitAuthorName     string `validate:"required"` // author name
	CommitAuthorEmail    string `validate:"required"` // author email
	CommitCommitterName  string `validate:"required"` // committer name
	CommitCommitterEmail string `validate:"required"` // committer email
	CommitTitle          string `validate:"required"` // The title of the latest commit on the current reference.
	CommitDescription    string // The description of the latest commit on the current reference.
	CommitCount          string `validate:"required"` // The total amount of commits inside of the current reference, can be used as build number.

	LastreleaseRefName          string `validate:"required"` // Human readable name of the last stable release.
	LastreleaseRefSlug          string `validate:"required"` // Slug of the last stable release.
	LastreleaseRefVcs           string `validate:"required"` // Holds the vcs specific absolute reference name of the last stable release. (ex: `refs/heads/main`)
	LastreleaseCommitAfterCount string `validate:"required"` // Holds the count of commits since the last stable release.

	DeployFreeze string `validate:"required,boolean"` // Currently in a deploy freeze window? (`true`, `false`)
}

func getInputFromEnv(data map[string]string) map[string]string {
	var result map[string]string

	for k, v := range data {
		if strings.HasPrefix(k, NCI_PIPELINE_INPUT) {
			result[strings.TrimPrefix(k, NCI_PIPELINE_INPUT+"_")] = v
		}
	}

	return result
}

func OfMap(data map[string]string) NormalizeCISpec {
	return NormalizeCISpec{
		Found:       data[NCI],
		Version:     data[NCI_VERSION],
		ServiceName: data[NCI_SERVICE_NAME],
		ServiceSlug: data[NCI_SERVICE_SLUG],

		WorkerId:      data[NCI_WORKER_ID],
		WorkerName:    data[NCI_WORKER_NAME],
		WorkerType:    data[NCI_WORKER_TYPE],
		WorkerOS:      data[NCI_WORKER_OS],
		WorkerVersion: data[NCI_WORKER_VERSION],
		WorkerArch:    data[NCI_WORKER_ARCH],

		PipelineId:           data[NCI_PIPELINE_ID],
		PipelineTrigger:      data[NCI_PIPELINE_TRIGGER],
		PipelineStageId:      data[NCI_PIPELINE_STAGE_ID],
		PipelineStageName:    data[NCI_PIPELINE_STAGE_NAME],
		PipelineStageSlug:    data[NCI_PIPELINE_STAGE_SLUG],
		PipelineJobId:        data[NCI_PIPELINE_JOB_ID],
		PipelineJobName:      data[NCI_PIPELINE_JOB_NAME],
		PipelineJobSlug:      data[NCI_PIPELINE_JOB_SLUG],
		PipelineJobStartedAt: data[NCI_PIPELINE_JOB_STARTED_AT],
		PipelineAttempt:      data[NCI_PIPELINE_ATTEMPT],
		PipelineConfigFile:   data[NCI_PIPELINE_CONFIG_FILE],
		PipelineUrl:          data[NCI_PIPELINE_URL],
		PipelineInput:        getInputFromEnv(data),

		MergeRequestId:               data[NCI_MERGE_REQUEST_ID],
		MergeRequestSourceBranchName: data[NCI_MERGE_REQUEST_SOURCE_BRANCH_NAME],
		MergeRequestTargetBranchName: data[NCI_MERGE_REQUEST_TARGET_BRANCH_NAME],

		ProjectId:            data[NCI_PROJECT_ID],
		ProjectName:          data[NCI_PROJECT_NAME],
		ProjectPath:          data[NCI_PROJECT_PATH],
		ProjectSlug:          data[NCI_PROJECT_SLUG],
		ProjectDescription:   data[NCI_PROJECT_DESCRIPTION],
		ProjectTopics:        data[NCI_PROJECT_TOPICS],
		ProjectIssueUrl:      data[NCI_PROJECT_ISSUE_URL],
		ProjectStargazers:    data[NCI_PROJECT_STARGAZERS],
		ProjectForks:         data[NCI_PROJECT_FORKS],
		ProjectDefaultBranch: data[NCI_PROJECT_DEFAULT_BRANCH],
		ProjectUrl:           data[NCI_PROJECT_URL],
		ProjectDir:           data[NCI_PROJECT_DIR],

		ContainerRegistryHost:       data[NCI_CONTAINERREGISTRY_HOST],
		ContainerRegistryUsername:   data[NCI_CONTAINERREGISTRY_USERNAME],
		ContainerRegistryPassword:   data[NCI_CONTAINERREGISTRY_PASSWORD],
		ContainerRegistryRepository: data[NCI_CONTAINERREGISTRY_REPOSITORY],
		ContainerRegistryTag:        data[NCI_CONTAINERREGISTRY_TAG],
		RepositoryKind:              data[NCI_REPOSITORY_KIND],
		RepositoryRemote:            data[NCI_REPOSITORY_REMOTE],
		RepositoryHostServer:        data[NCI_REPOSITORY_HOST_SERVER],
		RepositoryHostType:          data[NCI_REPOSITORY_HOST_TYPE],
		RepositoryStatus:            data[NCI_REPOSITORY_STATUS],

		CommitRefType:        data[NCI_COMMIT_REF_TYPE],
		CommitRefName:        data[NCI_COMMIT_REF_NAME],
		CommitRefPath:        data[NCI_COMMIT_REF_PATH],
		CommitRefSlug:        data[NCI_COMMIT_REF_SLUG],
		CommitRefVcs:         data[NCI_COMMIT_REF_VCS],
		CommitRefRelease:     data[NCI_COMMIT_REF_RELEASE],
		CommitSha:            data[NCI_COMMIT_SHA],
		CommitShaShort:       data[NCI_COMMIT_SHA_SHORT],
		CommitAuthorName:     data[NCI_COMMIT_AUTHOR_NAME],
		CommitAuthorEmail:    data[NCI_COMMIT_AUTHOR_EMAIL],
		CommitCommitterName:  data[NCI_COMMIT_COMMITTER_NAME],
		CommitCommitterEmail: data[NCI_COMMIT_COMMITTER_EMAIL],
		CommitTitle:          data[NCI_COMMIT_TITLE],
		CommitDescription:    data[NCI_COMMIT_DESCRIPTION],
		CommitCount:          data[NCI_COMMIT_COUNT],

		LastreleaseRefName:          data[NCI_LASTRELEASE_REF_NAME],
		LastreleaseRefSlug:          data[NCI_LASTRELEASE_REF_SLUG],
		LastreleaseRefVcs:           data[NCI_LASTRELEASE_REF_VCS],
		LastreleaseCommitAfterCount: data[NCI_LASTRELEASE_COMMIT_AFTER_COUNT],

		DeployFreeze: data[NCI_DEPLOY_FREEZE],
	}
}

func ToMap(spec NormalizeCISpec) map[string]string {
	data := make(map[string]string)
	data[NCI] = spec.Found
	data[NCI_VERSION] = spec.Version
	data[NCI_SERVICE_NAME] = spec.ServiceName
	data[NCI_SERVICE_SLUG] = spec.ServiceSlug

	data[NCI_WORKER_ID] = spec.WorkerId
	data[NCI_WORKER_NAME] = spec.WorkerName
	data[NCI_WORKER_TYPE] = spec.WorkerType
	data[NCI_WORKER_OS] = spec.WorkerOS
	data[NCI_WORKER_VERSION] = spec.WorkerVersion
	data[NCI_WORKER_ARCH] = spec.WorkerArch

	data[NCI_PIPELINE_ID] = spec.PipelineId
	data[NCI_PIPELINE_TRIGGER] = spec.PipelineTrigger
	data[NCI_PIPELINE_STAGE_ID] = spec.PipelineStageId
	data[NCI_PIPELINE_STAGE_NAME] = spec.PipelineStageName
	data[NCI_PIPELINE_STAGE_SLUG] = spec.PipelineStageSlug
	data[NCI_PIPELINE_JOB_ID] = spec.PipelineJobId
	data[NCI_PIPELINE_JOB_NAME] = spec.PipelineJobName
	data[NCI_PIPELINE_JOB_SLUG] = spec.PipelineJobSlug
	data[NCI_PIPELINE_JOB_STARTED_AT] = spec.PipelineJobStartedAt
	data[NCI_PIPELINE_ATTEMPT] = spec.PipelineAttempt
	data[NCI_PIPELINE_CONFIG_FILE] = spec.PipelineConfigFile
	data[NCI_PIPELINE_URL] = spec.PipelineUrl
	if spec.PipelineInput != nil {
		for k, v := range spec.PipelineInput {
			data[NCI_PIPELINE_INPUT+"_"+k] = v
		}
	}

	data[NCI_MERGE_REQUEST_ID] = spec.MergeRequestId
	data[NCI_MERGE_REQUEST_SOURCE_BRANCH_NAME] = spec.MergeRequestSourceBranchName
	data[NCI_MERGE_REQUEST_TARGET_BRANCH_NAME] = spec.MergeRequestTargetBranchName

	data[NCI_PROJECT_ID] = spec.ProjectId
	data[NCI_PROJECT_NAME] = spec.ProjectName
	data[NCI_PROJECT_PATH] = spec.ProjectPath
	data[NCI_PROJECT_SLUG] = spec.ProjectSlug
	data[NCI_PROJECT_DESCRIPTION] = spec.ProjectDescription
	data[NCI_PROJECT_TOPICS] = spec.ProjectTopics
	data[NCI_PROJECT_ISSUE_URL] = spec.ProjectIssueUrl
	data[NCI_PROJECT_STARGAZERS] = spec.ProjectStargazers
	data[NCI_PROJECT_FORKS] = spec.ProjectForks
	data[NCI_PROJECT_DEFAULT_BRANCH] = spec.ProjectDefaultBranch
	data[NCI_PROJECT_URL] = spec.ProjectUrl
	data[NCI_PROJECT_DIR] = spec.ProjectDir

	data[NCI_CONTAINERREGISTRY_HOST] = spec.ContainerRegistryHost
	data[NCI_CONTAINERREGISTRY_USERNAME] = spec.ContainerRegistryUsername
	data[NCI_CONTAINERREGISTRY_PASSWORD] = spec.ContainerRegistryPassword
	data[NCI_CONTAINERREGISTRY_REPOSITORY] = spec.ContainerRegistryRepository
	data[NCI_CONTAINERREGISTRY_TAG] = spec.ContainerRegistryTag

	data[NCI_REPOSITORY_KIND] = spec.RepositoryKind
	data[NCI_REPOSITORY_REMOTE] = spec.RepositoryRemote
	data[NCI_REPOSITORY_HOST_SERVER] = spec.RepositoryHostServer
	data[NCI_REPOSITORY_HOST_TYPE] = spec.RepositoryHostType
	data[NCI_REPOSITORY_STATUS] = spec.RepositoryStatus
	data[NCI_COMMIT_REF_TYPE] = spec.CommitRefType
	data[NCI_COMMIT_REF_NAME] = spec.CommitRefName
	data[NCI_COMMIT_REF_PATH] = spec.CommitRefPath
	data[NCI_COMMIT_REF_SLUG] = spec.CommitRefSlug
	data[NCI_COMMIT_REF_VCS] = spec.CommitRefVcs
	data[NCI_COMMIT_REF_RELEASE] = spec.CommitRefRelease
	data[NCI_COMMIT_SHA] = spec.CommitSha
	data[NCI_COMMIT_SHA_SHORT] = spec.CommitShaShort
	data[NCI_COMMIT_AUTHOR_NAME] = spec.CommitAuthorName
	data[NCI_COMMIT_AUTHOR_EMAIL] = spec.CommitAuthorEmail
	data[NCI_COMMIT_COMMITTER_NAME] = spec.CommitCommitterName
	data[NCI_COMMIT_COMMITTER_EMAIL] = spec.CommitCommitterEmail
	data[NCI_COMMIT_TITLE] = spec.CommitTitle
	data[NCI_COMMIT_DESCRIPTION] = spec.CommitDescription
	data[NCI_COMMIT_COUNT] = spec.CommitCount

	data[NCI_LASTRELEASE_REF_NAME] = spec.LastreleaseRefName
	data[NCI_LASTRELEASE_REF_SLUG] = spec.LastreleaseRefSlug
	data[NCI_LASTRELEASE_REF_VCS] = spec.LastreleaseRefVcs
	data[NCI_LASTRELEASE_COMMIT_AFTER_COUNT] = spec.LastreleaseCommitAfterCount

	data[NCI_DEPLOY_FREEZE] = spec.DeployFreeze

	return data
}

// Validate will validate the spec for completion
func (spec *NormalizeCISpec) Validate() []validationError {
	return validateSpec(spec)
}
