package projectdetails

import (
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func TestGetProjectDetailsGitLab(t *testing.T) {
	// http mock
	gitlabMockClient = &http.Client{}
	httpmock.ActivateNonDefault(gitlabMockClient)
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("GET", "https://gitlab.com/api/v4/projects/PhilippHeuer%2Fevents4j", httpmock.NewStringResponder(200, `{"id":6364957,"description":"Java Event Dispatcher / Consumer","name":"Events4J","name_with_namespace":"Philipp Heuer / Events4J","path":"events4j","path_with_namespace":"PhilippHeuer/events4j","created_at":"2018-05-13T20:14:54.777Z","default_branch":"master","tag_list":[],"topics":[],"ssh_url_to_repo":"git@gitlab.com:PhilippHeuer/events4j.git","http_url_to_repo":"https://gitlab.com/PhilippHeuer/events4j.git","web_url":"https://gitlab.com/PhilippHeuer/events4j","readme_url":"https://gitlab.com/PhilippHeuer/events4j/-/blob/master/README.md","forks_count":0,"avatar_url":null,"star_count":0,"last_activity_at":"2022-12-04T19:25:24.947Z","namespace":{"id":1423465,"name":"Philipp Heuer","path":"PhilippHeuer","kind":"user","full_path":"PhilippHeuer","parent_id":null,"avatar_url":"https://secure.gravatar.com/avatar/06a6a5b8addc909ff8139c369d1c0d7c?s=80&d=identicon","web_url":"https://gitlab.com/PhilippHeuer"}}`))

	details, _ := GetProjectDetailsGitLab("gitlab.com", "https://gitlab.com/PhilippHeuer/events4j.git")

	assert.Equal(t, "6364957", details.Id)
	assert.Equal(t, "Events4J", details.Name)
	assert.Equal(t, "philipp-heuer-events4j", details.Slug)
	assert.Equal(t, "Java Event Dispatcher / Consumer", details.Description)
	assert.Equal(t, "", details.Topics)
	assert.Equal(t, "https://gitlab.com/PhilippHeuer/events4j/-/issues/{ID}", details.IssueUrl)
	assert.NotEmpty(t, details.Stargazers)
	assert.NotEmpty(t, details.Forks)
	assert.Equal(t, "https://gitlab.com/PhilippHeuer/events4j", details.Url)
	assert.Equal(t, "master", details.DefaultBranch)
}
