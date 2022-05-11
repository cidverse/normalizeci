package azuredevops

import (
	_ "embed"
	"github.com/cidverse/normalizeci/pkg/ncispec"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
	"os"
	"runtime"
	"strings"
	"testing"

	"github.com/cidverse/normalizeci/pkg/common"
)

//go:embed azuredevops.env
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
	assert.Equal(t, "true", normalized["NCI"])
	assert.Equal(t, normalizer.version, normalized["NCI_VERSION"])
	assert.Equal(t, normalizer.name, normalized["NCI_SERVICE_NAME"])
	assert.Equal(t, normalizer.slug, normalized["NCI_SERVICE_SLUG"])
	// - worker
	assert.Equal(t, "9", normalized["NCI_WORKER_ID"])
	assert.Equal(t, "fv-az158-714", normalized["NCI_WORKER_NAME"])
	assert.Equal(t, "2.202.1", normalized["NCI_WORKER_VERSION"])
	assert.Equal(t, runtime.GOOS+"/"+runtime.GOARCH, normalized["NCI_WORKER_ARCH"])
	// - pipeline
	assert.Equal(t, "push", normalized["NCI_PIPELINE_TRIGGER"])
	assert.Equal(t, "Build", normalized["NCI_PIPELINE_STAGE_NAME"])
	assert.Equal(t, "build", normalized["NCI_PIPELINE_STAGE_SLUG"])
	assert.Equal(t, "__default", normalized["NCI_PIPELINE_JOB_NAME"])
	assert.Equal(t, "default", normalized["NCI_PIPELINE_JOB_SLUG"])
}

func TestValidateSpec(t *testing.T) {
	var normalizer = NewNormalizer()
	var normalized = normalizer.Normalize(common.GetEnvironmentFrom(testEnvironment))

	nci := ncispec.OfMap(normalized)

	err := nci.Validate()
	assert.Emptyf(t, err, "there shouldn't be any validation errors")
}
