package localgit

import (
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"

	"github.com/qubid/normalizeci/pkg/common"
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
		t.Log(key+"="+element)
	}

	// validate fields
	// - common
	assert.Equal(t, "true", normalized["NCI"])
	assert.Equal(t, normalizer.version, normalized["NCI_VERSION"])
	assert.Equal(t, normalizer.name, normalized["NCI_SERVICE_NAME"])
	assert.Equal(t, normalizer.slug, normalized["NCI_SERVICE_SLUG"])
	// - server
	assert.Equal(t, "local", normalized["NCI_SERVER_NAME"])
	assert.Equal(t, "localhost", normalized["NCI_SERVER_HOST"])
	assert.Equal(t, "", normalized["NCI_SERVER_VERSION"])
	// - worker
	// - pipeline
	// - container registry
	assert.Equal(t, "", normalized["NCI_CONTAINERREGISTRY_HOST"])
	assert.Equal(t, "", normalized["NCI_CONTAINERREGISTRY_USERNAME"])
	assert.Equal(t, "", normalized["NCI_CONTAINERREGISTRY_PASSWORD"])
	// - project
}
