package projectdetails

import (
	"context"
	"net/http"
	"os"
	"strconv"
	"strings"

	v1 "github.com/cidverse/normalizeci/pkg/ncispec/v1"
	"github.com/cidverse/normalizeci/pkg/normalizer/api"
	"github.com/google/go-github/v55/github"
	"github.com/gosimple/slug"
	"golang.org/x/oauth2"
)

var githubMockClient *http.Client

func GetProjectDetailsGitHub(host string, repoRemote string) (v1.Project, error) {
	result := v1.Project{}

	repoPath := strings.TrimSuffix(strings.TrimPrefix(strings.TrimPrefix(repoRemote, "https://github.com/"), "git@github.com:"), ".git")
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

	result.Id = strconv.FormatInt(*repo.ID, 10)
	result.Name = *repo.Name
	result.Path = *repo.FullName
	result.Slug = slug.Make(*repo.FullName)
	if repo.Description != nil {
		result.Description = *repo.Description
	} else {
		result.Description = ""
	}
	result.Topics = strings.Join(repo.Topics, ",")
	result.IssueUrl = strings.Replace(*repo.IssuesURL, "{/number}", "/{ID}", 1)
	result.Stargazers = strconv.Itoa(*repo.StargazersCount)
	result.Forks = strconv.Itoa(*repo.ForksCount)
	result.DefaultBranch = *repo.DefaultBranch
	result.Url = strings.TrimSuffix(*repo.CloneURL, ".git")

	return result, nil
}
