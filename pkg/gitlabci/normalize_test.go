package gitlabci

import (
	_ "embed"
	"runtime"
	"testing"

	"github.com/cidverse/normalizeci/pkg/ncispec"
	"github.com/stretchr/testify/assert"
)

func TestNormalizer_Normalize_Common(t *testing.T) {
	var normalizer = NewNormalizer()
	var normalized = normalizer.Normalize(map[string]string{})

	assert.Equal(t, "true", normalized[ncispec.NCI])
	assert.Equal(t, normalizer.version, normalized[ncispec.NCI_VERSION])
	assert.Equal(t, normalizer.name, normalized[ncispec.NCI_SERVICE_NAME])
	assert.Equal(t, normalizer.slug, normalized[ncispec.NCI_SERVICE_SLUG])
}

func TestNormalizer_Normalize_Worker(t *testing.T) {
	var normalizer = NewNormalizer()
	var normalized = normalizer.Normalize(map[string]string{
		"CI_RUNNER_ID":          "12270837",
		"CI_RUNNER_DESCRIPTION": "4-blue.shared.runners-manager.gitlab.com/default",
		"CI_RUNNER_VERSION":     "14.10.0~beta.50.g1f2fe53e",
	})

	assert.Equal(t, "12270837", normalized[ncispec.NCI_WORKER_ID])
	assert.Equal(t, "4-blue.shared.runners-manager.gitlab.com/default", normalized[ncispec.NCI_WORKER_NAME])
	assert.Equal(t, "gitlab_hosted_vm", normalized[ncispec.NCI_WORKER_TYPE])
	assert.Equal(t, "14.10.0~beta.50.g1f2fe53e", normalized[ncispec.NCI_WORKER_VERSION])
	assert.Equal(t, runtime.GOOS+"/"+runtime.GOARCH, normalized[ncispec.NCI_WORKER_ARCH])
}

func TestNormalizer_Normalize_Pipeline(t *testing.T) {
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

	assert.Equal(t, "535898514", normalized[ncispec.NCI_PIPELINE_ID])
	assert.Equal(t, "push", normalized[ncispec.NCI_PIPELINE_TRIGGER])
	assert.Equal(t, "build", normalized[ncispec.NCI_PIPELINE_STAGE_NAME])
	assert.Equal(t, "build", normalized[ncispec.NCI_PIPELINE_STAGE_SLUG])
	assert.Equal(t, "build", normalized[ncispec.NCI_PIPELINE_JOB_NAME])
	assert.Equal(t, "build", normalized[ncispec.NCI_PIPELINE_JOB_SLUG])
	assert.Equal(t, "2022-05-10T20:20:01Z", normalized[ncispec.NCI_PIPELINE_JOB_STARTED_AT])
	assert.Equal(t, "1", normalized[ncispec.NCI_PIPELINE_ATTEMPT])
	assert.Equal(t, "gitlab-ci.yml", normalized[ncispec.NCI_PIPELINE_CONFIG_FILE])
	assert.Equal(t, "https://gitlab.com/cidverse/cienvsamples/-/jobs/2438765887", normalized[ncispec.NCI_PIPELINE_URL])
}

func TestNormalizer_Normalize_MergeRequest(t *testing.T) {
	var normalizer = NewNormalizer()
	var normalized = normalizer.Normalize(map[string]string{
		"CI_MERGE_REQUEST_IID":                "153",
		"CI_MERGE_REQUEST_SOURCE_BRANCH_NAME": "feat/new-feature",
		"CI_MERGE_REQUEST_TARGET_BRANCH_NAME": "main",
	})

	assert.Equal(t, "153", normalized[ncispec.NCI_MERGE_REQUEST_ID])
	assert.Equal(t, "feat/new-feature", normalized[ncispec.NCI_MERGE_REQUEST_SOURCE_BRANCH_NAME])
	assert.Equal(t, "main", normalized[ncispec.NCI_MERGE_REQUEST_TARGET_BRANCH_NAME])
}

func TestNormalizer_Normalize_Project(t *testing.T) {
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

	assert.Equal(t, "35974876", normalized[ncispec.NCI_PROJECT_ID])
	assert.Equal(t, "cienvsamples", normalized[ncispec.NCI_PROJECT_NAME])
	assert.Equal(t, "cienvsamples", normalized[ncispec.NCI_PROJECT_PATH])
	assert.Equal(t, "cidverse-cienvsamples", normalized[ncispec.NCI_PROJECT_SLUG])
	assert.Equal(t, "A tool to turn the continuous integration / deployment variables into a common format for generally usable scripts without any dependencies.", normalized[ncispec.NCI_PROJECT_DESCRIPTION])
	assert.Equal(t, "main", normalized[ncispec.NCI_PROJECT_DEFAULT_BRANCH])
	assert.Equal(t, "https://gitlab.com/cidverse/cienvsamples", normalized[ncispec.NCI_PROJECT_URL])
}

func TestNormalizer_Normalize_WorkflowAPI(t *testing.T) {

}
