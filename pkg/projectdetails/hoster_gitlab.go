package projectdetails

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/cidverse/normalizeci/pkg/ncispec"
	"github.com/gosimple/slug"
	"github.com/hashicorp/go-retryablehttp"
	"github.com/xanzy/go-gitlab"
)

func GetGitLabToken(host string) string {
	if len(os.Getenv(ToEnvName(host)+"_TOKEN")) > 0 {
		return os.Getenv(ToEnvName(host) + "_TOKEN")
	} else if len(os.Getenv("GITLAB_TOKEN")) > 0 {
		return os.Getenv("GITLAB_TOKEN")
	} else if os.Getenv("CI") == "true" && len(os.Getenv("CI_BUILD_TOKEN")) > 0 {
		return os.Getenv("CI_BUILD_TOKEN")
	} else if os.Getenv("CI") == "true" && len(os.Getenv("CI_JOB_TOKEN")) > 0 {
		return os.Getenv("CI_JOB_TOKEN")
	}

	return ""
}

func GetProjectDetailsGitLab(host string, repoRemote string) (map[string]string, error) {
	projectDetails := make(map[string]string)
	repoPath := strings.TrimSuffix(strings.TrimPrefix(strings.TrimPrefix(repoRemote, fmt.Sprintf("https://%s/", host)), fmt.Sprintf("git@%s:", host)), ".git")
	glToken := GetGitLabToken(host)

	// client
	gitlabClient, gitlabClientErr := gitlab.NewClient(glToken, gitlab.WithBaseURL("https://"+host))
	if gitlabClientErr != nil {
		return nil, gitlabClientErr
	}

	// client with InsecureSkipVerify for .local domains
	if strings.HasSuffix(host, ".local") {
		tr := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
		retryClient := retryablehttp.NewClient()
		retryClient.HTTPClient.Transport = tr
		gitlabClient, gitlabClientErr = gitlab.NewClient(
			glToken,
			gitlab.WithBaseURL("https://"+host),
			gitlab.WithHTTPClient(retryClient.HTTPClient),
		)
		if gitlabClientErr != nil {
			return nil, gitlabClientErr
		}
	}

	// query project
	project, _, projectErr := gitlabClient.Projects.GetProject(repoPath, &gitlab.GetProjectOptions{})
	if projectErr != nil {
		return nil, projectErr
	}

	projectDetails[ncispec.NCI_PROJECT_ID] = strconv.Itoa(project.ID)
	projectDetails[ncispec.NCI_PROJECT_NAME] = project.Name
	projectDetails[ncispec.NCI_PROJECT_PATH] = project.NameWithNamespace
	projectDetails[ncispec.NCI_PROJECT_SLUG] = slug.Make(project.NameWithNamespace)
	projectDetails[ncispec.NCI_PROJECT_DESCRIPTION] = project.Description
	projectDetails[ncispec.NCI_PROJECT_TOPICS] = strings.Join(project.TagList, ",")
	projectDetails[ncispec.NCI_PROJECT_ISSUE_URL] = project.WebURL + "/-/issues/{ID}"
	projectDetails[ncispec.NCI_PROJECT_STARGAZERS] = strconv.Itoa(project.StarCount)
	projectDetails[ncispec.NCI_PROJECT_FORKS] = strconv.Itoa(project.ForksCount)
	projectDetails[ncispec.NCI_PROJECT_DEFAULT_BRANCH] = project.DefaultBranch

	return projectDetails, nil
}
