package projectdetails

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/cidverse/normalizeci/pkg/ncispec"
	"github.com/cidverse/normalizeci/pkg/normalizer/common"
	"github.com/gosimple/slug"
	"github.com/xanzy/go-gitlab"
)

var gitlabMockClient *http.Client

func GetGitLabToken(host string) string {
	if len(os.Getenv(common.ToEnvName(host)+"_TOKEN")) > 0 {
		return os.Getenv(common.ToEnvName(host) + "_TOKEN")
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
	gitlabUrl := "https://" + host

	// client
	gitlabClient, gitlabClientErr := gitlab.NewClient(glToken, gitlab.WithBaseURL(gitlabUrl))
	if gitlabClientErr != nil {
		return nil, gitlabClientErr
	}
	if gitlabMockClient != nil {
		gitlabClient, _ = gitlab.NewClient(
			glToken,
			gitlab.WithBaseURL(gitlabUrl),
			gitlab.WithHTTPClient(gitlabMockClient),
		)
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
	projectDetails[ncispec.NCI_PROJECT_URL] = project.WebURL

	return projectDetails, nil
}
