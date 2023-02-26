package localgit

import (
	"os"
	"testing"

	"github.com/cidverse/normalizeci/pkg/ncispec"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"

	"github.com/cidverse/normalizeci/pkg/common"
)

var testEnvironment []string

func TestMain(m *testing.M) {
	zerolog.SetGlobalLevel(zerolog.DebugLevel)
	code := m.Run()
	os.Exit(code)
}

func TestEnvironmentCheck(t *testing.T) {
	var normalizer = NewNormalizer()
	if normalizer.Check(common.GetEnvironmentFrom(testEnvironment)) != true {
		t.Errorf("Check should succeed, since this project is a git repository")
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
	assert.Equal(t, "local", normalized[ncispec.NCI_WORKER_ID])
	assert.Equal(t, "localhost", normalized[ncispec.NCI_WORKER_NAME])
	assert.Equal(t, "local", normalized[ncispec.NCI_WORKER_TYPE])
	assert.NotNil(t, normalized[ncispec.NCI_WORKER_OS])
	assert.Equal(t, "1.0.0", normalized[ncispec.NCI_WORKER_VERSION])
	assert.NotNil(t, normalized[ncispec.NCI_WORKER_ARCH])
	// - pipeline
	assert.NotNil(t, normalized[ncispec.NCI_PIPELINE_ID])
	assert.Equal(t, ncispec.PipelineTriggerCLI, normalized[ncispec.NCI_PIPELINE_TRIGGER])
	assert.NotNil(t, normalized[ncispec.NCI_PIPELINE_STAGE_ID])
	assert.Equal(t, ncispec.PipelineStageDefault, normalized[ncispec.NCI_PIPELINE_STAGE_NAME])
	assert.Equal(t, ncispec.PipelineStageDefault, normalized[ncispec.NCI_PIPELINE_STAGE_SLUG])
	assert.NotNil(t, normalized[ncispec.NCI_PIPELINE_JOB_ID])
	assert.Equal(t, ncispec.PipelineJobDefault, normalized[ncispec.NCI_PIPELINE_JOB_NAME])
	assert.Equal(t, ncispec.PipelineJobDefault, normalized[ncispec.NCI_PIPELINE_JOB_SLUG])
	assert.NotNil(t, normalized[ncispec.NCI_PIPELINE_JOB_STARTED_AT])
	assert.Equal(t, "1", normalized["NCI_PIPELINE_ATTEMPT"])
	// - container registry
	assert.Equal(t, "", normalized[ncispec.NCI_CONTAINERREGISTRY_HOST])
	assert.Equal(t, "", normalized[ncispec.NCI_CONTAINERREGISTRY_USERNAME])
	assert.Equal(t, "", normalized[ncispec.NCI_CONTAINERREGISTRY_PASSWORD])
	// - project
}

func TestValidateSpec(t *testing.T) {
	var normalizer = NewNormalizer()
	var normalized = normalizer.Normalize(common.GetEnvironmentFrom(testEnvironment))

	nci := ncispec.OfMap(normalized)

	err := nci.Validate()
	assert.Emptyf(t, err, "there shouldn't be any validation errors")
}
