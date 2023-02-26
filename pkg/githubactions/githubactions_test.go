package githubactions

import (
	_ "embed"
	"os"
	"runtime"
	"strings"
	"testing"

	"github.com/cidverse/normalizeci/pkg/ncispec"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"

	"github.com/cidverse/normalizeci/pkg/common"
)

//go:embed githubactions.env
var testEnvironmentFile string
var testEnvironment = strings.Split(testEnvironmentFile, "\n")

func TestMain(m *testing.M) {
	zerolog.SetGlobalLevel(zerolog.DebugLevel)
	code := m.Run()
	os.Exit(code)
}

func TestEnvironmentCheck(t *testing.T) {
	var normalizer = NewNormalizer()
	if normalizer.Check(common.GetEnvironmentFrom(testEnvironment)) != true {
		t.Errorf("Check should succeed with the provided sample data!")
	}
}

func TestEnvironmentNormalizer(t *testing.T) {
	var normalizer = NewNormalizer()
	var normalized = normalizer.Normalize(common.GetEnvironmentFrom(testEnvironment))

	// log all normalized values
	for key, element := range normalized {
		t.Log(key + "=" + element)
	}

	// validate fields
	// - common
	assert.Equal(t, "true", normalized[ncispec.NCI])
	assert.Equal(t, normalizer.version, normalized[ncispec.NCI_VERSION])
	assert.Equal(t, normalizer.name, normalized[ncispec.NCI_SERVICE_NAME])
	assert.Equal(t, normalizer.slug, normalized[ncispec.NCI_SERVICE_SLUG])
	// - worker
	assert.Equal(t, "github_969396af-1899-4849-9318-7807141c54e9", normalized[ncispec.NCI_WORKER_ID])
	assert.Equal(t, "github_969396af-1899-4849-9318-7807141c54e9", normalized[ncispec.NCI_WORKER_NAME])
	assert.Equal(t, "github_hosted_vm", normalized[ncispec.NCI_WORKER_TYPE])
	assert.Equal(t, "ubuntu20:20220503.1", normalized[ncispec.NCI_WORKER_OS])
	assert.Equal(t, "latest", normalized[ncispec.NCI_WORKER_VERSION])
	assert.Equal(t, runtime.GOOS+"/"+runtime.GOARCH, normalized[ncispec.NCI_WORKER_ARCH])
	// - pipeline
	assert.Equal(t, "2303126757", normalized[ncispec.NCI_PIPELINE_ID])
	assert.Equal(t, "push", normalized[ncispec.NCI_PIPELINE_TRIGGER])
	assert.Equal(t, "ci", normalized[ncispec.NCI_PIPELINE_STAGE_NAME])
	assert.Equal(t, "ci", normalized[ncispec.NCI_PIPELINE_STAGE_SLUG])
	assert.Equal(t, "__run", normalized[ncispec.NCI_PIPELINE_JOB_NAME])
	assert.Equal(t, "run", normalized[ncispec.NCI_PIPELINE_JOB_SLUG])
	assert.NotNil(t, normalized[ncispec.NCI_PIPELINE_JOB_STARTED_AT])
	assert.Equal(t, "1", normalized[ncispec.NCI_PIPELINE_ATTEMPT])
	assert.Equal(t, "https://github.com/cidverse/cienvsamples/actions/runs/2303126757", normalized[ncispec.NCI_PIPELINE_URL])
	// - project
	assert.Equal(t, "https://github.com/cidverse/cienvsamples", normalized[ncispec.NCI_PROJECT_URL])
}

func TestEnvironmentNormalizerPullRequestId(t *testing.T) {
	var normalizer = NewNormalizer()
	var normalized = normalizer.Normalize(common.GetEnvironmentMerge(testEnvironment, []string{"GITHUB_EVENT_NAME=pull_request", "GITHUB_REF=refs/pull/519/merge"}))

	assert.Equal(t, "519", normalized[ncispec.NCI_MERGE_REQUEST_ID])
}

func TestValidateSpec(t *testing.T) {
	var normalizer = NewNormalizer()
	var normalized = normalizer.Normalize(common.GetEnvironmentFrom(testEnvironment))

	nci := ncispec.OfMap(normalized)

	err := nci.Validate()
	assert.Emptyf(t, err, "there shouldn't be any validation errors")
}
