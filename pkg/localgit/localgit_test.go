package localgit

import (
	"testing"

	"github.com/PhilippHeuer/normalize-ci/pkg/common"
)

var testEnvironment = []string{
	"NCI_CONTAINERREGISTRY_USERNAME=ci-token",
	"NCI_CONTAINERREGISTRY_PASSWORD=secret",
	"NCI_CONTAINERREGISTRY_HOST=registry.gitlab.com",
}

func TestEnvironmentCheck(t *testing.T) {
	var normalizer = NewNormalizer()
	if normalizer.Check(testEnvironment) != true {
		t.Errorf("Check should succeed, since this project is a git repository")
	}
}

func TestEnvironmentNormalizer(t *testing.T) {
	var normalizer = NewNormalizer()
	var normalized = normalizer.Normalize(testEnvironment)

	// log all normalized values
	for _, envvar := range normalized {
		t.Log(envvar)
	}

	// validate fields
	// - common
	common.AssertThatEnvEquals(t, normalized, "NCI", "true")
	common.AssertThatEnvEquals(t, normalized, "NCI_VERSION", normalizer.version)
	common.AssertThatEnvEquals(t, normalized, "NCI_SERVICE_NAME", normalizer.name)
	common.AssertThatEnvEquals(t, normalized, "NCI_SERVICE_SLUG", normalizer.slug)
	// - server
	common.AssertThatEnvEquals(t, normalized, "NCI_SERVER_NAME", "local")
	common.AssertThatEnvEquals(t, normalized, "NCI_SERVER_HOST", "localhost")
	common.AssertThatEnvEquals(t, normalized, "NCI_SERVER_VERSION", "")
	// - worker
	// - pipeline
	// - container registry
	common.AssertThatEnvEquals(t, normalized, "NCI_CONTAINERREGISTRY_HOST", "registry.gitlab.com")
	common.AssertThatEnvEquals(t, normalized, "NCI_CONTAINERREGISTRY_USERNAME", "ci-token")
	common.AssertThatEnvEquals(t, normalized, "NCI_CONTAINERREGISTRY_PASSWORD", "secret")
	// - project
}
