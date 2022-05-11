package gitlabci

import (
	"github.com/cidverse/normalizeci/pkg/common"
	"github.com/cidverse/normalizeci/pkg/ncispec"
	"github.com/cidverse/normalizeci/pkg/projectdetails"
	"github.com/cidverse/normalizeci/pkg/vcsrepository"
	"github.com/gosimple/slug"
	"runtime"
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
	if env["GITLAB_CI"] == "true" {
		return true
	}

	return false
}

// Normalize normalizes the environment variables into the common format
func (n Normalizer) Normalize(env map[string]string) map[string]string {
	nci := ncispec.OfMap(env)

	// common
	nci.NCI = "true"
	nci.NCI_VERSION = n.version
	nci.NCI_SERVICE_NAME = n.name
	nci.NCI_SERVICE_SLUG = n.slug

	// worker
	nci.NCI_WORKER_ID = env["CI_RUNNER_ID"]
	nci.NCI_WORKER_NAME = env["CI_RUNNER_DESCRIPTION"]
	nci.NCI_WORKER_VERSION = env["CI_RUNNER_VERSION"]
	nci.NCI_WORKER_ARCH = runtime.GOOS + "/" + runtime.GOARCH

	// pipeline
	nci.NCI_PIPELINE_TRIGGER = env["CI_PIPELINE_SOURCE"]
	if nci.NCI_PIPELINE_TRIGGER == ncispec.PipelineTriggerPullRequest {
		nci.NCI_PIPELINE_PULL_REQUEST_ID = env["CI_MERGE_REQUEST_IID"]
	}

	nci.NCI_PIPELINE_STAGE_NAME = env["CI_JOB_STAGE"]
	nci.NCI_PIPELINE_STAGE_SLUG = slug.Make(env["CI_JOB_STAGE"])
	nci.NCI_PIPELINE_JOB_NAME = env["CI_JOB_NAME"]
	nci.NCI_PIPELINE_JOB_SLUG = slug.Make(env["CI_JOB_NAME"])

	// repository
	projectDir := vcsrepository.FindRepositoryDirectory(common.GetWorkingDirectory())
	vcsData, addDataErr := vcsrepository.GetVCSRepositoryInformation(projectDir)
	if addDataErr != nil {
		panic(addDataErr)
	}
	nci.NCI_REPOSITORY_KIND = vcsData[ncispec.NCI_REPOSITORY_KIND]
	nci.NCI_REPOSITORY_REMOTE = vcsData[ncispec.NCI_REPOSITORY_REMOTE]
	if len(env["CI_COMMIT_TAG"]) > 0 {
		nci.NCI_COMMIT_REF_TYPE = "tag"
		nci.NCI_COMMIT_REF_NAME = env["CI_COMMIT_TAG"]
		nci.NCI_COMMIT_REF_PATH = nci.NCI_COMMIT_REF_TYPE + "/" + env["CI_COMMIT_TAG"]
		nci.NCI_COMMIT_REF_SLUG = slug.Make(env["CI_COMMIT_TAG"])
		nci.NCI_COMMIT_REF_VCS = "refs/tags/" + env["CI_COMMIT_TAG"]
	} else {
		nci.NCI_COMMIT_REF_TYPE = "branch"
		nci.NCI_COMMIT_REF_NAME = env["CI_COMMIT_REF_NAME"]
		nci.NCI_COMMIT_REF_PATH = nci.NCI_COMMIT_REF_TYPE + "/" + env["CI_COMMIT_REF_NAME"]
		nci.NCI_COMMIT_REF_SLUG = slug.Make(env["CI_COMMIT_REF_NAME"])
		nci.NCI_COMMIT_REF_VCS = "refs/heads/" + env["CI_COMMIT_REF_NAME"]
	}
	nci.NCI_COMMIT_REF_RELEASE = vcsData[ncispec.NCI_COMMIT_REF_RELEASE]
	nci.NCI_COMMIT_SHA = vcsData[ncispec.NCI_COMMIT_SHA]
	nci.NCI_COMMIT_SHA_SHORT = vcsData[ncispec.NCI_COMMIT_SHA_SHORT]
	nci.NCI_COMMIT_TITLE = vcsData[ncispec.NCI_COMMIT_TITLE]
	nci.NCI_COMMIT_DESCRIPTION = vcsData[ncispec.NCI_COMMIT_DESCRIPTION]
	nci.NCI_COMMIT_AUTHOR_NAME = vcsData[ncispec.NCI_COMMIT_AUTHOR_NAME]
	nci.NCI_COMMIT_AUTHOR_EMAIL = vcsData[ncispec.NCI_COMMIT_AUTHOR_EMAIL]
	nci.NCI_COMMIT_COMMITTER_NAME = vcsData[ncispec.NCI_COMMIT_COMMITTER_NAME]
	nci.NCI_COMMIT_COMMITTER_EMAIL = vcsData[ncispec.NCI_COMMIT_COMMITTER_EMAIL]
	nci.NCI_COMMIT_COUNT = vcsData[ncispec.NCI_COMMIT_COUNT]
	nci.NCI_LASTRELEASE_REF_NAME = vcsData[ncispec.NCI_LASTRELEASE_REF_NAME]
	nci.NCI_LASTRELEASE_REF_SLUG = vcsData[ncispec.NCI_LASTRELEASE_REF_SLUG]
	nci.NCI_LASTRELEASE_REF_VCS = vcsData[ncispec.NCI_LASTRELEASE_REF_VCS]
	nci.NCI_LASTRELEASE_COMMIT_AFTER_COUNT = vcsData[ncispec.NCI_LASTRELEASE_COMMIT_AFTER_COUNT]

	// project details
	projectData := projectdetails.GetProjectDetails(nci.NCI_REPOSITORY_KIND, nci.NCI_REPOSITORY_REMOTE)
	if projectData != nil {
		nci.NCI_PROJECT_ID = projectData[ncispec.NCI_PROJECT_ID]
		nci.NCI_PROJECT_NAME = projectData[ncispec.NCI_PROJECT_NAME]
		nci.NCI_PROJECT_PATH = projectData[ncispec.NCI_PROJECT_PATH]
		nci.NCI_PROJECT_SLUG = projectData[ncispec.NCI_PROJECT_SLUG]
		nci.NCI_PROJECT_DESCRIPTION = projectData[ncispec.NCI_PROJECT_DESCRIPTION]
		nci.NCI_PROJECT_TOPICS = projectData[ncispec.NCI_PROJECT_TOPICS]
		nci.NCI_PROJECT_ISSUE_URL = projectData[ncispec.NCI_PROJECT_ISSUE_URL]
		nci.NCI_PROJECT_STARGAZERS = projectData[ncispec.NCI_PROJECT_STARGAZERS]
		nci.NCI_PROJECT_FORKS = projectData[ncispec.NCI_PROJECT_FORKS]
	}
	nci.NCI_PROJECT_DIR = projectDir

	// container registry
	nci.NCI_CONTAINERREGISTRY_HOST = env["CI_REGISTRY"]
	nci.NCI_CONTAINERREGISTRY_REPOSITORY = env["CI_REGISTRY_IMAGE"]
	if len(env["CI_DEPLOY_USER"]) > 0 {
		nci.NCI_CONTAINERREGISTRY_USERNAME = env["CI_DEPLOY_USER"]
		nci.NCI_CONTAINERREGISTRY_PASSWORD = env["CI_DEPLOY_PASSWORD"]
	} else {
		nci.NCI_CONTAINERREGISTRY_USERNAME = env["CI_REGISTRY_USER"]
		nci.NCI_CONTAINERREGISTRY_PASSWORD = env["CI_REGISTRY_PASSWORD"]
	}
	nci.NCI_CONTAINERREGISTRY_TAG = nci.NCI_COMMIT_REF_RELEASE

	// control
	if _, ok := env["CI_DEPLOY_FREEZE"]; ok {
		nci.NCI_DEPLOY_FREEZE = env["CI_DEPLOY_FREEZE"]
	} else {
		nci.NCI_DEPLOY_FREEZE = "false"
	}

	return ncispec.ToMap(nci)
}

func (n Normalizer) Denormalize(env map[string]string) map[string]string {
	data := make(map[string]string)

	data["GITLAB_CI"] = "true"

	return data
}

// NewNormalizer gets a instance of the normalizer
func NewNormalizer() Normalizer {
	entity := Normalizer{
		version: "0.3.0",
		name:    "GitLab CI",
		slug:    "gitlab-ci",
	}

	return entity
}
