package gitlabci

import (
	_ "embed"
	"runtime"
	"testing"

	"github.com/cidverse/normalizeci/pkg/nciutil"
	"github.com/stretchr/testify/assert"
)

func TestNormalizer_Normalize_Common(t *testing.T) {
	nciutil.MockVCSClient(t)

	var normalizer = NewNormalizer()
	var normalized, err = normalizer.Normalize(map[string]string{})

	assert.NoError(t, err)
	assert.Equal(t, "true", normalized.Found)
	assert.Equal(t, "1.0.0", normalized.Version)
	assert.Equal(t, normalizer.name, normalized.ServiceName)
	assert.Equal(t, normalizer.slug, normalized.ServiceSlug)
}

func TestNormalizer_Normalize_Worker(t *testing.T) {
	nciutil.MockVCSClient(t)

	var normalizer = NewNormalizer()
	var normalized, err = normalizer.Normalize(map[string]string{
		"CI_RUNNER_ID":          "12270837",
		"CI_RUNNER_DESCRIPTION": "4-blue.shared.runners-manager.gitlab.com/default",
		"CI_RUNNER_VERSION":     "14.10.0~beta.50.g1f2fe53e",
	})

	assert.NoError(t, err)
	assert.Equal(t, "12270837", normalized.Worker.Id)
	assert.Equal(t, "4-blue.shared.runners-manager.gitlab.com/default", normalized.Worker.Name)
	assert.Equal(t, "gitlab_hosted_vm", normalized.Worker.Type)
	assert.Equal(t, "14.10.0~beta.50.g1f2fe53e", normalized.Worker.Version)
	assert.Equal(t, runtime.GOOS+"/"+runtime.GOARCH, normalized.Worker.Arch)
}

func TestNormalizer_Normalize_Pipeline(t *testing.T) {
	nciutil.MockVCSClient(t)

	var normalizer = NewNormalizer()
	var normalized, err = normalizer.Normalize(map[string]string{
		"CI_PIPELINE_ID":     "535898514",
		"CI_PIPELINE_SOURCE": "push",
		"CI_JOB_STAGE":       "build",
		"CI_JOB_ID":          "build",
		"CI_JOB_NAME":        "build",
		"CI_JOB_STARTED_AT":  "2022-05-10T20:20:01Z",
		"CI_JOB_URL":         "https://gitlab.com/cidverse/cienvsamples/-/jobs/2438765887",
	})

	assert.NoError(t, err)
	assert.Equal(t, "535898514", normalized.Pipeline.Id)
	assert.Equal(t, "push", normalized.Pipeline.Trigger)
	assert.Equal(t, "build", normalized.Pipeline.StageName)
	assert.Equal(t, "build", normalized.Pipeline.StageSlug)
	assert.Equal(t, "build", normalized.Pipeline.JobName)
	assert.Equal(t, "build", normalized.Pipeline.JobSlug)
	assert.Equal(t, "2022-05-10T20:20:01Z", normalized.Pipeline.JobStartedAt)
	assert.Equal(t, "1", normalized.Pipeline.Attempt)
	assert.Equal(t, "gitlab-ci.yml", normalized.Pipeline.ConfigFile)
	assert.Equal(t, "https://gitlab.com/cidverse/cienvsamples/-/jobs/2438765887", normalized.Pipeline.Url)
}

func TestNormalizer_Normalize_MergeRequest(t *testing.T) {
	nciutil.MockVCSClient(t)

	var normalizer = NewNormalizer()
	var normalized, err = normalizer.Normalize(map[string]string{
		"CI_MERGE_REQUEST_IID":                "153",
		"CI_MERGE_REQUEST_SOURCE_BRANCH_NAME": "feat/new-feature",
		"CI_MERGE_REQUEST_TARGET_BRANCH_NAME": "main",
	})

	assert.NoError(t, err)
	assert.Equal(t, "153", normalized.MergeRequest.Id)
	assert.Equal(t, "feat/new-feature", normalized.MergeRequest.SourceBranchName)
	assert.Equal(t, "main", normalized.MergeRequest.TargetBranchName)
}

func TestNormalizer_Normalize_Project(t *testing.T) {
	nciutil.MockVCSClient(t)

	var normalizer = NewNormalizer()
	var normalized, err = normalizer.Normalize(map[string]string{
		"CI_PROJECT_ID":          "35974876",
		"CI_PROJECT_TITLE":       "cienvsamples",
		"CI_PROJECT_NAME":        "cienvsamples",
		"CI_PROJECT_PATH_SLUG":   "cidverse-cienvsamples",
		"CI_PROJECT_DESCRIPTION": "A tool to turn the continuous integration / deployment variables into a common format for generally usable scripts without any dependencies.",
		"CI_DEFAULT_BRANCH":      "main",
		"CI_PROJECT_URL":         "https://gitlab.com/cidverse/cienvsamples",
	})

	assert.NoError(t, err)
	assert.Equal(t, "35974876", normalized.Project.Id)
	assert.Equal(t, "cienvsamples", normalized.Project.Name)
	assert.Equal(t, "cienvsamples", normalized.Project.Path)
	assert.Equal(t, "cidverse-cienvsamples", normalized.Project.Slug)
	assert.Equal(t, "A tool to turn the continuous integration / deployment variables into a common format for generally usable scripts without any dependencies.", normalized.Project.Description)
	assert.Equal(t, "main", normalized.Project.DefaultBranch)
	assert.Equal(t, "https://gitlab.com/cidverse/cienvsamples", normalized.Project.Url)
}

func TestNormalizer_Normalize_WorkflowAPI(t *testing.T) {

}
