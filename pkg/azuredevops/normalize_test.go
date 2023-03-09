package azuredevops

import (
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
		"AGENT_ID":          "9",
		"AGENT_MACHINENAME": "fv-az158-714",
		"ImageOS":           "ubuntu20",
		"ImageVersion":      "20220503.1",
		"AGENT_VERSION":     "2.202.1",
	})

	assert.Equal(t, "9", normalized[ncispec.NCI_WORKER_ID])
	assert.Equal(t, "fv-az158-714", normalized[ncispec.NCI_WORKER_NAME])
	assert.Equal(t, "azuredevops_hosted_vm", normalized[ncispec.NCI_WORKER_TYPE])
	assert.Equal(t, "ubuntu20:20220503.1", normalized[ncispec.NCI_WORKER_OS])
	assert.Equal(t, "2.202.1", normalized[ncispec.NCI_WORKER_VERSION])
	assert.Equal(t, runtime.GOOS+"/"+runtime.GOARCH, normalized[ncispec.NCI_WORKER_ARCH])
}

func TestNormalizer_Normalize_Pipeline(t *testing.T) {
	var normalizer = NewNormalizer()
	var normalized = normalizer.Normalize(map[string]string{
		"SYSTEM_PHASEID":                 "a11efe29-9b58-5a6c-3fa4-3e36996dcbd8",
		"BUILD_REASON":                   "IndividualCI",
		"SYSTEM_STAGEID":                 "6884a131-87da-5381-61f3-d7acc3b91d76",
		"SYSTEM_STAGENAME":               "__run",
		"SYSTEM_JOBID":                   "3dc8fd7e-4368-5a92-293e-d53cefc8c4b3",
		"SYSTEM_JOBNAME":                 "__default",
		"SYSTEM_JOBATTEMPT":              "3",
		"SYSTEM_TEAMFOUNDATIONSERVERURI": "https://heuer.visualstudio.com/",
		"SYSTEM_TEAMPROJECT":             "cienvsamples",
		"BUILD_BUILDID":                  "11",
	})

	assert.Equal(t, "a11efe29-9b58-5a6c-3fa4-3e36996dcbd8", normalized[ncispec.NCI_PIPELINE_ID])
	assert.Equal(t, "push", normalized[ncispec.NCI_PIPELINE_TRIGGER])
	assert.Equal(t, "6884a131-87da-5381-61f3-d7acc3b91d76", normalized[ncispec.NCI_PIPELINE_STAGE_ID])
	assert.Equal(t, "__run", normalized[ncispec.NCI_PIPELINE_STAGE_NAME])
	assert.Equal(t, "run", normalized[ncispec.NCI_PIPELINE_STAGE_SLUG])
	assert.Equal(t, "3dc8fd7e-4368-5a92-293e-d53cefc8c4b3", normalized[ncispec.NCI_PIPELINE_JOB_ID])
	assert.Equal(t, "__default", normalized[ncispec.NCI_PIPELINE_JOB_NAME])
	assert.Equal(t, "default", normalized[ncispec.NCI_PIPELINE_JOB_SLUG])
	assert.NotNil(t, normalized[ncispec.NCI_PIPELINE_JOB_STARTED_AT])
	assert.Equal(t, "3", normalized[ncispec.NCI_PIPELINE_ATTEMPT])
	assert.Equal(t, "https://heuer.visualstudio.com/cienvsamples/_build/results?buildId=11", normalized[ncispec.NCI_PIPELINE_URL])
}

func TestNormalizer_Normalize_Project(t *testing.T) {
}

func TestNormalizer_Normalize_PullRequest(t *testing.T) {
}

func TestNormalizer_Normalize_WorkflowAPI(t *testing.T) {
}
