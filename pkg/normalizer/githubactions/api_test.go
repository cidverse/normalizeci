package githubactions

import (
	_ "embed"
	"os"
	"testing"

	"github.com/google/go-github/v61/github"
	"github.com/stretchr/testify/assert"
)

const workflowTriggerPayload = `{"inputs":{"example-value": "1","logLevel": "info"},"organization":{"avatar_url":"https://avatars.githubusercontent.com/u/84687161?v=4","description":"A home for projects that aim to improve and unify CI/CD related processes in a platform agnostic way.","events_url":"https://api.github.com/orgs/cidverse/events","hooks_url":"https://api.github.com/orgs/cidverse/hooks","id":84687161,"issues_url":"https://api.github.com/orgs/cidverse/issues","login":"cidverse","members_url":"https://api.github.com/orgs/cidverse/members{/member}","node_id":"MDEyOk9yZ2FuaXphdGlvbjg0Njg3MTYx","public_members_url":"https://api.github.com/orgs/cidverse/public_members{/member}","repos_url":"https://api.github.com/orgs/cidverse/repos","url":"https://api.github.com/orgs/cidverse"},"ref":"refs/heads/main","repository":{"allow_forking":true,"archive_url":"https://api.github.com/repos/cidverse/test/{archive_format}{/ref}","archived":false,"assignees_url":"https://api.github.com/repos/cidverse/test/assignees{/user}","blobs_url":"https://api.github.com/repos/cidverse/test/git/blobs{/sha}","branches_url":"https://api.github.com/repos/cidverse/test/branches{/branch}","clone_url":"https://github.com/cidverse/test.git","collaborators_url":"https://api.github.com/repos/cidverse/test/collaborators{/collaborator}","comments_url":"https://api.github.com/repos/cidverse/test/comments{/number}","commits_url":"https://api.github.com/repos/cidverse/test/commits{/sha}","compare_url":"https://api.github.com/repos/cidverse/test/compare/{base}...{head}","contents_url":"https://api.github.com/repos/cidverse/test/contents/{+path}","contributors_url":"https://api.github.com/repos/cidverse/test/contributors","created_at":"2023-02-06T20:19:13Z","default_branch":"main","deployments_url":"https://api.github.com/repos/cidverse/test/deployments","description":null,"disabled":false,"downloads_url":"https://api.github.com/repos/cidverse/test/downloads","events_url":"https://api.github.com/repos/cidverse/test/events","fork":false,"forks":0,"forks_count":0,"forks_url":"https://api.github.com/repos/cidverse/test/forks","full_name":"cidverse/test","git_commits_url":"https://api.github.com/repos/cidverse/test/git/commits{/sha}","git_refs_url":"https://api.github.com/repos/cidverse/test/git/refs{/sha}","git_tags_url":"https://api.github.com/repos/cidverse/test/git/tags{/sha}","git_url":"git://github.com/cidverse/test.git","has_discussions":false,"has_downloads":true,"has_issues":true,"has_pages":false,"has_projects":true,"has_wiki":true,"homepage":null,"hooks_url":"https://api.github.com/repos/cidverse/test/hooks","html_url":"https://github.com/cidverse/test","id":598298087,"is_template":false,"issue_comment_url":"https://api.github.com/repos/cidverse/test/issues/comments{/number}","issue_events_url":"https://api.github.com/repos/cidverse/test/issues/events{/number}","issues_url":"https://api.github.com/repos/cidverse/test/issues{/number}","keys_url":"https://api.github.com/repos/cidverse/test/keys{/key_id}","labels_url":"https://api.github.com/repos/cidverse/test/labels{/name}","language":"Go","languages_url":"https://api.github.com/repos/cidverse/test/languages","license":null,"merges_url":"https://api.github.com/repos/cidverse/test/merges","milestones_url":"https://api.github.com/repos/cidverse/test/milestones{/number}","mirror_url":null,"name":"test","node_id":"R_kgDOI6lN5w","notifications_url":"https://api.github.com/repos/cidverse/test/notifications{?since,all,participating}","open_issues":1,"open_issues_count":1,"owner":{"avatar_url":"https://avatars.githubusercontent.com/u/84687161?v=4","events_url":"https://api.github.com/users/cidverse/events{/privacy}","followers_url":"https://api.github.com/users/cidverse/followers","following_url":"https://api.github.com/users/cidverse/following{/other_user}","gists_url":"https://api.github.com/users/cidverse/gists{/gist_id}","gravatar_id":"","html_url":"https://github.com/cidverse","id":84687161,"login":"cidverse","node_id":"MDEyOk9yZ2FuaXphdGlvbjg0Njg3MTYx","organizations_url":"https://api.github.com/users/cidverse/orgs","received_events_url":"https://api.github.com/users/cidverse/received_events","repos_url":"https://api.github.com/users/cidverse/repos","site_admin":false,"starred_url":"https://api.github.com/users/cidverse/starred{/owner}{/repo}","subscriptions_url":"https://api.github.com/users/cidverse/subscriptions","type":"Organization","url":"https://api.github.com/users/cidverse"},"private":false,"pulls_url":"https://api.github.com/repos/cidverse/test/pulls{/number}","pushed_at":"2023-02-12T22:15:37Z","releases_url":"https://api.github.com/repos/cidverse/test/releases{/id}","size":31,"ssh_url":"git@github.com:cidverse/test.git","stargazers_count":0,"stargazers_url":"https://api.github.com/repos/cidverse/test/stargazers","statuses_url":"https://api.github.com/repos/cidverse/test/statuses/{sha}","subscribers_url":"https://api.github.com/repos/cidverse/test/subscribers","subscription_url":"https://api.github.com/repos/cidverse/test/subscription","svn_url":"https://github.com/cidverse/test","tags_url":"https://api.github.com/repos/cidverse/test/tags","teams_url":"https://api.github.com/repos/cidverse/test/teams","topics":[],"trees_url":"https://api.github.com/repos/cidverse/test/git/trees{/sha}","updated_at":"2023-02-06T20:25:57Z","url":"https://api.github.com/repos/cidverse/test","visibility":"public","watchers":0,"watchers_count":0,"web_commit_signoff_required":false},"sender":{"avatar_url":"https://avatars.githubusercontent.com/u/10275049?v=4","events_url":"https://api.github.com/users/PhilippHeuer/events{/privacy}","followers_url":"https://api.github.com/users/PhilippHeuer/followers","following_url":"https://api.github.com/users/PhilippHeuer/following{/other_user}","gists_url":"https://api.github.com/users/PhilippHeuer/gists{/gist_id}","gravatar_id":"","html_url":"https://github.com/PhilippHeuer","id":10275049,"login":"PhilippHeuer","node_id":"MDQ6VXNlcjEwMjc1MDQ5","organizations_url":"https://api.github.com/users/PhilippHeuer/orgs","received_events_url":"https://api.github.com/users/PhilippHeuer/received_events","repos_url":"https://api.github.com/users/PhilippHeuer/repos","site_admin":false,"starred_url":"https://api.github.com/users/PhilippHeuer/starred{/owner}{/repo}","subscriptions_url":"https://api.github.com/users/PhilippHeuer/subscriptions","type":"User","url":"https://api.github.com/users/PhilippHeuer"},"workflow":".github/workflows/ci.yml"}`

//go:embed examples/workflow_dispatch.json
var workflowDispatchJSON string

//go:embed examples/pullrequest.json
var pullRequestJSON string

func TestParseGithubWorkflowDispatchEvent(t *testing.T) {
	// create a temporary file with a test event
	eventFile, err := os.CreateTemp("", "github-event-test-*.json")
	if err != nil {
		t.Fatalf("Failed to create temporary file: %s", err)
	}
	defer os.Remove(eventFile.Name())
	if _, err := eventFile.Write([]byte(workflowDispatchJSON)); err != nil {
		t.Fatalf("Failed to write test event data to temporary file: %s", err)
	}
	if err := eventFile.Close(); err != nil {
		t.Fatalf("Failed to close temporary file: %s", err)
	}

	// Call the function and check the result
	event, err := ParseGithubEvent("PushEvent", eventFile.Name())
	if err != nil {
		t.Fatalf("Failed to parse test event: %s", err)
	}
	assert.NotNil(t, event)
}

func TestParseGithubPREvent(t *testing.T) {
	// create a temporary file with a test event
	eventFile, err := os.CreateTemp("", "github-event-test-*.json")
	if err != nil {
		t.Fatalf("Failed to create temporary file: %s", err)
	}
	defer os.Remove(eventFile.Name())
	if _, err := eventFile.Write([]byte(pullRequestJSON)); err != nil {
		t.Fatalf("Failed to write test event data to temporary file: %s", err)
	}
	if err := eventFile.Close(); err != nil {
		t.Fatalf("Failed to close temporary file: %s", err)
	}

	// Call the function and check the result
	event, err := ParseGithubEvent("PullRequestEvent", eventFile.Name())
	if err != nil {
		t.Fatalf("Failed to parse test event: %s", err)
	}
	assert.NotNil(t, event)

	// cast
	githubEvent, ok := event.(*github.PullRequestEvent)
	assert.True(t, ok)

	// content
	assert.Equal(t, 17, githubEvent.PullRequest.GetNumber())
	assert.Equal(t, "311b1ba11b054c4aab8baeca5ea21efb0e591380", githubEvent.PullRequest.GetBase().GetSHA())
}
