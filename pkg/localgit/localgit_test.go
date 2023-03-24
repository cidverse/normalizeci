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

func TestNormalizer_Normalize_Common(t *testing.T) {
	var normalizer = NewNormalizer()
	var normalized = normalizer.Normalize(map[string]string{})

	assert.Equal(t, "true", normalized.Found)
	assert.Equal(t, normalizer.version, normalized.Version)
	assert.Equal(t, normalizer.name, normalized.ServiceName)
	assert.Equal(t, normalizer.slug, normalized.ServiceSlug)
}

func TestNormalizer_Normalize_Worker(t *testing.T) {
	var normalizer = NewNormalizer()
	var normalized = normalizer.Normalize(map[string]string{})

	assert.Equal(t, "local", normalized.WorkerId)
	assert.Equal(t, "localhost", normalized.WorkerName)
	assert.Equal(t, "local", normalized.WorkerType)
	assert.NotNil(t, normalized.WorkerOS)
	assert.Equal(t, "1.0.0", normalized.WorkerVersion)
	assert.NotNil(t, normalized.WorkerArch)
}

func TestNormalizer_Normalize_Pipeline(t *testing.T) {
	var normalizer = NewNormalizer()
	var normalized = normalizer.Normalize(map[string]string{})

	assert.NotNil(t, normalized.PipelineId)
	assert.Equal(t, ncispec.PipelineTriggerCLI, normalized.PipelineTrigger)
	assert.NotNil(t, normalized.PipelineStageId)
	assert.Equal(t, ncispec.PipelineStageDefault, normalized.PipelineStageName)
	assert.Equal(t, ncispec.PipelineStageDefault, normalized.PipelineStageSlug)
	assert.NotNil(t, normalized.PipelineJobId)
	assert.Equal(t, ncispec.PipelineJobDefault, normalized.PipelineJobName)
	assert.Equal(t, ncispec.PipelineJobDefault, normalized.PipelineJobSlug)
	assert.NotNil(t, normalized.PipelineJobStartedAt)
	assert.Equal(t, "1", normalized.PipelineAttempt)
}
