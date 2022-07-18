package vcsrepository

import (
	"github.com/cidverse/normalizeci/pkg/common"
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

func TestFindGitCommitsBetweenRefs(t *testing.T) {
	projectDir := FindRepositoryDirectory(common.GetWorkingDirectory())

	commits, commitsErr := FindCommitsBetweenRefs(projectDir, "refs/tags/v1.0.0", "refs/tags/v0.9.0")
	assert.NoError(t, commitsErr)
	assert.NotNil(t, commits)
	assert.Equal(t, 2, len(commits))

	// commit 1
	assert.Equal(t, "chore: update workflow script", commits[0].Message)
	assert.Equal(t, "\n", commits[0].Description)
	assert.Equal(t, "Philipp Heuer", commits[0].Author.Name)
	assert.Equal(t, "git@philippheuer.me", commits[0].Author.Email)
	assert.Equal(t, "Philipp Heuer", commits[0].Committer.Name)
	assert.Equal(t, "git@philippheuer.me", commits[0].Committer.Email)
	assert.Equal(t, "c1604a3", commits[0].ShortHash)
	assert.Equal(t, "c1604a3bf7b1b686608616206e357b1aae07ec45", commits[0].Hash)
	assert.Equal(t, int64(1578348804000000000), commits[0].AuthoredAt.UnixNano())
	assert.Equal(t, 1, len(commits[0].Tags))
	assert.Equal(t, "v1.0.0", commits[0].Tags[0].Name)
	assert.Equal(t, "refs/tags/v1.0.0", commits[0].Tags[0].VCSRef)
	assert.Nil(t, commits[0].Context)

	// commit 2
	assert.Equal(t, "fix: escape special chars in commit title / message and set default values for empty repos", commits[1].Message)
	assert.Equal(t, "\n", commits[1].Description)
	assert.Equal(t, "Philipp Heuer", commits[1].Author.Name)
	assert.Equal(t, "git@philippheuer.me", commits[1].Author.Email)
	assert.Equal(t, "Philipp Heuer", commits[1].Committer.Name)
	assert.Equal(t, "git@philippheuer.me", commits[1].Committer.Email)
	assert.Equal(t, "f3d7bd7", commits[1].ShortHash)
	assert.Equal(t, "f3d7bd736652725711fc4dc1dab0b3206ec4d3ae", commits[1].Hash)
	assert.Equal(t, int64(1578348473000000000), commits[1].AuthoredAt.UnixNano())
	assert.Nil(t, commits[1].Context)
}

func TestFindLatestGitReleaseFromCommit(t *testing.T) {
	projectDir := FindRepositoryDirectory(common.GetWorkingDirectory())

	release, releaseErr := FindLatestRelease(projectDir, "refs/tags/v1.0.0", true, false)
	assert.NoError(t, releaseErr)
	assert.NotNil(t, release)
	assert.Equal(t, "1.0.0", release.Name)
	assert.Equal(t, "refs/tags/v1.0.0", release.Reference)
}

func TestFindLatestGitReleaseSkipAndUnstable(t *testing.T) {
	projectDir := FindRepositoryDirectory(common.GetWorkingDirectory())

	release, releaseErr := FindLatestRelease(projectDir, "refs/tags/v1.0.0", true, true)
	assert.NoError(t, releaseErr)
	assert.NotNil(t, release)
	assert.Equal(t, "0.9.0", release.Name)
	assert.Equal(t, "refs/tags/v0.9.0", release.Reference)
}
