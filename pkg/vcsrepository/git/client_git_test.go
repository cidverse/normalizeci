package gitclient

import (
	"strings"
	"testing"

	"github.com/cidverse/normalizeci/pkg/normalizer/common"
	"github.com/cidverse/normalizeci/pkg/vcsrepository/vcsapi"
	"github.com/stretchr/testify/assert"
)

func TestGetCommitByHashWithChanges(t *testing.T) {
	projectDir, _ := common.GetProjectDirectory()

	client, clientErr := NewGitClient(projectDir)
	assert.NoError(t, clientErr)
	assert.NotNil(t, client)

	commit, commitErr := client.FindCommitByHash("93591c7d4d56d961b09d44ec6a866e4c53a2b947", true)
	assert.NoError(t, commitErr)
	assert.NotNil(t, commit)

	assert.Equal(t, "93591c7", commit.ShortHash)
	assert.Equal(t, "93591c7d4d56d961b09d44ec6a866e4c53a2b947", commit.Hash)
	assert.Equal(t, "chore(deps): update module github.com/xanzy/go-gitlab to v0.76.0 (#20)", commit.Message)
	assert.Equal(t, "Co-authored-by: renovate[bot] <29139614+renovate[bot]@users.noreply.github.com>", commit.Description)
	assert.Equal(t, "renovate[bot]", commit.Author.Name)
	assert.Equal(t, "29139614+renovate[bot]@users.noreply.github.com", commit.Author.Email)
	assert.Equal(t, "GitHub", commit.Committer.Name)
	assert.Equal(t, "noreply@github.com", commit.Committer.Email)
	assert.Nil(t, commit.Context)

	assert.Equal(t, "update", commit.Changes[0].Type)
	assert.Equal(t, "go.mod", commit.Changes[0].FileFrom.Name)
	assert.Equal(t, "48f397f4b543c87e1c80992b0ff4ddc881a42eae", commit.Changes[0].FileFrom.Hash)
	assert.Equal(t, "go.mod", commit.Changes[0].FileTo.Name)
	assert.Equal(t, "3489cc8928678653ceadf47954a1c2f5d8f01e0c", commit.Changes[0].FileTo.Hash)
	assert.NotNil(t, commit.Changes[0].Patch)
}

func TestFindGitCommitsBetweenRefs(t *testing.T) {
	projectDir, _ := common.GetProjectDirectory()

	client, clientErr := NewGitClient(projectDir)
	assert.NoError(t, clientErr)
	assert.NotNil(t, client)

	commits, commitsErr := client.FindCommitsBetween(&vcsapi.VCSRef{Type: "tag", Value: "v1.0.0"}, &vcsapi.VCSRef{Type: "tag", Value: "v0.9.0"}, false, 0)
	assert.NoError(t, commitsErr)
	assert.NotNil(t, commits)
	assert.Equal(t, 2, len(commits))

	// commit 1
	assert.Equal(t, "chore: update workflow script", commits[0].Message)
	assert.Equal(t, "", commits[0].Description)
	assert.Equal(t, "Philipp Heuer", commits[0].Author.Name)
	assert.Equal(t, "git@philippheuer.me", commits[0].Author.Email)
	assert.Equal(t, "Philipp Heuer", commits[0].Committer.Name)
	assert.Equal(t, "git@philippheuer.me", commits[0].Committer.Email)
	assert.Equal(t, "c1604a3", commits[0].ShortHash)
	assert.Equal(t, "c1604a3bf7b1b686608616206e357b1aae07ec45", commits[0].Hash)
	assert.Equal(t, int64(1578348804000000000), commits[0].AuthoredAt.UnixNano())
	assert.Equal(t, 1, len(commits[0].Tags))
	assert.Equal(t, "tag", commits[0].Tags[0].Type)
	assert.Equal(t, "v1.0.0", commits[0].Tags[0].Value)
	assert.Nil(t, commits[0].Context)

	// commit 2
	assert.Equal(t, "fix: escape special chars in commit title / message and set default values for empty repos", commits[1].Message)
	assert.Equal(t, "", commits[1].Description)
	assert.Equal(t, "Philipp Heuer", commits[1].Author.Name)
	assert.Equal(t, "git@philippheuer.me", commits[1].Author.Email)
	assert.Equal(t, "Philipp Heuer", commits[1].Committer.Name)
	assert.Equal(t, "git@philippheuer.me", commits[1].Committer.Email)
	assert.Equal(t, "f3d7bd7", commits[1].ShortHash)
	assert.Equal(t, "f3d7bd736652725711fc4dc1dab0b3206ec4d3ae", commits[1].Hash)
	assert.Equal(t, int64(1578348473000000000), commits[1].AuthoredAt.UnixNano())
	assert.Equal(t, 0, len(commits[1].Tags))
	assert.Nil(t, commits[1].Context)
}

func TestFindGitCommitsBetweenHashRefs(t *testing.T) {
	projectDir, _ := common.GetProjectDirectory()

	client, clientErr := NewGitClient(projectDir)
	assert.NoError(t, clientErr)
	assert.NotNil(t, client)

	commits, commitsErr := client.FindCommitsBetween(&vcsapi.VCSRef{Type: "hash", Hash: "c1604a3bf7b1b686608616206e357b1aae07ec45"}, &vcsapi.VCSRef{Type: "tag", Value: "v0.9.0"}, false, 0)
	assert.NoError(t, commitsErr)
	assert.NotNil(t, commits)
	assert.Equal(t, 2, len(commits))

	// commit 1
	assert.Equal(t, "chore: update workflow script", commits[0].Message)
	assert.Equal(t, "", commits[0].Description)
	assert.Equal(t, "Philipp Heuer", commits[0].Author.Name)
	assert.Equal(t, "git@philippheuer.me", commits[0].Author.Email)
	assert.Equal(t, "Philipp Heuer", commits[0].Committer.Name)
	assert.Equal(t, "git@philippheuer.me", commits[0].Committer.Email)
	assert.Equal(t, "c1604a3", commits[0].ShortHash)
	assert.Equal(t, "c1604a3bf7b1b686608616206e357b1aae07ec45", commits[0].Hash)
	assert.Equal(t, int64(1578348804000000000), commits[0].AuthoredAt.UnixNano())
	assert.Equal(t, 1, len(commits[0].Tags))
	assert.Equal(t, "tag", commits[0].Tags[0].Type)
	assert.Equal(t, "v1.0.0", commits[0].Tags[0].Value)
	assert.Nil(t, commits[0].Context)

	// commit 2
	assert.Equal(t, "fix: escape special chars in commit title / message and set default values for empty repos", commits[1].Message)
	assert.Equal(t, "", commits[1].Description)
	assert.Equal(t, "Philipp Heuer", commits[1].Author.Name)
	assert.Equal(t, "git@philippheuer.me", commits[1].Author.Email)
	assert.Equal(t, "Philipp Heuer", commits[1].Committer.Name)
	assert.Equal(t, "git@philippheuer.me", commits[1].Committer.Email)
	assert.Equal(t, "f3d7bd7", commits[1].ShortHash)
	assert.Equal(t, "f3d7bd736652725711fc4dc1dab0b3206ec4d3ae", commits[1].Hash)
	assert.Equal(t, int64(1578348473000000000), commits[1].AuthoredAt.UnixNano())
	assert.Equal(t, 0, len(commits[1].Tags))
	assert.Nil(t, commits[1].Context)
}

func TestFindLatestGitReleaseFromCommit(t *testing.T) {
	projectDir, _ := common.GetProjectDirectory()

	client, clientErr := NewGitClient(projectDir)
	assert.NoError(t, clientErr)
	assert.NotNil(t, client)

	release, releaseErr := client.FindLatestRelease(true)
	assert.NoError(t, releaseErr)
	assert.NotNil(t, release)
	assert.Equal(t, "tag", release.Type)
	assert.True(t, true, strings.HasPrefix(release.Value, "v"))
	assert.Regexp(t, "v[0-9]+.[0-9]+.[0-9]+", release.Value)
	assert.Regexp(t, "[0-9]+.[0-9]+.[0-9]+", release.Version)
}

func TestParseGitRefLogLine_Tag(t *testing.T) {
	vcsRef := ParseGitRefLogLine("0000000000000000000000000000000000000000 1cafbbdb80ce27304ac92a9e2fde6c3df8119a19 runner <runner@fv-az554-304.(none)> 1679700466 +0000\tcheckout: moving from master to refs/tags/v2.0.0-alpha.1", "1cafbbdb80ce27304ac92a9e2fde6c3df8119a19")
	assert.Equal(t, "1cafbbdb80ce27304ac92a9e2fde6c3df8119a19", vcsRef.Hash)
	assert.Equal(t, "tag", vcsRef.Type)
	assert.Equal(t, "v2.0.0-alpha.1", vcsRef.Value)
}

func TestParseGitRefLogLine_LocalBranch(t *testing.T) {
	vcsRef := ParseGitRefLogLine("0000000000000000000000000000000000000000 1cafbbdb80ce27304ac92a9e2fde6c3df8119a19 runner <runner@fv-az554-304.(none)> 1679700466 +0000\tcheckout: moving from master to feature-branch", "1cafbbdb80ce27304ac92a9e2fde6c3df8119a19")
	assert.Equal(t, "1cafbbdb80ce27304ac92a9e2fde6c3df8119a19", vcsRef.Hash)
	assert.Equal(t, "branch", vcsRef.Type)
	assert.Equal(t, "feature-branch", vcsRef.Value)
}

func TestParseGitRefLogLine_Hash(t *testing.T) {
	vcsRef := ParseGitRefLogLine("0000000000000000000000000000000000000000 1cafbbdb80ce27304ac92a9e2fde6c3df8119a19 runner <runner@fv-az554-304.(none)> 1679700466 +0000\tcheckout: moving from master to 1cafbbdb80ce27304ac92a9e2fde6c3df8119a19", "1cafbbdb80ce27304ac92a9e2fde6c3df8119a19")
	assert.Equal(t, "1cafbbdb80ce27304ac92a9e2fde6c3df8119a19", vcsRef.Hash)
	assert.Equal(t, "branch", vcsRef.Type)
	assert.Equal(t, "master", vcsRef.Value)
}
