package projectdetails

import (
	"net/http"
	"os"
	"strconv"
	"strings"

	v1 "github.com/cidverse/normalizeci/pkg/ncispec/v1"
	"github.com/cidverse/normalizeci/pkg/normalizer/api"
	"github.com/gosimple/slug"
	"gitlab.com/gitlab-org/api/client-go"
)

var gitlabMockClient *http.Client

func GetGitLabToken(host string) string {
	if len(os.Getenv(api.ToEnvName(host)+"_TOKEN")) > 0 {
		return os.Getenv(api.ToEnvName(host) + "_TOKEN")
	} else if len(os.Getenv("GITLAB_TOKEN")) > 0 {
		return os.Getenv("GITLAB_TOKEN")
	} else if os.Getenv("CI") == "true" && len(os.Getenv("CI_JOB_TOKEN")) > 0 {
		return os.Getenv("CI_JOB_TOKEN")
	}

	return ""
}

func GetProjectDetailsGitLab(host string, repoRemote string) (v1.Project, error) {
	result := v1.Project{}
	repoPath, err := repoPathFromRemote(repoRemote, host)
	if err != nil {
		return result, err
	}
	glToken := GetGitLabToken(host)
	gitlabUrl := "https://" + host

	// client
	gitlabClient, gitlabClientErr := gitlab.NewClient(glToken, gitlab.WithBaseURL(gitlabUrl))
	if gitlabClientErr != nil {
		return result, gitlabClientErr
	}
	if gitlabMockClient != nil {
		gitlabClient, _ = gitlab.NewClient("", gitlab.WithBaseURL(gitlabUrl), gitlab.WithHTTPClient(gitlabMockClient))
	}

	// query project
	project, _, projectErr := gitlabClient.Projects.GetProject(repoPath, &gitlab.GetProjectOptions{})
	if projectErr != nil {
		return result, projectErr
	}

	result.ID = strconv.Itoa(project.ID)
	result.Name = project.Name
	result.Path = project.NameWithNamespace
	result.Slug = slug.Make(project.NameWithNamespace)
	result.Namespace = project.Namespace.Path
	result.NamespaceSlug = slug.Make(project.Namespace.Path)
	result.Description = project.Description
	result.Topics = strings.Join(project.TagList, ",")
	result.IssueUrl = project.WebURL + "/-/issues/{ID}"
	result.Stargazers = strconv.Itoa(project.StarCount)
	result.Forks = strconv.Itoa(project.ForksCount)
	result.DefaultBranch = project.DefaultBranch
	result.Url = project.WebURL

	return result, nil
}
