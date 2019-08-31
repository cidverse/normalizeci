package localgit

import (
	"testing"
)

var testEnvironment = []string{}

func TestEnvironmentCheck(t *testing.T) {
	var normalizer = NewNormalizer()
	if normalizer.Check(testEnvironment) != true {
		t.Errorf("Check should succeed, since this project is a git repository")
	}
}
