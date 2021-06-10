package projectdetails

import (
	"context"
	"github.com/google/go-github/v35/github"
	"github.com/gosimple/slug"
	"golang.org/x/oauth2"
	"os"
	"strconv"
	"strings"
)

func GetProjectDetailsGitHub(repoRemote string) (map[string]string, error) {
	projectDetails := make(map[string]string)

	repoPath := strings.TrimSuffix(strings.TrimPrefix(strings.TrimPrefix(repoRemote, "https://github.com/"), "git@github.com:"), ".git")
	repoPathSplit := strings.SplitN(repoPath, "/", 2)

	ctx := context.Background()
	client := github.NewClient(nil)
	if len(os.Getenv("GITHUB_TOKEN")) > 0 {
		ts := oauth2.StaticTokenSource(
			&oauth2.Token{AccessToken: os.Getenv("GITHUB_TOKEN")},
		)
		tc := oauth2.NewClient(ctx, ts)
		client = github.NewClient(tc)
	}

	repo, _, repoErr := client.Repositories.Get(ctx, repoPathSplit[0], repoPathSplit[1])
	if repoErr != nil {
		return nil, repoErr
	}

	projectDetails["NCI_PROJECT_ID"] = strconv.FormatInt(*repo.ID, 10)
	projectDetails["NCI_PROJECT_NAME"] = *repo.Name
	projectDetails["NCI_PROJECT_SLUG"] = slug.Make(*repo.FullName)
	projectDetails["NCI_PROJECT_DESCRIPTION"] = *repo.Description
	projectDetails["NCI_PROJECT_TOPICS"] = strings.Join(repo.Topics, ",")
	projectDetails["NCI_PROJECT_ISSUE_URL"] = strings.Replace(*repo.IssuesURL, "{/number}", "/{ID}", 1)
	projectDetails["NCI_PROJECT_STARGAZERS"] = strconv.Itoa(*repo.StargazersCount)
	projectDetails["NCI_PROJECT_FORKS"] = strconv.Itoa(*repo.ForksCount)

	return projectDetails, nil
}