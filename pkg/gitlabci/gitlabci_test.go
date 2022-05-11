package gitlabci

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

//go:embed gitlabci.env
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
		t.Errorf("Check should succeed with the provided gitlab ci sample data!")
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
	assert.Equal(t, "12270837", normalized["NCI_WORKER_ID"])
	assert.Equal(t, "4-blue.shared.runners-manager.gitlab.com/default", normalized["NCI_WORKER_NAME"])
	assert.Equal(t, "14.10.0~beta.50.g1f2fe53e", normalized["NCI_WORKER_VERSION"])
	assert.Equal(t, runtime.GOOS+"/"+runtime.GOARCH, normalized["NCI_WORKER_ARCH"])
	// - pipeline
	assert.Equal(t, "push", normalized["NCI_PIPELINE_TRIGGER"])
	assert.Equal(t, "build", normalized["NCI_PIPELINE_STAGE_NAME"])
	assert.Equal(t, "build", normalized["NCI_PIPELINE_STAGE_SLUG"])
	assert.Equal(t, "build", normalized["NCI_PIPELINE_JOB_NAME"])
	assert.Equal(t, "build", normalized["NCI_PIPELINE_JOB_SLUG"])
	// - container registry
	assert.Equal(t, "registry.gitlab.com", normalized["NCI_CONTAINERREGISTRY_HOST"])
	assert.Equal(t, "registry.gitlab.com/cidverse/cienvsamples", normalized["NCI_CONTAINERREGISTRY_REPOSITORY"])
	assert.Equal(t, "gitlab-ci-token", normalized["NCI_CONTAINERREGISTRY_USERNAME"])
	assert.Equal(t, "secret", normalized["NCI_CONTAINERREGISTRY_PASSWORD"])
}

func TestValidateSpec(t *testing.T) {
	var normalizer = NewNormalizer()
	var normalized = normalizer.Normalize(common.GetEnvironmentFrom(testEnvironment))

	nci := ncispec.OfMap(normalized)

	err := nci.Validate()
	assert.Emptyf(t, err, "there shouldn't be any validation errors")
}
