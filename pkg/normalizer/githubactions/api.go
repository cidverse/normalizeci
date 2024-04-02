package githubactions

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/google/go-github/v61/github"
	"golang.org/x/oauth2"
)

var githubMockClient *http.Client

// GetGithubWorkflowRun retrieves a GitHub workflow run and its associated workflow for a given repository path
// and run ID. It returns the resulting GitHub workflow run and workflow objects, along with an error, if any.
//
// The function requires a valid GitHub access token to be set in the GITHUB_TOKEN environment variable.
//
// Parameters:
//   - repositoryPath: The repository path should be in the format "owner/repo"
//   - runId: a string representation of the numeric run ID for the workflow
//
// Returns:
//   - *github.WorkflowRun: A pointer to the retrieved GitHub workflow run object.
//   - *github.Workflow: A pointer to the retrieved GitHub workflow object associated with the workflow run.
//   - error: An error value, if any.
func GetGithubWorkflowRun(repositoryPath string, runId string) (*github.WorkflowRun, *github.Workflow, error) {
	if repositoryPath == "" {
		return nil, nil, fmt.Errorf("no repositoryPath provided")
	}
	rPath := strings.SplitN(repositoryPath, "/", 2)
	owner := rPath[0]
	name := rPath[1]
	if owner == "" || name == "" {
		return nil, nil, fmt.Errorf("invalid repositoryPath provided: %s", repositoryPath)
	}

	// GitHub client
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("GITHUB_TOKEN")},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)
	if githubMockClient != nil {
		client = github.NewClient(githubMockClient)
	}

	// parse runID
	runID, err := strconv.ParseInt(runId, 10, 64)
	if err != nil {
		return nil, nil, fmt.Errorf("parsing run ID %q: %w", runID, err)
	}

	// query run
	workflowRun, _, err := client.Actions.GetWorkflowRunByID(ctx, owner, name, runID)
	if err != nil {
		return nil, nil, fmt.Errorf("getting workflow run: %w", err)
	}

	// query workflow
	workflow, _, err := client.Actions.GetWorkflowByID(ctx, owner, name, workflowRun.GetWorkflowID())
	if err != nil {
		return nil, nil, fmt.Errorf("getting workflow: %w", err)
	}

	return workflowRun, workflow, nil
}

// ParseGithubEvent reads a JSON file containing a GitHub event
//
// Parameters:
//   - eventType: the GitHub event type
//   - eventFile: the GitHub event json file
//
// Returns:
//   - github.Event: A struct representing the parsed GitHub event.
//   - error: An error value, if any. If an error occurs while reading or parsing the file, it will be returned along with an informative error message.
func ParseGithubEvent(eventType string, eventFile string) (interface{}, error) {
	// maps GITHUB_EVENT_NAME to the corresponding event type
	if eventType == "pull_request" {
		eventType = "PullRequestEvent"
	} else if eventType == "workflow_dispatch" {
		eventType = "WorkflowDispatchEvent"
	}

	// file exists?
	if _, err := os.Stat(eventFile); os.IsNotExist(err) {
		return github.Event{}, fmt.Errorf("GITHUB_EVENT_PATH file does not exist: %w", err)
	}

	// read payload
	eventJSONBytes, err := os.ReadFile(eventFile) // just pass the file name
	if err != nil {
		return github.Event{}, fmt.Errorf("failed to read GITHUB_EVENT_PATH file: %w", err)
	}
	eventJSON := json.RawMessage(eventJSONBytes)
	event := github.Event{Type: &eventType, RawPayload: &eventJSON}

	// parse payload
	payload, err := event.ParsePayload()
	if err != nil {
		return github.Event{}, fmt.Errorf("failed to parse event payload: %w", err)
	}

	return payload, nil
}
