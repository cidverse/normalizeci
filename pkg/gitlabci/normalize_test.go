package gitlabci

import (
	_ "embed"
	"runtime"
	"testing"

	"github.com/cidverse/normalizeci/pkg/vcsrepository"
	"github.com/stretchr/testify/assert"
)

func TestNormalizer_Normalize_Common(t *testing.T) {
	vcsrepository.MockClient = MockVCSClient(t)
	var normalizer = NewNormalizer()
	var normalized = normalizer.Normalize(map[string]string{})

	assert.Equal(t, "true", normalized.Found)
	assert.Equal(t, normalizer.version, normalized.Version)
	assert.Equal(t, normalizer.name, normalized.ServiceName)
	assert.Equal(t, normalizer.slug, normalized.ServiceSlug)
}

func TestNormalizer_Normalize_Worker(t *testing.T) {
	vcsrepository.MockClient = MockVCSClient(t)
	var normalizer = NewNormalizer()
	var normalized = normalizer.Normalize(map[string]string{
		"CI_RUNNER_ID":          "12270837",
		"CI_RUNNER_DESCRIPTION": "4-blue.shared.runners-manager.gitlab.com/default",
		"CI_RUNNER_VERSION":     "14.10.0~beta.50.g1f2fe53e",
	})

	assert.Equal(t, "12270837", normalized.WorkerId)
	assert.Equal(t, "4-blue.shared.runners-manager.gitlab.com/default", normalized.WorkerName)
	assert.Equal(t, "gitlab_hosted_vm", normalized.WorkerType)
	assert.Equal(t, "14.10.0~beta.50.g1f2fe53e", normalized.WorkerVersion)
	assert.Equal(t, runtime.GOOS+"/"+runtime.GOARCH, normalized.WorkerArch)
}

func TestNormalizer_Normalize_Pipeline(t *testing.T) {
	vcsrepository.MockClient = MockVCSClient(t)
	var normalizer = NewNormalizer()
	var normalized = normalizer.Normalize(map[string]string{
		"CI_PIPELINE_ID":     "535898514",
		"CI_PIPELINE_SOURCE": "push",
		"CI_JOB_STAGE":       "build",
		"CI_JOB_ID":          "build",
		"CI_JOB_NAME":        "build",
		"CI_JOB_STARTED_AT":  "2022-05-10T20:20:01Z",
		"CI_JOB_URL":         "https://gitlab.com/cidverse/cienvsamples/-/jobs/2438765887",
	})

	assert.Equal(t, "535898514", normalized.PipelineId)
	assert.Equal(t, "push", normalized.PipelineTrigger)
	assert.Equal(t, "build", normalized.PipelineStageName)
	assert.Equal(t, "build", normalized.PipelineStageSlug)
	assert.Equal(t, "build", normalized.PipelineJobName)
	assert.Equal(t, "build", normalized.PipelineJobSlug)
	assert.Equal(t, "2022-05-10T20:20:01Z", normalized.PipelineJobStartedAt)
	assert.Equal(t, "1", normalized.PipelineAttempt)
	assert.Equal(t, "gitlab-ci.yml", normalized.PipelineConfigFile)
	assert.Equal(t, "https://gitlab.com/cidverse/cienvsamples/-/jobs/2438765887", normalized.PipelineUrl)
}

func TestNormalizer_Normalize_MergeRequest(t *testing.T) {
	vcsrepository.MockClient = MockVCSClient(t)
	var normalizer = NewNormalizer()
	var normalized = normalizer.Normalize(map[string]string{
		"CI_MERGE_REQUEST_IID":                "153",
		"CI_MERGE_REQUEST_SOURCE_BRANCH_NAME": "feat/new-feature",
		"CI_MERGE_REQUEST_TARGET_BRANCH_NAME": "main",
	})

	assert.Equal(t, "153", normalized.MergeRequestId)
	assert.Equal(t, "feat/new-feature", normalized.MergeRequestSourceBranchName)
	assert.Equal(t, "main", normalized.MergeRequestTargetBranchName)
}

func TestNormalizer_Normalize_Project(t *testing.T) {
	vcsrepository.MockClient = MockVCSClient(t)
	var normalizer = NewNormalizer()
	var normalized = normalizer.Normalize(map[string]string{
		"CI_PROJECT_ID":          "35974876",
		"CI_PROJECT_TITLE":       "cienvsamples",
		"CI_PROJECT_NAME":        "cienvsamples",
		"CI_PROJECT_PATH_SLUG":   "cidverse-cienvsamples",
		"CI_PROJECT_DESCRIPTION": "A tool to turn the continuous integration / deployment variables into a common format for generally usable scripts without any dependencies.",
		"CI_DEFAULT_BRANCH":      "main",
		"CI_PROJECT_URL":         "https://gitlab.com/cidverse/cienvsamples",
	})

	assert.Equal(t, "35974876", normalized.ProjectId)
	assert.Equal(t, "cienvsamples", normalized.ProjectName)
	assert.Equal(t, "cienvsamples", normalized.ProjectPath)
	assert.Equal(t, "cidverse-cienvsamples", normalized.ProjectSlug)
	assert.Equal(t, "A tool to turn the continuous integration / deployment variables into a common format for generally usable scripts without any dependencies.", normalized.ProjectDescription)
	assert.Equal(t, "main", normalized.ProjectDefaultBranch)
	assert.Equal(t, "https://gitlab.com/cidverse/cienvsamples", normalized.ProjectUrl)
}

func TestNormalizer_Normalize_WorkflowAPI(t *testing.T) {

}
