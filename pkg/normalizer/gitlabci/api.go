package gitlabci

import (
	"net/http"
	"strconv"

	"gitlab.com/gitlab-org/api/client-go"
)

var gitlabMockClient *http.Client

func GetGitlabPipelineRun(server string, project string, pipelineIdText string, token string) ([]*gitlab.PipelineVariable, error) {
	// client
	client, clientErr := gitlab.NewClient(token, gitlab.WithBaseURL(server))
	if clientErr != nil {
		return nil, clientErr
	}

	if gitlabMockClient != nil {
		client, _ = gitlab.NewClient(
			token,
			gitlab.WithBaseURL(server),
			gitlab.WithHTTPClient(gitlabMockClient),
		)
	}

	// query
	pipelineId, err := strconv.Atoi(pipelineIdText)
	if err != nil {
		return nil, err
	}
	variables, _, err := client.Pipelines.GetPipelineVariables(project, int64(pipelineId))
	if err != nil {
		return nil, err
	}

	return variables, nil
}
