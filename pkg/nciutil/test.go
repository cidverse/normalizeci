package nciutil

import (
	"testing"

	"github.com/cidverse/go-vcs"
	"github.com/cidverse/go-vcs/mocks"
	"github.com/cidverse/go-vcs/vcsapi"
	v1 "github.com/cidverse/normalizeci/pkg/ncispec/v1"
	"github.com/cidverse/normalizeci/pkg/projectdetails"
)

func MockVCSClient(t *testing.T) vcsapi.Client {
	mockClient := mocks.NewClient(t)
	mockClient.On("VCSType").Return("git")
	mockClient.On("VCSRemote").Return("https://github.com/cidverse/normalizeci.git")
	mockClient.On("VCSHostServer", "https://github.com/cidverse/normalizeci.git").Return("github.com")
	mockClient.On("VCSHostType", "github.com").Return("github")
	mockClient.On("VCSHead").Return(vcsapi.VCSRef{Type: "branch", Value: "main"}, nil)
	mockClient.On("VCSRefToInternalRef", vcsapi.VCSRef{Type: "branch", Value: "main"}).Return("refs/heads/main", nil)
	mockClient.On("FindCommitByHash", "", false).Return(vcsapi.Commit{}, nil)

	vcs.MockClient = mockClient
	projectdetails.MockProjectDetails = &v1.Project{
		Id:            "205438004",
		Name:          "normalizeci",
		Path:          "cidverse/normalizeci",
		Slug:          "cidverse-normalizeci",
		Description:   "A tool to turn the continuous integration / deployment variables into a common format for generally usable scripts without any dependencies.",
		Topics:        "cicd,normalization",
		IssueUrl:      "https://api.github.com/repos/cidverse/normalizeci/issues/{ID}",
		Stargazers:    "5",
		Forks:         "3",
		Dir:           "",
		Url:           "https://github.com/cidverse/normalizeci",
		DefaultBranch: "main",
	}

	return mockClient
}
