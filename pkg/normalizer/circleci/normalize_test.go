package circleci

import (
	"runtime"
	"testing"

	"github.com/cidverse/normalizeci/pkg/nciutil"
	"github.com/stretchr/testify/assert"
)

var circleciEnv = map[string]string{
	"CIRCLE_BRANCH":                "main",
	"CIRCLE_BUILD_NUM":             "9",
	"CIRCLE_BUILD_URL":             "https://app.circleci.com/jobs/circleci/0d12c8fa-cf92-415b-a14c-4787b20c52e2/aa4638b8-064e-4321-b1ee-794f3028b01f/9",
	"CIRCLE_INTERNAL_SCRATCH":      "/tmp/circleci-4169931879",
	"CIRCLE_INTERNAL_TASK_DATA":    "/tmp/.circleci-task-data-7564f453-074e-43cf-aef3-a393a9909474-0-build",
	"CIRCLE_JOB":                   "publish-env",
	"CIRCLE_NODE_INDEX":            "0",
	"CIRCLE_NODE_TOTAL":            "1",
	"CIRCLE_OIDC_TOKEN":            "secret",
	"CIRCLE_OIDC_TOKEN_V2":         "secret",
	"CIRCLE_ORGANIZATION_ID":       "0d12c8fa-cf92-415b-a14c-4787b20c52e2",
	"CIRCLE_PIPELINE_ID":           "ec202ec0-b88f-47cf-9df0-f85979ea3426",
	"CIRCLE_PROJECT_ID":            "aa4638b8-064e-4321-b1ee-794f3028b01f",
	"CIRCLE_PROJECT_REPONAME":      "cienvsamples",
	"CIRCLE_PROJECT_USERNAME":      "CIDVerse",
	"CIRCLE_REPOSITORY_URL":        "",
	"CIRCLE_SHA1":                  "790efd9b96e59d9b3c3f1899284c85fa91efbcbc",
	"CIRCLE_SHELL_ENV":             "/tmp/.bash_env-7564f453-074e-43cf-aef3-a393a9909474-0-build",
	"CIRCLE_USERNAME":              "dummy@example.com",
	"CIRCLE_WORKFLOW_ID":           "00adcc65-3be1-463b-b5f6-f2102329bb00",
	"CIRCLE_WORKFLOW_JOB_ID":       "7564f453-074e-43cf-aef3-a393a9909474",
	"CIRCLE_WORKFLOW_WORKSPACE_ID": "00adcc65-3be1-463b-b5f6-f2102329bb00",
	"CIRCLE_WORKING_DIRECTORY":     "/root/project",
}

func TestNormalizer_Normalize_Common(t *testing.T) {
	nciutil.MockVCSClient(t)

	var normalizer = NewNormalizer()
	var normalized, err = normalizer.Normalize(map[string]string{})

	assert.NoError(t, err)
	assert.Equal(t, "true", normalized.Found)
	assert.Equal(t, "1.0.0", normalized.Version)
	assert.Equal(t, normalizer.name, normalized.ServiceName)
	assert.Equal(t, normalizer.slug, normalized.ServiceSlug)
}

func TestNormalizer_Normalize_Worker(t *testing.T) {
	nciutil.MockVCSClient(t)

	var normalizer = NewNormalizer()
	var normalized, err = normalizer.Normalize(circleciEnv)

	assert.NoError(t, err)
	assert.Equal(t, "0", normalized.Worker.Id)
	assert.Equal(t, "0", normalized.Worker.Name)
	assert.Equal(t, "circleci_hosted_vm", normalized.Worker.Type)
	assert.Equal(t, "unknown", normalized.Worker.OS)
	assert.Equal(t, "latest", normalized.Worker.Version)
	assert.Equal(t, runtime.GOOS+"/"+runtime.GOARCH, normalized.Worker.Arch)
}

func TestNormalizer_Normalize_Pipeline(t *testing.T) {
	nciutil.MockVCSClient(t)

	var normalizer = NewNormalizer()
	var normalized, err = normalizer.Normalize(circleciEnv)

	assert.NoError(t, err)
	assert.Equal(t, "ec202ec0-b88f-47cf-9df0-f85979ea3426", normalized.Pipeline.Id)
	assert.Equal(t, "unknown", normalized.Pipeline.Trigger)
	assert.Equal(t, "00adcc65-3be1-463b-b5f6-f2102329bb00", normalized.Pipeline.StageId)
	assert.Equal(t, "default", normalized.Pipeline.StageName)
	assert.Equal(t, "default", normalized.Pipeline.StageSlug)
	assert.Equal(t, "7564f453-074e-43cf-aef3-a393a9909474", normalized.Pipeline.JobId)
	assert.Equal(t, "publish-env", normalized.Pipeline.JobName)
	assert.Equal(t, "publish-env", normalized.Pipeline.JobSlug)
	assert.NotNil(t, normalized.Pipeline.JobStartedAt)
	assert.Equal(t, "0", normalized.Pipeline.Attempt)
	assert.Equal(t, "https://app.circleci.com/jobs/circleci/0d12c8fa-cf92-415b-a14c-4787b20c52e2/aa4638b8-064e-4321-b1ee-794f3028b01f/9", normalized.Pipeline.Url)
}

func TestNormalizer_Normalize_Project(t *testing.T) {
}

func TestNormalizer_Normalize_PullRequest(t *testing.T) {
}

func TestNormalizer_Normalize_WorkflowAPI(t *testing.T) {
}
