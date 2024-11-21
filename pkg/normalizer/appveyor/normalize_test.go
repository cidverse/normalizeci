package appveyor

import (
	"runtime"
	"testing"

	"github.com/cidverse/normalizeci/pkg/nciutil"
	"github.com/stretchr/testify/assert"
)

var appveyorEnv = map[string]string{
	"APPVEYOR":                               "true",
	"APPVEYOR_ACCOUNT_NAME":                  "PhilippHeuer",
	"APPVEYOR_API_URL":                       "http://localhost:36401/",
	"APPVEYOR_BUILD_FOLDER":                  "/home/appveyor/projects/cienvsamples",
	"APPVEYOR_BUILD_ID":                      "51035157",
	"APPVEYOR_BUILD_NUMBER":                  "19",
	"APPVEYOR_BUILD_VERSION":                 "1.0.19",
	"APPVEYOR_BUILD_WORKER_IMAGE":            "Ubuntu2004",
	"APPVEYOR_CONSOLE_DISABLE_PTY":           "false",
	"APPVEYOR_JOB_ID":                        "yyrrqc54e997xewx",
	"APPVEYOR_JOB_NAME":                      "",
	"APPVEYOR_JOB_NUMBER":                    "1",
	"APPVEYOR_PROJECT_ID":                    "928594",
	"APPVEYOR_PROJECT_NAME":                  "cienvsamples",
	"APPVEYOR_PROJECT_SLUG":                  "cienvsamples",
	"APPVEYOR_PULL_REQUEST_HEAD_COMMIT":      "",
	"APPVEYOR_PULL_REQUEST_HEAD_REPO_BRANCH": "",
	"APPVEYOR_PULL_REQUEST_HEAD_REPO_NAME":   "",
	"APPVEYOR_PULL_REQUEST_NUMBER":           "",
	"APPVEYOR_PULL_REQUEST_TITLE":            "",
	"APPVEYOR_REPO_BRANCH":                   "main",
	"APPVEYOR_REPO_COMMIT":                   "790efd9b96e59d9b3c3f1899284c85fa91efbcbc",
	"APPVEYOR_REPO_COMMIT_AUTHOR":            "Philipp Heuer",
	"APPVEYOR_REPO_COMMIT_AUTHOR_EMAIL":      "dummy@example.com",
	"APPVEYOR_REPO_COMMIT_MESSAGE":           "feat: add gitlab sync job",
	"APPVEYOR_REPO_COMMIT_MESSAGE_EXTENDED":  "",
	"APPVEYOR_REPO_COMMIT_TIMESTAMP":         "2024-11-21T20:07:21.0000000Z",
	"APPVEYOR_REPO_NAME":                     "cidverse/cienvsamples",
	"APPVEYOR_REPO_PROVIDER":                 "gitHub",
	"APPVEYOR_REPO_SCM":                      "git",
	"APPVEYOR_REPO_TAG":                      "false",
	"APPVEYOR_REPO_TAG_NAME":                 "",
	"APPVEYOR_URL":                           "https://ci.appveyor.com",
	"ASPNETCORE_ENVIRONMENT":                 "Production",
	"CI":                                     "true",
	"CI_SERVICE_NAME":                        "appveyor",
}

func TestNormalizer_Normalize_Common(t *testing.T) {
	nciutil.MockVCSClient(t)

	var normalizer = NewNormalizer()
	var normalized, err = normalizer.Normalize(appveyorEnv)

	assert.NoError(t, err)
	assert.Equal(t, "true", normalized.Found)
	assert.Equal(t, "1.0.0", normalized.Version)
	assert.Equal(t, normalizer.name, normalized.ServiceName)
	assert.Equal(t, normalizer.slug, normalized.ServiceSlug)
}

func TestNormalizer_Normalize_Worker(t *testing.T) {
	nciutil.MockVCSClient(t)

	var normalizer = NewNormalizer()
	var normalized, err = normalizer.Normalize(appveyorEnv)

	assert.NoError(t, err)
	assert.Equal(t, "0", normalized.Worker.Id)
	assert.Equal(t, "unknown", normalized.Worker.Name)
	assert.Equal(t, "appveyor_hosted_vm", normalized.Worker.Type)
	assert.Equal(t, "Ubuntu2004", normalized.Worker.OS)
	assert.Equal(t, "latest", normalized.Worker.Version)
	assert.Equal(t, runtime.GOOS+"/"+runtime.GOARCH, normalized.Worker.Arch)
}

func TestNormalizer_Normalize_Pipeline(t *testing.T) {
	nciutil.MockVCSClient(t)

	var normalizer = NewNormalizer()
	var normalized, err = normalizer.Normalize(appveyorEnv)

	assert.NoError(t, err)
	assert.Equal(t, "51035157", normalized.Pipeline.Id)
	assert.Equal(t, "push", normalized.Pipeline.Trigger)
	assert.Equal(t, "51035157", normalized.Pipeline.StageId)
	assert.Equal(t, "default", normalized.Pipeline.StageName)
	assert.Equal(t, "default", normalized.Pipeline.StageSlug)
	assert.Equal(t, "yyrrqc54e997xewx", normalized.Pipeline.JobId)
	assert.Equal(t, "", normalized.Pipeline.JobName)
	assert.Equal(t, "", normalized.Pipeline.JobSlug)
	assert.NotNil(t, normalized.Pipeline.JobStartedAt)
	assert.Equal(t, "1", normalized.Pipeline.Attempt)
	assert.Equal(t, "https://ci.appveyor.com/project/PhilippHeuer/cienvsamples/builds/51035157", normalized.Pipeline.Url)
}

func TestNormalizer_Normalize_Project(t *testing.T) {
}

func TestNormalizer_Normalize_PullRequest(t *testing.T) {
}

func TestNormalizer_Normalize_WorkflowAPI(t *testing.T) {
}
