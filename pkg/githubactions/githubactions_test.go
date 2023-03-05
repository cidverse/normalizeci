package githubactions

import (
	_ "embed"
	"net/http"
	"os"
	"runtime"
	"strings"
	"testing"

	"github.com/cidverse/normalizeci/pkg/ncispec"
	"github.com/jarcoal/httpmock"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"

	"github.com/cidverse/normalizeci/pkg/common"
)

//go:embed githubactions.env
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
	githubMockClient = &http.Client{}
	httpmock.ActivateNonDefault(githubMockClient)
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("GET", "https://api.github.com/repos/cidverse/cienvsamples/actions/runs/2303126757", httpmock.NewStringResponder(200, `{"id":2303126757,"name":"ci","node_id":"WFR_kwLOHTHf4s6JRuzl","head_branch":"main","head_sha":"1b37fdecbab29370c0715489429dbaed6581c678","path":".github/workflows/ci.yml","display_title":"feat: add azure-devops to update script","run_number":11,"event":"push","status":"completed","conclusion":"success","workflow_id":25656602,"check_suite_id":6453158213,"check_suite_node_id":"CS_kwDOHTHf4s8AAAABgKNhRQ","url":"https://api.github.com/repos/cidverse/cienvsamples/actions/runs/2303126757","html_url":"https://github.com/cidverse/cienvsamples/actions/runs/2303126757","pull_requests":[],"created_at":"2022-05-10T20:20:59Z","updated_at":"2022-05-10T20:21:20Z","actor":{"login":"PhilippHeuer","id":10275049,"node_id":"MDQ6VXNlcjEwMjc1MDQ5","avatar_url":"https://avatars.githubusercontent.com/u/10275049?v=4","gravatar_id":"","url":"https://api.github.com/users/PhilippHeuer","html_url":"https://github.com/PhilippHeuer","followers_url":"https://api.github.com/users/PhilippHeuer/followers","following_url":"https://api.github.com/users/PhilippHeuer/following{/other_user}","gists_url":"https://api.github.com/users/PhilippHeuer/gists{/gist_id}","starred_url":"https://api.github.com/users/PhilippHeuer/starred{/owner}{/repo}","subscriptions_url":"https://api.github.com/users/PhilippHeuer/subscriptions","organizations_url":"https://api.github.com/users/PhilippHeuer/orgs","repos_url":"https://api.github.com/users/PhilippHeuer/repos","events_url":"https://api.github.com/users/PhilippHeuer/events{/privacy}","received_events_url":"https://api.github.com/users/PhilippHeuer/received_events","type":"User","site_admin":false},"run_attempt":1,"referenced_workflows":[],"run_started_at":"2022-05-10T20:20:59Z","triggering_actor":{"login":"PhilippHeuer","id":10275049,"node_id":"MDQ6VXNlcjEwMjc1MDQ5","avatar_url":"https://avatars.githubusercontent.com/u/10275049?v=4","gravatar_id":"","url":"https://api.github.com/users/PhilippHeuer","html_url":"https://github.com/PhilippHeuer","followers_url":"https://api.github.com/users/PhilippHeuer/followers","following_url":"https://api.github.com/users/PhilippHeuer/following{/other_user}","gists_url":"https://api.github.com/users/PhilippHeuer/gists{/gist_id}","starred_url":"https://api.github.com/users/PhilippHeuer/starred{/owner}{/repo}","subscriptions_url":"https://api.github.com/users/PhilippHeuer/subscriptions","organizations_url":"https://api.github.com/users/PhilippHeuer/orgs","repos_url":"https://api.github.com/users/PhilippHeuer/repos","events_url":"https://api.github.com/users/PhilippHeuer/events{/privacy}","received_events_url":"https://api.github.com/users/PhilippHeuer/received_events","type":"User","site_admin":false},"jobs_url":"https://api.github.com/repos/cidverse/cienvsamples/actions/runs/2303126757/jobs","logs_url":"https://api.github.com/repos/cidverse/cienvsamples/actions/runs/2303126757/logs","check_suite_url":"https://api.github.com/repos/cidverse/cienvsamples/check-suites/6453158213","artifacts_url":"https://api.github.com/repos/cidverse/cienvsamples/actions/runs/2303126757/artifacts","cancel_url":"https://api.github.com/repos/cidverse/cienvsamples/actions/runs/2303126757/cancel","rerun_url":"https://api.github.com/repos/cidverse/cienvsamples/actions/runs/2303126757/rerun","previous_attempt_url":null,"workflow_url":"https://api.github.com/repos/cidverse/cienvsamples/actions/workflows/25656602","head_commit":{"id":"1b37fdecbab29370c0715489429dbaed6581c678","tree_id":"97c2e0439666b82d0b5d2a2875dd651a37d9c21f","message":"feat: add azure-devops to update script","timestamp":"2022-05-10T20:20:54Z","author":{"name":"Philipp Heuer","email":"git@philippheuer.me"},"committer":{"name":"Philipp Heuer","email":"git@philippheuer.me"}},"repository":{"id":489807842,"node_id":"R_kgDOHTHf4g","name":"cienvsamples","full_name":"cidverse/cienvsamples","private":false,"owner":{"login":"cidverse","id":84687161,"node_id":"MDEyOk9yZ2FuaXphdGlvbjg0Njg3MTYx","avatar_url":"https://avatars.githubusercontent.com/u/84687161?v=4","gravatar_id":"","url":"https://api.github.com/users/cidverse","html_url":"https://github.com/cidverse","followers_url":"https://api.github.com/users/cidverse/followers","following_url":"https://api.github.com/users/cidverse/following{/other_user}","gists_url":"https://api.github.com/users/cidverse/gists{/gist_id}","starred_url":"https://api.github.com/users/cidverse/starred{/owner}{/repo}","subscriptions_url":"https://api.github.com/users/cidverse/subscriptions","organizations_url":"https://api.github.com/users/cidverse/orgs","repos_url":"https://api.github.com/users/cidverse/repos","events_url":"https://api.github.com/users/cidverse/events{/privacy}","received_events_url":"https://api.github.com/users/cidverse/received_events","type":"Organization","site_admin":false},"html_url":"https://github.com/cidverse/cienvsamples","description":null,"fork":false,"url":"https://api.github.com/repos/cidverse/cienvsamples","forks_url":"https://api.github.com/repos/cidverse/cienvsamples/forks","keys_url":"https://api.github.com/repos/cidverse/cienvsamples/keys{/key_id}","collaborators_url":"https://api.github.com/repos/cidverse/cienvsamples/collaborators{/collaborator}","teams_url":"https://api.github.com/repos/cidverse/cienvsamples/teams","hooks_url":"https://api.github.com/repos/cidverse/cienvsamples/hooks","issue_events_url":"https://api.github.com/repos/cidverse/cienvsamples/issues/events{/number}","events_url":"https://api.github.com/repos/cidverse/cienvsamples/events","assignees_url":"https://api.github.com/repos/cidverse/cienvsamples/assignees{/user}","branches_url":"https://api.github.com/repos/cidverse/cienvsamples/branches{/branch}","tags_url":"https://api.github.com/repos/cidverse/cienvsamples/tags","blobs_url":"https://api.github.com/repos/cidverse/cienvsamples/git/blobs{/sha}","git_tags_url":"https://api.github.com/repos/cidverse/cienvsamples/git/tags{/sha}","git_refs_url":"https://api.github.com/repos/cidverse/cienvsamples/git/refs{/sha}","trees_url":"https://api.github.com/repos/cidverse/cienvsamples/git/trees{/sha}","statuses_url":"https://api.github.com/repos/cidverse/cienvsamples/statuses/{sha}","languages_url":"https://api.github.com/repos/cidverse/cienvsamples/languages","stargazers_url":"https://api.github.com/repos/cidverse/cienvsamples/stargazers","contributors_url":"https://api.github.com/repos/cidverse/cienvsamples/contributors","subscribers_url":"https://api.github.com/repos/cidverse/cienvsamples/subscribers","subscription_url":"https://api.github.com/repos/cidverse/cienvsamples/subscription","commits_url":"https://api.github.com/repos/cidverse/cienvsamples/commits{/sha}","git_commits_url":"https://api.github.com/repos/cidverse/cienvsamples/git/commits{/sha}","comments_url":"https://api.github.com/repos/cidverse/cienvsamples/comments{/number}","issue_comment_url":"https://api.github.com/repos/cidverse/cienvsamples/issues/comments{/number}","contents_url":"https://api.github.com/repos/cidverse/cienvsamples/contents/{+path}","compare_url":"https://api.github.com/repos/cidverse/cienvsamples/compare/{base}...{head}","merges_url":"https://api.github.com/repos/cidverse/cienvsamples/merges","archive_url":"https://api.github.com/repos/cidverse/cienvsamples/{archive_format}{/ref}","downloads_url":"https://api.github.com/repos/cidverse/cienvsamples/downloads","issues_url":"https://api.github.com/repos/cidverse/cienvsamples/issues{/number}","pulls_url":"https://api.github.com/repos/cidverse/cienvsamples/pulls{/number}","milestones_url":"https://api.github.com/repos/cidverse/cienvsamples/milestones{/number}","notifications_url":"https://api.github.com/repos/cidverse/cienvsamples/notifications{?since,all,participating}","labels_url":"https://api.github.com/repos/cidverse/cienvsamples/labels{/name}","releases_url":"https://api.github.com/repos/cidverse/cienvsamples/releases{/id}","deployments_url":"https://api.github.com/repos/cidverse/cienvsamples/deployments"},"head_repository":{"id":489807842,"node_id":"R_kgDOHTHf4g","name":"cienvsamples","full_name":"cidverse/cienvsamples","private":false,"owner":{"login":"cidverse","id":84687161,"node_id":"MDEyOk9yZ2FuaXphdGlvbjg0Njg3MTYx","avatar_url":"https://avatars.githubusercontent.com/u/84687161?v=4","gravatar_id":"","url":"https://api.github.com/users/cidverse","html_url":"https://github.com/cidverse","followers_url":"https://api.github.com/users/cidverse/followers","following_url":"https://api.github.com/users/cidverse/following{/other_user}","gists_url":"https://api.github.com/users/cidverse/gists{/gist_id}","starred_url":"https://api.github.com/users/cidverse/starred{/owner}{/repo}","subscriptions_url":"https://api.github.com/users/cidverse/subscriptions","organizations_url":"https://api.github.com/users/cidverse/orgs","repos_url":"https://api.github.com/users/cidverse/repos","events_url":"https://api.github.com/users/cidverse/events{/privacy}","received_events_url":"https://api.github.com/users/cidverse/received_events","type":"Organization","site_admin":false},"html_url":"https://github.com/cidverse/cienvsamples","description":null,"fork":false,"url":"https://api.github.com/repos/cidverse/cienvsamples","forks_url":"https://api.github.com/repos/cidverse/cienvsamples/forks","keys_url":"https://api.github.com/repos/cidverse/cienvsamples/keys{/key_id}","collaborators_url":"https://api.github.com/repos/cidverse/cienvsamples/collaborators{/collaborator}","teams_url":"https://api.github.com/repos/cidverse/cienvsamples/teams","hooks_url":"https://api.github.com/repos/cidverse/cienvsamples/hooks","issue_events_url":"https://api.github.com/repos/cidverse/cienvsamples/issues/events{/number}","events_url":"https://api.github.com/repos/cidverse/cienvsamples/events","assignees_url":"https://api.github.com/repos/cidverse/cienvsamples/assignees{/user}","branches_url":"https://api.github.com/repos/cidverse/cienvsamples/branches{/branch}","tags_url":"https://api.github.com/repos/cidverse/cienvsamples/tags","blobs_url":"https://api.github.com/repos/cidverse/cienvsamples/git/blobs{/sha}","git_tags_url":"https://api.github.com/repos/cidverse/cienvsamples/git/tags{/sha}","git_refs_url":"https://api.github.com/repos/cidverse/cienvsamples/git/refs{/sha}","trees_url":"https://api.github.com/repos/cidverse/cienvsamples/git/trees{/sha}","statuses_url":"https://api.github.com/repos/cidverse/cienvsamples/statuses/{sha}","languages_url":"https://api.github.com/repos/cidverse/cienvsamples/languages","stargazers_url":"https://api.github.com/repos/cidverse/cienvsamples/stargazers","contributors_url":"https://api.github.com/repos/cidverse/cienvsamples/contributors","subscribers_url":"https://api.github.com/repos/cidverse/cienvsamples/subscribers","subscription_url":"https://api.github.com/repos/cidverse/cienvsamples/subscription","commits_url":"https://api.github.com/repos/cidverse/cienvsamples/commits{/sha}","git_commits_url":"https://api.github.com/repos/cidverse/cienvsamples/git/commits{/sha}","comments_url":"https://api.github.com/repos/cidverse/cienvsamples/comments{/number}","issue_comment_url":"https://api.github.com/repos/cidverse/cienvsamples/issues/comments{/number}","contents_url":"https://api.github.com/repos/cidverse/cienvsamples/contents/{+path}","compare_url":"https://api.github.com/repos/cidverse/cienvsamples/compare/{base}...{head}","merges_url":"https://api.github.com/repos/cidverse/cienvsamples/merges","archive_url":"https://api.github.com/repos/cidverse/cienvsamples/{archive_format}{/ref}","downloads_url":"https://api.github.com/repos/cidverse/cienvsamples/downloads","issues_url":"https://api.github.com/repos/cidverse/cienvsamples/issues{/number}","pulls_url":"https://api.github.com/repos/cidverse/cienvsamples/pulls{/number}","milestones_url":"https://api.github.com/repos/cidverse/cienvsamples/milestones{/number}","notifications_url":"https://api.github.com/repos/cidverse/cienvsamples/notifications{?since,all,participating}","labels_url":"https://api.github.com/repos/cidverse/cienvsamples/labels{/name}","releases_url":"https://api.github.com/repos/cidverse/cienvsamples/releases{/id}","deployments_url":"https://api.github.com/repos/cidverse/cienvsamples/deployments"}}`))
	httpmock.RegisterResponder("GET", "https://api.github.com/repos/cidverse/cienvsamples/actions/workflows/25656602", httpmock.NewStringResponder(200, `{"id":25656602,"node_id":"W_kwDOHTHf4s4Bh30a","name":"ci","path":".github/workflows/ci.yml","state":"active","created_at":"2022-05-08T01:55:02.000Z","updated_at":"2022-05-08T01:55:02.000Z","url":"https://api.github.com/repos/cidverse/cienvsamples/actions/workflows/25656602","html_url":"https://github.com/cidverse/cienvsamples/blob/main/.github/workflows/ci.yml","badge_url":"https://github.com/cidverse/cienvsamples/workflows/ci/badge.svg"}`))

	var normalizer = NewNormalizer()
	var normalized = normalizer.Normalize(common.GetEnvironmentFrom(testEnvironment))

	// log all normalized values
	for key, element := range normalized {
		t.Log(key + "=" + element)
	}

	// validate fields
	// - common
	assert.Equal(t, "true", normalized[ncispec.NCI])
	assert.Equal(t, normalizer.version, normalized[ncispec.NCI_VERSION])
	assert.Equal(t, normalizer.name, normalized[ncispec.NCI_SERVICE_NAME])
	assert.Equal(t, normalizer.slug, normalized[ncispec.NCI_SERVICE_SLUG])
	// - worker
	assert.Equal(t, "github_969396af-1899-4849-9318-7807141c54e9", normalized[ncispec.NCI_WORKER_ID])
	assert.Equal(t, "github_969396af-1899-4849-9318-7807141c54e9", normalized[ncispec.NCI_WORKER_NAME])
	assert.Equal(t, "github_hosted_vm", normalized[ncispec.NCI_WORKER_TYPE])
	assert.Equal(t, "ubuntu20:20220503.1", normalized[ncispec.NCI_WORKER_OS])
	assert.Equal(t, "latest", normalized[ncispec.NCI_WORKER_VERSION])
	assert.Equal(t, runtime.GOOS+"/"+runtime.GOARCH, normalized[ncispec.NCI_WORKER_ARCH])
	// - pipeline
	assert.Equal(t, "2303126757", normalized[ncispec.NCI_PIPELINE_ID])
	assert.Equal(t, "push", normalized[ncispec.NCI_PIPELINE_TRIGGER])
	assert.Equal(t, "ci", normalized[ncispec.NCI_PIPELINE_STAGE_NAME])
	assert.Equal(t, "ci", normalized[ncispec.NCI_PIPELINE_STAGE_SLUG])
	assert.Equal(t, "__run", normalized[ncispec.NCI_PIPELINE_JOB_NAME])
	assert.Equal(t, "run", normalized[ncispec.NCI_PIPELINE_JOB_SLUG])
	assert.NotNil(t, normalized[ncispec.NCI_PIPELINE_JOB_STARTED_AT])
	assert.Equal(t, "1", normalized[ncispec.NCI_PIPELINE_ATTEMPT])
	assert.Equal(t, "https://github.com/cidverse/cienvsamples/actions/runs/2303126757", normalized[ncispec.NCI_PIPELINE_URL])
	// - project
	assert.Equal(t, "https://github.com/cidverse/cienvsamples", normalized[ncispec.NCI_PROJECT_URL])
}

func TestEnvironmentNormalizerPullRequestId(t *testing.T) {
	var normalizer = NewNormalizer()
	var normalized = normalizer.Normalize(common.GetEnvironmentMerge(testEnvironment, []string{"GITHUB_EVENT_NAME=pull_request", "GITHUB_REF=refs/pull/519/merge"}))

	assert.Equal(t, "519", normalized[ncispec.NCI_MERGE_REQUEST_ID])
}

func TestValidateSpec(t *testing.T) {
	var normalizer = NewNormalizer()
	var normalized = normalizer.Normalize(common.GetEnvironmentFrom(testEnvironment))

	nci := ncispec.OfMap(normalized)

	err := nci.Validate()
	assert.Emptyf(t, err, "there shouldn't be any validation errors")
}
