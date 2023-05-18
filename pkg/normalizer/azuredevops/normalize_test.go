package azuredevops

import (
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
	assert.Equal(t, "1.0.0", normalized.Version)
	assert.Equal(t, normalizer.name, normalized.ServiceName)
	assert.Equal(t, normalizer.slug, normalized.ServiceSlug)
}

func TestNormalizer_Normalize_Worker(t *testing.T) {
	vcsrepository.MockClient = MockVCSClient(t)
	var normalizer = NewNormalizer()
	var normalized = normalizer.Normalize(map[string]string{
		"AGENT_ID":          "9",
		"AGENT_MACHINENAME": "fv-az158-714",
		"ImageOS":           "ubuntu20",
		"ImageVersion":      "20220503.1",
		"AGENT_VERSION":     "2.202.1",
	})

	assert.Equal(t, "9", normalized.Worker.Id)
	assert.Equal(t, "fv-az158-714", normalized.Worker.Name)
	assert.Equal(t, "azuredevops_hosted_vm", normalized.Worker.Type)
	assert.Equal(t, "ubuntu20:20220503.1", normalized.Worker.OS)
	assert.Equal(t, "2.202.1", normalized.Worker.Version)
	assert.Equal(t, runtime.GOOS+"/"+runtime.GOARCH, normalized.Worker.Arch)
}

func TestNormalizer_Normalize_Pipeline(t *testing.T) {
	vcsrepository.MockClient = MockVCSClient(t)
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

	assert.Equal(t, "a11efe29-9b58-5a6c-3fa4-3e36996dcbd8", normalized.Pipeline.Id)
	assert.Equal(t, "push", normalized.Pipeline.Trigger)
	assert.Equal(t, "6884a131-87da-5381-61f3-d7acc3b91d76", normalized.Pipeline.StageId)
	assert.Equal(t, "__run", normalized.Pipeline.StageName)
	assert.Equal(t, "run", normalized.Pipeline.StageSlug)
	assert.Equal(t, "3dc8fd7e-4368-5a92-293e-d53cefc8c4b3", normalized.Pipeline.JobId)
	assert.Equal(t, "__default", normalized.Pipeline.JobName)
	assert.Equal(t, "default", normalized.Pipeline.JobSlug)
	assert.NotNil(t, normalized.Pipeline.JobStartedAt)
	assert.Equal(t, "3", normalized.Pipeline.Attempt)
	assert.Equal(t, "https://heuer.visualstudio.com/cienvsamples/_build/results?buildId=11", normalized.Pipeline.Url)
}

func TestNormalizer_Normalize_Project(t *testing.T) {
}

func TestNormalizer_Normalize_PullRequest(t *testing.T) {
}

func TestNormalizer_Normalize_WorkflowAPI(t *testing.T) {
}
