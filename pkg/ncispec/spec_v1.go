package ncispec

const NCI = "NCI"                           // Will be set the true, if the variables have been normalized. (this script)
const NCI_VERSION = "NCI_VERSION"           // The revision of nci that was used to generate the normalized variables.
const NCI_SERVICE_NAME = "NCI_SERVICE_NAME" // The commercial name of the used ci service. (e.g. GitLab CI, Travis CI, CircleCI, Jenkins)
const NCI_SERVICE_SLUG = "NCI_SERVICE_SLUG" // The commercial name normalized as slug for use in scripts, will not be changed.

const NCI_WORKER_ID = "NCI_WORKER_ID"           // A unique id of the ci worker.
const NCI_WORKER_NAME = "NCI_WORKER_NAME"       // The human readable name of the ci worker.
const NCI_WORKER_VERSION = "NCI_WORKER_VERSION" // The version of the ci worker.
const NCI_WORKER_ARCH = "NCI_WORKER_ARCH"       // The arch of the ci worker. (ie. linux/amd64)

const NCI_PIPELINE_TRIGGER = "NCI_PIPELINE_TRIGGER"                 // What triggered the pipeline. (ie. manual/push/trigger/api/schedule/pull_request/build)
const NCI_PIPELINE_STAGE_NAME = "NCI_PIPELINE_STAGE_NAME"           // Human readable name of the current stage.
const NCI_PIPELINE_STAGE_SLUG = "NCI_PIPELINE_STAGE_SLUG"           // Slug of the current stage.
const NCI_PIPELINE_JOB_NAME = "NCI_PIPELINE_JOB_NAME"               // Human readable name of the current job.
const NCI_PIPELINE_JOB_SLUG = "NCI_PIPELINE_JOB_SLUG"               // Slug of the current job.
const NCI_PIPELINE_PULL_REQUEST_ID = "NCI_PIPELINE_PULL_REQUEST_ID" // The number of the pull request, is only present if `NCI_PIPELINE_TRIGGER` = pull_request.

const NCI_PROJECT_ID = "NCI_PROJECT_ID"                   // Unique project id, can be used in deployments.
const NCI_PROJECT_NAME = "NCI_PROJECT_NAME"               // Unique project id, can be used in deployments.
const NCI_PROJECT_PATH = "NCI_PROJECT_PATH"               // Path of the namespace and the project
const NCI_PROJECT_SLUG = "NCI_PROJECT_SLUG"               // Project slug, that can be used in deployments.
const NCI_PROJECT_DESCRIPTION = "NCI_PROJECT_DESCRIPTION" // The project description.
const NCI_PROJECT_TOPICS = "NCI_PROJECT_TOPICS"           // The topics / tags of the project.
const NCI_PROJECT_ISSUE_URL = "NCI_PROJECT_ISSUE_URL"     // A template for links to issues, contains a `{ID}` placeholder.
const NCI_PROJECT_STARGAZERS = "NCI_PROJECT_STARGAZERS"   // The number of people who `follow` / `bookmarked` the project.
const NCI_PROJECT_FORKS = "NCI_PROJECT_FORKS"             // The number of forks of the project.
const NCI_PROJECT_DIR = "NCI_PROJECT_DIR"                 // Project directory on the local filesystem.

const NCI_CONTAINERREGISTRY_HOST = "NCI_CONTAINERREGISTRY_HOST"             // The hostname of the container registry.
const NCI_CONTAINERREGISTRY_USERNAME = "NCI_CONTAINERREGISTRY_USERNAME"     // The username used for container registry authentication.
const NCI_CONTAINERREGISTRY_PASSWORD = "NCI_CONTAINERREGISTRY_PASSWORD"     // The password used for container registry authentication.
const NCI_CONTAINERREGISTRY_REPOSITORY = "NCI_CONTAINERREGISTRY_REPOSITORY" // The repository, that should be used for the current project.
const NCI_CONTAINERREGISTRY_TAG = "NCI_CONTAINERREGISTRY_TAG"               // The tag that should be build.

const NCI_REPOSITORY_KIND = "NCI_REPOSITORY_KIND"               //  The used version control system. (git)
const NCI_REPOSITORY_REMOTE = "NCI_REPOSITORY_REMOTE"           // The remote url pointing at the repository. (git remote url or `local` if no remote was found)
const NCI_COMMIT_REF_TYPE = "NCI_COMMIT_REF_TYPE"               // The reference type. (branch / tag)
const NCI_COMMIT_REF_NAME = "NCI_COMMIT_REF_NAME"               // Human readable name of the current repository reference.
const NCI_COMMIT_REF_PATH = "NCI_COMMIT_REF_PATH"               // Combination of the ref type and ref name. (tag/v1.0.0 or branch/main)
const NCI_COMMIT_REF_SLUG = "NCI_COMMIT_REF_SLUG"               // Slug of the current repository reference.
const NCI_COMMIT_REF_VCS = "NCI_COMMIT_REF_VCS"                 // Holds the vcs specific absolute reference name. (ex: `refs/heads/main`)
const NCI_COMMIT_REF_RELEASE = "NCI_COMMIT_REF_RELEASE"         // Release version of the artifact, without leading `v` or `/` - should be in format `x.y.z` or `feature-abc`.
const NCI_COMMIT_SHA = "NCI_COMMIT_SHA"                         // A unique hash, that each commit gets.
const NCI_COMMIT_SHA_SHORT = "NCI_COMMIT_SHA_SHORT"             // A short form of the unique commit hash. (8 chars)
const NCI_COMMIT_AUTHOR_NAME = "NCI_COMMIT_AUTHOR_NAME"         // author name
const NCI_COMMIT_AUTHOR_EMAIL = "NCI_COMMIT_AUTHOR_EMAIL"       // author email
const NCI_COMMIT_COMMITTER_NAME = "NCI_COMMIT_COMMITTER_NAME"   // committer name
const NCI_COMMIT_COMMITTER_EMAIL = "NCI_COMMIT_COMMITTER_EMAIL" // committer email
const NCI_COMMIT_TITLE = "NCI_COMMIT_TITLE"                     // The title of the latest commit on the current reference.
const NCI_COMMIT_DESCRIPTION = "NCI_COMMIT_DESCRIPTION"         // The description of the latest commit on the current reference.
const NCI_COMMIT_COUNT = "NCI_COMMIT_COUNT"                     // The total amount of commits inside of the current reference, can be used as build number.

const NCI_LASTRELEASE_REF_NAME = "NCI_LASTRELEASE_REF_NAME"                     // Human readable name of the last stable release.
const NCI_LASTRELEASE_REF_SLUG = "NCI_LASTRELEASE_REF_SLUG"                     // Slug of the last stable release.
const NCI_LASTRELEASE_REF_VCS = "NCI_LASTRELEASE_REF_VCS"                       // Holds the vcs specific absolute reference name of the last stable release. (ex: `refs/heads/main`)
const NCI_LASTRELEASE_COMMIT_AFTER_COUNT = "NCI_LASTRELEASE_COMMIT_AFTER_COUNT" // Holds the count of commits since the last stable release.

const NCI_DEPLOY_FREEZE = "NCI_DEPLOY_FREEZE" // Currently in a deploy freeze window? (`true`, `false`)

type NormalizeCISpec struct {
	DATA map[string]string // storage to hold any data unrelated to NCI

	NCI              string `validate:"required"`         // Will be set the true, if the variables have been normalized. (this script)
	NCI_VERSION      string `validate:"required"`         // The revision of nci that was used to generate the normalized variables.
	NCI_SERVICE_NAME string `validate:"required"`         // The commercial name of the used ci service. (e.g. GitLab CI, Travis CI, CircleCI, Jenkins)
	NCI_SERVICE_SLUG string `validate:"required,is-slug"` // The commercial name normalized as slug for use in scripts, will not be changed.

	NCI_WORKER_ID      string `validate:"required"`         // A unique id of the ci worker.
	NCI_WORKER_NAME    string `validate:"required"`         // The human readable name of the ci worker.
	NCI_WORKER_VERSION string `validate:"required"`         // The version of the ci worker.
	NCI_WORKER_ARCH    string `validate:"required,is-arch"` // The arch of the ci worker. (ie. linux/amd64)

	NCI_PIPELINE_TRIGGER         string `validate:"required,oneof=cli manual push trigger api schedule pull_request build"` // What triggered the pipeline. (ie. manual/push/trigger/api/schedule/pull_request/build)
	NCI_PIPELINE_STAGE_NAME      string `validate:"required"`                                                               // Human readable name of the current stage.
	NCI_PIPELINE_STAGE_SLUG      string `validate:"required,is-slug"`                                                       // Slug of the current stage.
	NCI_PIPELINE_JOB_NAME        string `validate:"required"`                                                               // Human readable name of the current job.
	NCI_PIPELINE_JOB_SLUG        string `validate:"required,is-slug"`                                                       // Slug of the current job.
	NCI_PIPELINE_PULL_REQUEST_ID string `validate:"required_if=NCI_PIPELINE_TRIGGER pull_request"`                          // The number of the pull request, is only present if `NCI_PIPELINE_TRIGGER` = pull_request.

	NCI_PROJECT_ID          string // Unique project id, can be used in deployments.
	NCI_PROJECT_NAME        string // Unique project id, can be used in deployments.
	NCI_PROJECT_PATH        string // Path of the Namespace and the project
	NCI_PROJECT_SLUG        string `validate:"required,is-slug"` // Project slug, that can be used in deployments.
	NCI_PROJECT_DESCRIPTION string // The project description.
	NCI_PROJECT_TOPICS      string // The topics / tags of the project.
	NCI_PROJECT_ISSUE_URL   string // A template for links to issues, contains a `{ID}` placeholder.
	NCI_PROJECT_STARGAZERS  string `validate:"number"` // The number of people who `follow` / `bookmarked` the project.
	NCI_PROJECT_FORKS       string `validate:"number"` // The number of forks of the project.
	NCI_PROJECT_DIR         string // Project directory on the local filesystem.

	NCI_CONTAINERREGISTRY_HOST       string // The hostname of the container registry.
	NCI_CONTAINERREGISTRY_USERNAME   string // The username used for container registry authentication.
	NCI_CONTAINERREGISTRY_PASSWORD   string // The password used for container registry authentication.
	NCI_CONTAINERREGISTRY_REPOSITORY string `validate:"required"` // The repository, that should be used for the current project.
	NCI_CONTAINERREGISTRY_TAG        string `validate:"required"` // The tag that should be build.

	NCI_REPOSITORY_KIND        string `validate:"required"` //  The used version control system. (git)
	NCI_REPOSITORY_REMOTE      string `validate:"required"` // The remote url pointing at the repository. (git remote url or `local` if no remote was found)
	NCI_COMMIT_REF_TYPE        string `validate:"required"` // The reference type. (branch / tag)
	NCI_COMMIT_REF_NAME        string `validate:"required"` // Human readable name of the current repository reference.
	NCI_COMMIT_REF_PATH        string `validate:"required"` // Combination of the ref type and ref name. (tag/v1.0.0 or branch/main)
	NCI_COMMIT_REF_SLUG        string `validate:"required"` // Slug of the current repository reference.
	NCI_COMMIT_REF_VCS         string `validate:"required"` // Holds the vcs specific absolute reference name. (ex: `refs/heads/main`)// Release version of the artifact, without leading `v` or `/` - should be in format `x.y.z` or `feature-abc`.
	NCI_COMMIT_REF_RELEASE     string `validate:"required"` // Release version of the artifact, without leading `v` or `/` - should be in format `x.y.z` or `feature-abc`.
	NCI_COMMIT_SHA             string `validate:"required"` // A unique hash, that each commit gets.
	NCI_COMMIT_SHA_SHORT       string `validate:"required"` // A short form of the unique commit hash. (8 chars)
	NCI_COMMIT_AUTHOR_NAME     string `validate:"required"` // author name
	NCI_COMMIT_AUTHOR_EMAIL    string `validate:"required"` // author email
	NCI_COMMIT_COMMITTER_NAME  string `validate:"required"` // committer name
	NCI_COMMIT_COMMITTER_EMAIL string `validate:"required"` // committer email
	NCI_COMMIT_TITLE           string `validate:"required"` // The title of the latest commit on the current reference.
	NCI_COMMIT_DESCRIPTION     string // The description of the latest commit on the current reference.
	NCI_COMMIT_COUNT           string `validate:"required"` // The total amount of commits inside of the current reference, can be used as build number.

	NCI_LASTRELEASE_REF_NAME           string `validate:"required"` // Human readable name of the last stable release.
	NCI_LASTRELEASE_REF_SLUG           string `validate:"required"` // Slug of the last stable release.
	NCI_LASTRELEASE_REF_VCS            string `validate:"required"` // Holds the vcs specific absolute reference name of the last stable release. (ex: `refs/heads/main`)
	NCI_LASTRELEASE_COMMIT_AFTER_COUNT string `validate:"required"` // Holds the count of commits since the last stable release.

	NCI_DEPLOY_FREEZE string `validate:"required,boolean"` // Currently in a deploy freeze window? (`true`, `false`)
}

func OfMap(data map[string]string) NormalizeCISpec {
	return NormalizeCISpec{
		DATA:                               data,
		NCI:                                data[NCI],
		NCI_VERSION:                        data[NCI_VERSION],
		NCI_SERVICE_NAME:                   data[NCI_SERVICE_NAME],
		NCI_SERVICE_SLUG:                   data[NCI_SERVICE_SLUG],
		NCI_WORKER_ID:                      data[NCI_WORKER_ID],
		NCI_WORKER_NAME:                    data[NCI_WORKER_NAME],
		NCI_WORKER_VERSION:                 data[NCI_WORKER_VERSION],
		NCI_WORKER_ARCH:                    data[NCI_WORKER_ARCH],
		NCI_PIPELINE_TRIGGER:               data[NCI_PIPELINE_TRIGGER],
		NCI_PIPELINE_STAGE_NAME:            data[NCI_PIPELINE_STAGE_NAME],
		NCI_PIPELINE_STAGE_SLUG:            data[NCI_PIPELINE_STAGE_SLUG],
		NCI_PIPELINE_JOB_NAME:              data[NCI_PIPELINE_JOB_NAME],
		NCI_PIPELINE_JOB_SLUG:              data[NCI_PIPELINE_JOB_SLUG],
		NCI_PIPELINE_PULL_REQUEST_ID:       data[NCI_PIPELINE_PULL_REQUEST_ID],
		NCI_PROJECT_ID:                     data[NCI_PROJECT_ID],
		NCI_PROJECT_NAME:                   data[NCI_PROJECT_NAME],
		NCI_PROJECT_PATH:                   data[NCI_PROJECT_PATH],
		NCI_PROJECT_SLUG:                   data[NCI_PROJECT_SLUG],
		NCI_PROJECT_DESCRIPTION:            data[NCI_PROJECT_DESCRIPTION],
		NCI_PROJECT_TOPICS:                 data[NCI_PROJECT_TOPICS],
		NCI_PROJECT_ISSUE_URL:              data[NCI_PROJECT_ISSUE_URL],
		NCI_PROJECT_STARGAZERS:             data[NCI_PROJECT_STARGAZERS],
		NCI_PROJECT_FORKS:                  data[NCI_PROJECT_FORKS],
		NCI_PROJECT_DIR:                    data[NCI_PROJECT_DIR],
		NCI_CONTAINERREGISTRY_HOST:         data[NCI_CONTAINERREGISTRY_HOST],
		NCI_CONTAINERREGISTRY_USERNAME:     data[NCI_CONTAINERREGISTRY_USERNAME],
		NCI_CONTAINERREGISTRY_PASSWORD:     data[NCI_CONTAINERREGISTRY_PASSWORD],
		NCI_CONTAINERREGISTRY_REPOSITORY:   data[NCI_CONTAINERREGISTRY_REPOSITORY],
		NCI_CONTAINERREGISTRY_TAG:          data[NCI_CONTAINERREGISTRY_TAG],
		NCI_REPOSITORY_KIND:                data[NCI_REPOSITORY_KIND],
		NCI_REPOSITORY_REMOTE:              data[NCI_REPOSITORY_REMOTE],
		NCI_COMMIT_REF_TYPE:                data[NCI_COMMIT_REF_TYPE],
		NCI_COMMIT_REF_NAME:                data[NCI_COMMIT_REF_NAME],
		NCI_COMMIT_REF_PATH:                data[NCI_COMMIT_REF_PATH],
		NCI_COMMIT_REF_SLUG:                data[NCI_COMMIT_REF_SLUG],
		NCI_COMMIT_REF_VCS:                 data[NCI_COMMIT_REF_VCS],
		NCI_COMMIT_REF_RELEASE:             data[NCI_COMMIT_REF_RELEASE],
		NCI_COMMIT_SHA:                     data[NCI_COMMIT_SHA],
		NCI_COMMIT_SHA_SHORT:               data[NCI_COMMIT_SHA_SHORT],
		NCI_COMMIT_AUTHOR_NAME:             data[NCI_COMMIT_AUTHOR_NAME],
		NCI_COMMIT_AUTHOR_EMAIL:            data[NCI_COMMIT_AUTHOR_EMAIL],
		NCI_COMMIT_COMMITTER_NAME:          data[NCI_COMMIT_COMMITTER_NAME],
		NCI_COMMIT_COMMITTER_EMAIL:         data[NCI_COMMIT_COMMITTER_EMAIL],
		NCI_COMMIT_TITLE:                   data[NCI_COMMIT_TITLE],
		NCI_COMMIT_DESCRIPTION:             data[NCI_COMMIT_DESCRIPTION],
		NCI_COMMIT_COUNT:                   data[NCI_COMMIT_COUNT],
		NCI_LASTRELEASE_REF_NAME:           data[NCI_LASTRELEASE_REF_NAME],
		NCI_LASTRELEASE_REF_SLUG:           data[NCI_LASTRELEASE_REF_SLUG],
		NCI_LASTRELEASE_REF_VCS:            data[NCI_LASTRELEASE_REF_VCS],
		NCI_LASTRELEASE_COMMIT_AFTER_COUNT: data[NCI_LASTRELEASE_COMMIT_AFTER_COUNT],
		NCI_DEPLOY_FREEZE:                  data[NCI_DEPLOY_FREEZE],
	}
}

func ToMap(spec NormalizeCISpec) map[string]string {
	data := spec.DATA
	if data == nil {
		data = make(map[string]string)
	}

	data[NCI] = spec.NCI
	data[NCI_VERSION] = spec.NCI_VERSION
	data[NCI_SERVICE_NAME] = spec.NCI_SERVICE_NAME
	data[NCI_SERVICE_SLUG] = spec.NCI_SERVICE_SLUG

	data[NCI_WORKER_ID] = spec.NCI_WORKER_ID
	data[NCI_WORKER_NAME] = spec.NCI_WORKER_NAME
	data[NCI_WORKER_VERSION] = spec.NCI_WORKER_VERSION
	data[NCI_WORKER_ARCH] = spec.NCI_WORKER_ARCH

	data[NCI_PIPELINE_TRIGGER] = spec.NCI_PIPELINE_TRIGGER
	data[NCI_PIPELINE_STAGE_NAME] = spec.NCI_PIPELINE_STAGE_NAME
	data[NCI_PIPELINE_STAGE_SLUG] = spec.NCI_PIPELINE_STAGE_SLUG
	data[NCI_PIPELINE_JOB_NAME] = spec.NCI_PIPELINE_JOB_NAME
	data[NCI_PIPELINE_JOB_SLUG] = spec.NCI_PIPELINE_JOB_SLUG
	data[NCI_PIPELINE_PULL_REQUEST_ID] = spec.NCI_PIPELINE_PULL_REQUEST_ID

	data[NCI_PROJECT_ID] = spec.NCI_PROJECT_ID
	data[NCI_PROJECT_NAME] = spec.NCI_PROJECT_NAME
	data[NCI_PROJECT_PATH] = spec.NCI_PROJECT_PATH
	data[NCI_PROJECT_SLUG] = spec.NCI_PROJECT_SLUG
	data[NCI_PROJECT_DESCRIPTION] = spec.NCI_PROJECT_DESCRIPTION
	data[NCI_PROJECT_TOPICS] = spec.NCI_PROJECT_TOPICS
	data[NCI_PROJECT_ISSUE_URL] = spec.NCI_PROJECT_ISSUE_URL
	data[NCI_PROJECT_STARGAZERS] = spec.NCI_PROJECT_STARGAZERS
	data[NCI_PROJECT_FORKS] = spec.NCI_PROJECT_FORKS
	data[NCI_PROJECT_DIR] = spec.NCI_PROJECT_DIR

	data[NCI_CONTAINERREGISTRY_HOST] = spec.NCI_CONTAINERREGISTRY_HOST
	data[NCI_CONTAINERREGISTRY_USERNAME] = spec.NCI_CONTAINERREGISTRY_USERNAME
	data[NCI_CONTAINERREGISTRY_PASSWORD] = spec.NCI_CONTAINERREGISTRY_PASSWORD
	data[NCI_CONTAINERREGISTRY_REPOSITORY] = spec.NCI_CONTAINERREGISTRY_REPOSITORY
	data[NCI_CONTAINERREGISTRY_TAG] = spec.NCI_CONTAINERREGISTRY_TAG

	data[NCI_REPOSITORY_KIND] = spec.NCI_REPOSITORY_KIND
	data[NCI_REPOSITORY_REMOTE] = spec.NCI_REPOSITORY_REMOTE
	data[NCI_COMMIT_REF_TYPE] = spec.NCI_COMMIT_REF_TYPE
	data[NCI_COMMIT_REF_NAME] = spec.NCI_COMMIT_REF_NAME
	data[NCI_COMMIT_REF_PATH] = spec.NCI_COMMIT_REF_PATH
	data[NCI_COMMIT_REF_SLUG] = spec.NCI_COMMIT_REF_SLUG
	data[NCI_COMMIT_REF_VCS] = spec.NCI_COMMIT_REF_VCS
	data[NCI_COMMIT_REF_RELEASE] = spec.NCI_COMMIT_REF_RELEASE
	data[NCI_COMMIT_SHA] = spec.NCI_COMMIT_SHA
	data[NCI_COMMIT_SHA_SHORT] = spec.NCI_COMMIT_SHA_SHORT
	data[NCI_COMMIT_AUTHOR_NAME] = spec.NCI_COMMIT_AUTHOR_NAME
	data[NCI_COMMIT_AUTHOR_EMAIL] = spec.NCI_COMMIT_AUTHOR_EMAIL
	data[NCI_COMMIT_COMMITTER_NAME] = spec.NCI_COMMIT_COMMITTER_NAME
	data[NCI_COMMIT_COMMITTER_EMAIL] = spec.NCI_COMMIT_COMMITTER_EMAIL
	data[NCI_COMMIT_TITLE] = spec.NCI_COMMIT_TITLE
	data[NCI_COMMIT_DESCRIPTION] = spec.NCI_COMMIT_DESCRIPTION
	data[NCI_COMMIT_COUNT] = spec.NCI_COMMIT_COUNT

	data[NCI_LASTRELEASE_REF_NAME] = spec.NCI_LASTRELEASE_REF_NAME
	data[NCI_LASTRELEASE_REF_SLUG] = spec.NCI_LASTRELEASE_REF_SLUG
	data[NCI_LASTRELEASE_REF_VCS] = spec.NCI_LASTRELEASE_REF_VCS
	data[NCI_LASTRELEASE_COMMIT_AFTER_COUNT] = spec.NCI_LASTRELEASE_COMMIT_AFTER_COUNT

	data[NCI_DEPLOY_FREEZE] = spec.NCI_DEPLOY_FREEZE

	return data
}

// Validate will validate the spec for completion
func (spec *NormalizeCISpec) Validate() []validationError {
	return validateSpec(spec)
}
