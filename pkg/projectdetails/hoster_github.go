package projectdetails

import (
	"context"
	"github.com/cidverse/normalizeci/pkg/ncispec"
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

	ghToken := ""
	if len(os.Getenv("GITHUB_TOKEN")) > 0 {
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

	repo, _, repoErr := client.Repositories.Get(ctx, repoPathSplit[0], repoPathSplit[1])
	if repoErr != nil {
		return nil, repoErr
	}

	projectDetails[ncispec.NCI_PROJECT_ID] = strconv.FormatInt(*repo.ID, 10)
	projectDetails[ncispec.NCI_PROJECT_NAME] = *repo.Name
	projectDetails[ncispec.NCI_PROJECT_PATH] = *repo.FullName
	projectDetails[ncispec.NCI_PROJECT_SLUG] = slug.Make(*repo.FullName)
	if repo.Description != nil {
		projectDetails[ncispec.NCI_PROJECT_DESCRIPTION] = *repo.Description
	} else {
		projectDetails[ncispec.NCI_PROJECT_DESCRIPTION] = ""
	}
	projectDetails[ncispec.NCI_PROJECT_TOPICS] = strings.Join(repo.Topics, ",")
	projectDetails[ncispec.NCI_PROJECT_ISSUE_URL] = strings.Replace(*repo.IssuesURL, "{/number}", "/{ID}", 1)
	projectDetails[ncispec.NCI_PROJECT_STARGAZERS] = strconv.Itoa(*repo.StargazersCount)
	projectDetails[ncispec.NCI_PROJECT_FORKS] = strconv.Itoa(*repo.ForksCount)

	return projectDetails, nil
}
