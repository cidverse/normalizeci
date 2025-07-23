package projectdetails

import (
	"context"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/cidverse/go-ptr"
	v1 "github.com/cidverse/normalizeci/pkg/ncispec/v1"
	"github.com/cidverse/normalizeci/pkg/normalizer/api"
	"github.com/google/go-github/v74/github"
	"github.com/gosimple/slug"
	"golang.org/x/oauth2"
)

var githubMockClient *http.Client

func GetProjectDetailsGitHub(host string, repoRemote string) (v1.Project, error) {
	result := v1.Project{}
	repoPath, err := repoPathFromRemote(repoRemote, host)
	if err != nil {
		return result, err
	}
	repoPathSplit := strings.SplitN(repoPath, "/", 2)

	ghToken := ""
	if len(os.Getenv(api.ToEnvName(host)+"_TOKEN")) > 0 {
		ghToken = os.Getenv(api.ToEnvName(host) + "_TOKEN")
	} else if len(os.Getenv("GITHUB_TOKEN")) > 0 {
		ghToken = os.Getenv("GITHUB_TOKEN")
	}

	ctx := context.Background()
	client := github.NewClient(nil)
	if len(ghToken) > 0 {
		ts := oauth2.StaticTokenSource(
			&oauth2.Token{AccessToken: ghToken},
		)
		tc := oauth2.NewClient(ctx, ts)
		client = github.NewClient(tc)
	}
	if githubMockClient != nil {
		client = github.NewClient(githubMockClient)
	}

	repo, _, repoErr := client.Repositories.Get(ctx, repoPathSplit[0], repoPathSplit[1])
	if repoErr != nil {
		return result, repoErr
	}

	result.ID = strconv.FormatInt(ptr.Value(repo.ID), 10)
	result.Name = ptr.Value(repo.Name)
	result.Path = ptr.Value(repo.FullName)
	result.Slug = slug.Make(ptr.Value(repo.FullName))
	result.Namespace = ptr.Value(repo.Owner.Login)
	result.NamespaceSlug = slug.Make(ptr.Value(repo.Owner.Login))
	result.Description = ptr.Value(repo.Description)
	result.Topics = strings.Join(repo.Topics, ",")
	result.IssueUrl = strings.Replace(ptr.Value(repo.IssuesURL), "{/number}", "/{ID}", 1)
	result.Stargazers = strconv.Itoa(ptr.Value(repo.StargazersCount))
	result.Forks = strconv.Itoa(ptr.Value(repo.ForksCount))
	result.DefaultBranch = ptr.Value(repo.DefaultBranch)
	result.Url = strings.TrimSuffix(ptr.Value(repo.CloneURL), ".git")

	return result, nil
}
