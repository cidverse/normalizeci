package gitlabci

import (
	"os"
	"testing"

	"github.com/cidverse/go-vcs/mocks"
	"github.com/cidverse/go-vcs/vcsapi"
	"github.com/rs/zerolog"
)

func TestMain(m *testing.M) {
	zerolog.SetGlobalLevel(zerolog.DebugLevel)
	code := m.Run()
	os.Exit(code)
}

func MockVCSClient(t *testing.T) vcsapi.Client {
	mockClient := mocks.NewClient(t)
	mockClient.On("VCSType").Return("git")
	mockClient.On("VCSRemote").Return("https://github.com/cidverse/normalizeci.git")
	mockClient.On("VCSHostServer", "https://github.com/cidverse/normalizeci.git").Return("github.com")
	mockClient.On("VCSHostType", "github.com").Return("github")
	mockClient.On("VCSHead").Return(vcsapi.VCSRef{Type: "branch", Value: "main"}, nil)
	mockClient.On("VCSRefToInternalRef", vcsapi.VCSRef{Type: "branch", Value: "main"}).Return("refs/heads/main", nil)
	mockClient.On("FindCommitByHash", "", false).Return(vcsapi.Commit{}, nil)

	return mockClient
}
