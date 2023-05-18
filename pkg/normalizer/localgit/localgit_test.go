package localgit

import (
	"os"
	"testing"

	"github.com/cidverse/normalizeci/pkg/ncispec/common"
	"github.com/cidverse/normalizeci/pkg/normalizer/api"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
)

var testEnvironment []string

func TestMain(m *testing.M) {
	zerolog.SetGlobalLevel(zerolog.DebugLevel)
	code := m.Run()
	os.Exit(code)
}

func TestEnvironmentCheck(t *testing.T) {
	var normalizer = NewNormalizer()
	if normalizer.Check(api.GetEnvironmentFrom(testEnvironment)) != true {
		t.Errorf("Check should succeed, since this project is a git repository")
	}
}

func TestNormalizer_Normalize_Common(t *testing.T) {
	var normalizer = NewNormalizer()
	var normalized = normalizer.Normalize(map[string]string{})

	assert.Equal(t, "true", normalized.Found)
	assert.Equal(t, "1.0.0", normalized.Version)
	assert.Equal(t, normalizer.name, normalized.ServiceName)
	assert.Equal(t, normalizer.slug, normalized.ServiceSlug)
}

func TestNormalizer_Normalize_Worker(t *testing.T) {
	var normalizer = NewNormalizer()
	var normalized = normalizer.Normalize(map[string]string{})

	assert.Equal(t, "local", normalized.Worker.Id)
	assert.Equal(t, "localhost", normalized.Worker.Name)
	assert.Equal(t, "local", normalized.Worker.Type)
	assert.NotNil(t, normalized.Worker.OS)
	assert.Equal(t, "1.0.0", normalized.Worker.Version)
	assert.NotNil(t, normalized.Worker.Arch)
}

func TestNormalizer_Normalize_Pipeline(t *testing.T) {
	var normalizer = NewNormalizer()
	var normalized = normalizer.Normalize(map[string]string{})

	assert.NotNil(t, normalized.Pipeline.Id)
	assert.Equal(t, common.PipelineTriggerCLI, normalized.Pipeline.Trigger)
	assert.NotNil(t, normalized.Pipeline.StageId)
	assert.Equal(t, common.PipelineStageDefault, normalized.Pipeline.StageName)
	assert.Equal(t, common.PipelineStageDefault, normalized.Pipeline.StageSlug)
	assert.NotNil(t, normalized.Pipeline.JobId)
	assert.Equal(t, common.PipelineJobDefault, normalized.Pipeline.JobName)
	assert.Equal(t, common.PipelineJobDefault, normalized.Pipeline.JobSlug)
	assert.NotNil(t, normalized.Pipeline.JobStartedAt)
	assert.Equal(t, "1", normalized.Pipeline.Attempt)
}
