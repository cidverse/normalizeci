package v1

type Spec struct {
	Found        string `env:"NCI" validate:"required"`                      // Will be set the true, if the variables have been normalized. (this script)
	Version      string `env:"NCI_VERSION" validate:"required"`              // The revision of nci that was used to generate the normalized variables.
	ServiceName  string `env:"NCI_SERVICE_NAME" validate:"required"`         // The commercial name of the used ci service. (e.g. GitLab CI, Travis CI, CircleCI, Jenkins)
	ServiceSlug  string `env:"NCI_SERVICE_SLUG" validate:"required,is-slug"` // The commercial name normalized as slug for use in scripts, will not be changed.
	Worker       Worker
	Pipeline     Pipeline
	Repository   Repository
	Project      Project
	Commit       Commit
	MergeRequest MergeRequest
	Flags        Flags
}

type Worker struct {
	Id      string `env:"NCI_WORKER_ID" validate:"required"`   // A unique id of the ci worker.
	Name    string `env:"NCI_WORKER_NAME" validate:"required"` // The human-readable name of the ci worker.
	Type    string `env:"NCI_WORKER_TYPE" validate:"required"`
	OS      string `env:"NCI_WORKER_OS"`                               // Worker OS or OS Image
	Version string `env:"NCI_WORKER_VERSION" validate:"required"`      // The version of the ci worker.
	Arch    string `env:"NCI_WORKER_ARCH" validate:"required,is-arch"` // The arch of the ci worker. (ie. linux/amd64)
}

type Pipeline struct {
	Id           string            `env:"NCI_PIPELINE_ID" validate:"required"`
	Trigger      string            `env:"NCI_PIPELINE_TRIGGER" validate:"required,oneof=cli manual push trigger api schedule pull_request build"` // What triggered the pipeline. (ie. manual/push/trigger/api/schedule/pull_request/build)
	StageId      string            `env:"NCI_PIPELINE_STAGE_ID"`
	StageName    string            `env:"NCI_PIPELINE_STAGE_NAME" validate:"required"`         // Human-readable name of the current stage.
	StageSlug    string            `env:"NCI_PIPELINE_STAGE_SLUG" validate:"required,is-slug"` // Slug of the current stage.
	JobId        string            `env:"NCI_PIPELINE_JOB_ID"`
	JobName      string            `env:"NCI_PIPELINE_JOB_NAME" validate:"required"`         // Human-readable name of the current job.
	JobSlug      string            `env:"NCI_PIPELINE_JOB_SLUG" validate:"required,is-slug"` // Slug of the current job.
	JobStartedAt string            `env:"NCI_PIPELINE_JOB_STARTED_AT" validate:"required"`   // Timestamp when the job started.
	Attempt      string            `env:"NCI_PIPELINE_ATTEMPT" validate:"number"`            // The current attempt number of the pipeline.
	ConfigFile   string            `env:"NCI_PIPELINE_CONFIG_FILE"`                          // Pipeline Config File
	Url          string            `env:"NCI_PIPELINE_URL"`                                  // Pipeline URL
	Input        map[string]string `env-prefix:"NCI_INPUT_"`
}

type Repository struct {
	Kind           string `env:"NCI_REPOSITORY_KIND" validate:"required"`                     //  The used version control system. (git)
	Remote         string `env:"NCI_REPOSITORY_REMOTE" validate:"required"`                   // The remote url pointing at the repository. (git remote url or `local` if no remote was found)
	HostServer     string `env:"NCI_REPOSITORY_HOST_SERVER" validate:"required"`              // Host of the git repository server, for example github.com
	HostServerSlug string `env:"NCI_REPOSITORY_HOST_SERVER_SLUG" validate:"required,is-slug"` // Host of the git repository server, for example github-com
	HostType       string `env:"NCI_REPOSITORY_HOST_TYPE" validate:"required"`                // Type of the git repository server (github, gitlab, ...)
	Status         string `env:"NCI_REPOSITORY_STATUS" validate:"required"`                   // The repository status (dirty, clean)
}

type Project struct {
	UID           string `env:"NCI_PROJECT_UID" validate:"required,is-slug"`  // UID returns a unique identifier by combining the host slug and project id. (e.g. github-com-123456)
	ID            string `env:"NCI_PROJECT_ID" validate:"required"`           // Unique project id, can be used in deployments.
	Name          string `env:"NCI_PROJECT_NAME" validate:"required"`         // Unique project id, can be used in deployments.
	Path          string `env:"NCI_PROJECT_PATH" validate:"required"`         // Path of the Namespace and the project
	Slug          string `env:"NCI_PROJECT_SLUG" validate:"required,is-slug"` // Project slug, that can be used in deployments.
	Description   string `env:"NCI_PROJECT_DESCRIPTION"`                      // The project description.
	Topics        string `env:"NCI_PROJECT_TOPICS"`                           // The topics / tags of the project.
	IssueUrl      string `env:"NCI_PROJECT_ISSUE_URL"`                        // A template for links to issues, contains a `{ID}` placeholder.
	Stargazers    string `env:"NCI_PROJECT_STARGAZERS"`                       // The number of people who `follow` / `bookmarked` the project.
	Forks         string `env:"NCI_PROJECT_FORKS"`                            // The number of forks of the project.
	Dir           string `env:"NCI_PROJECT_DIR" validate:"required"`          // Project directory on the local filesystem.
	Url           string `env:"NCI_PROJECT_URL"`                              // Project URL
	DefaultBranch string `env:"NCI_PROJECT_DEFAULT_BRANCH"`                   // The default branch
}

type Commit struct {
	RefType        string `env:"NCI_COMMIT_REF_TYPE" validate:"required"`        // The reference type. (branch / tag)
	RefName        string `env:"NCI_COMMIT_REF_NAME" validate:"required"`        // Human-readable name of the current repository reference.
	RefPath        string `env:"NCI_COMMIT_REF_PATH" validate:"required"`        // Combination of the ref type and ref name. (tag/v1.0.0 or branch/main)
	RefSlug        string `env:"NCI_COMMIT_REF_SLUG" validate:"required"`        // Slug of the current repository reference.
	RefVCS         string `env:"NCI_COMMIT_REF_VCS" validate:"required"`         // Holds the vcs specific absolute reference name. (ex: `refs/heads/main`)// Release version of the artifact, without leading `v` or `/` - should be in format `x.y.z` or `feature-abc`.
	RefRelease     string `env:"NCI_COMMIT_REF_RELEASE" validate:"required"`     // Release version of the artifact, without leading `v` or `/` - should be in format `x.y.z` or `feature-abc`.
	HashShort      string `env:"NCI_COMMIT_HASH_SHORT" validate:"required"`      // A short form of the unique commit hash. (8 chars)
	Hash           string `env:"NCI_COMMIT_HASH" validate:"required"`            //  A unique hash, that each commit gets.
	AuthorName     string `env:"NCI_COMMIT_AUTHOR_NAME" validate:"required"`     // author name
	AuthorEmail    string `env:"NCI_COMMIT_AUTHOR_EMAIL" validate:"required"`    // author email
	CommitterName  string `env:"NCI_COMMIT_COMMITTER_NAME" validate:"required"`  // committer name
	CommitterEmail string `env:"NCI_COMMIT_COMMITTER_EMAIL" validate:"required"` // committer email
	Title          string `env:"NCI_COMMIT_TITLE" validate:"required"`           // The title of the latest commit on the current reference.
	Description    string `env:"NCI_COMMIT_DESCRIPTION"`                         // The description of the latest commit on the current reference.
	Count          string `env:"NCI_COMMIT_COUNT" validate:"required"`           // The total amount of commits inside the current reference, can be used as build number.
}

type MergeRequest struct {
	Id               string `env:"NCI_MERGE_REQUEST_ID"`
	Title            string `env:"NCI_MERGE_REQUEST_TITLE"`
	SourceBranchName string `env:"NCI_MERGE_REQUEST_SOURCE_BRANCH_NAME"`
	SourceHash       string `env:"NCI_MERGE_REQUEST_SOURCE_HASH"`
	TargetBranchName string `env:"NCI_MERGE_REQUEST_TARGET_BRANCH_NAME"`
	TargetHash       string `env:"NCI_MERGE_REQUEST_TARGET_HASH"`

	/** extend with
	CI_MERGE_REQUEST_ID	11.6	all	The instance-level ID of the merge request. This is a unique ID across all projects on GitLab.
	CI_MERGE_REQUEST_IID	11.6	all	The project-level IID (internal ID) of the merge request. This ID is unique for the current project.

	CI_MERGE_REQUEST_LABELS	11.9	all	Comma-separated label names of the merge request.
	CI_MERGE_REQUEST_MILESTONE	11.9	all	The milestone title of the merge request.
	CI_MERGE_REQUEST_PROJECT_ID	11.6	all	The ID of the project of the merge request.
	*/
}

type Flags struct {
	DeployFreeze string `env:"NCI_DEPLOY_FREEZE"`
}

// Create creates a new Spec
func Create(serviceName string, serviceSlug string) Spec {
	return Spec{
		Found:        "true",
		Version:      "1.0.0",
		ServiceName:  serviceName,
		ServiceSlug:  serviceSlug,
		Worker:       Worker{},
		Pipeline:     Pipeline{},
		Repository:   Repository{},
		Project:      Project{},
		Commit:       Commit{},
		MergeRequest: MergeRequest{},
		Flags:        Flags{},
	}
}
