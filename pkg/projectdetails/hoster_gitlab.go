package projectdetails

import (
	"github.com/gosimple/slug"
	"github.com/xanzy/go-gitlab"
	"os"
	"strconv"
	"strings"
)

func GetProjectDetailsGitLab(repoRemote string) (map[string]string, error) {
	projectDetails := make(map[string]string)
	repoPath := strings.TrimSuffix(strings.TrimPrefix(strings.TrimPrefix(repoRemote, "https://gitlab.com/"), "git@gitlab.com:"), ".git")

	glToken := ""
	if os.Getenv("CI") == "true" && len(os.Getenv("CI_BUILD_TOKEN")) > 0 {
		glToken = os.Getenv("CI_BUILD_TOKEN")
	}
	if len(os.Getenv("GITLAB_TOKEN")) > 0 {
		glToken = os.Getenv("GITLAB_TOKEN")
	}

	gitlabClient, gitlabClientErr := gitlab.NewClient(glToken, gitlab.WithBaseURL("https://gitlab.com"))
	if gitlabClientErr != nil {
		return nil, gitlabClientErr
	}

	project, _, projectErr := gitlabClient.Projects.GetProject(repoPath, &gitlab.GetProjectOptions{})
	if projectErr != nil {
		return nil, projectErr
	}

	projectDetails["NCI_PROJECT_ID"] = strconv.Itoa(project.ID)
	projectDetails["NCI_PROJECT_NAME"] = project.Name
	projectDetails["NCI_PROJECT_SLUG"] = slug.Make(project.NameWithNamespace)
	projectDetails["NCI_PROJECT_DESCRIPTION"] = project.Description
	projectDetails["NCI_PROJECT_TOPICS"] = strings.Join(project.TagList, ",")
	projectDetails["NCI_PROJECT_ISSUE_URL"] = project.WebURL + "/-/issues/{ID}"
	projectDetails["NCI_PROJECT_STARGAZERS"] = strconv.Itoa(project.StarCount)
	projectDetails["NCI_PROJECT_FORKS"] = strconv.Itoa(project.ForksCount)

	return projectDetails, nil
}
