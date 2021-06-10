package projectdetails

import (
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	zerolog.SetGlobalLevel(zerolog.DebugLevel)
	code := m.Run()
	os.Exit(code)
}

func TestGetProjectDetailsGitHub(t *testing.T) {
	details := GetProjectDetails("git", "https://github.com/cidverse/normalizeci.git")

	assert.Equal(t, "205438004", details["NCI_PROJECT_ID"])
	assert.Equal(t, "normalizeci", details["NCI_PROJECT_NAME"])
	assert.Equal(t, "cidverse-normalizeci", details["NCI_PROJECT_SLUG"])
	assert.Equal(t, "A tool to turn the continious integration / deployment variables into a common format for generally usable scripts without any dependencies.", details["NCI_PROJECT_DESCRIPTION"])
	assert.Equal(t, "cicd,normalization", details["NCI_PROJECT_TOPICS"])
	assert.Equal(t, "https://api.github.com/repos/cidverse/normalizeci/issues/{ID}", details["NCI_PROJECT_ISSUE_URL"])
	assert.NotEmpty(t, details["NCI_PROJECT_STARGAZERS"])
	assert.NotEmpty(t, details["NCI_PROJECT_FORKS"])
}

func TestGetProjectDetailsGitLab(t *testing.T) {
	details := GetProjectDetails("git", "https://gitlab.com/gitlab-org/gitlab.git")

	assert.Equal(t, "278964", details["NCI_PROJECT_ID"])
	assert.Equal(t, "GitLab", details["NCI_PROJECT_NAME"])
	assert.Equal(t, "gitlab-org-gitlab", details["NCI_PROJECT_SLUG"])
	assert.Equal(t, "GitLab is an open source end-to-end software development platform with built-in version control, issue tracking, code review, CI/CD, and more. Self-host GitLab on your own servers, in a container, or on a cloud provider.", details["NCI_PROJECT_DESCRIPTION"])
	assert.Equal(t, "", details["NCI_PROJECT_TOPICS"])
	assert.Equal(t, "https://gitlab.com/gitlab-org/gitlab/-/issues/{ID}", details["NCI_PROJECT_ISSUE_URL"])
	assert.NotEmpty(t, details["NCI_PROJECT_STARGAZERS"])
	assert.NotEmpty(t, details["NCI_PROJECT_FORKS"])
}
