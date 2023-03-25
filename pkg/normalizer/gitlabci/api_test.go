package gitlabci

import (
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func TestGetGitlabPipelineRun(t *testing.T) {
	gitlabMockClient = &http.Client{}
	httpmock.ActivateNonDefault(gitlabMockClient)
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("GET", "https://gitlab.com/api/v4/projects/43228743/pipelines/801916361/variables", httpmock.NewStringResponder(200, `[{"variable_type":"file","key":"afile","value":"example text","raw":false},{"variable_type":"env_var","key":"hello","value":"world","raw":false},{"variable_type":"env_var","key":"name","value":"my-name","raw":false}]`))

	// call function
	variables, err := GetGitlabPipelineRun("https://gitlab.com", "43228743", "801916361", "invalid-token")
	assert.NoError(t, err)
	assert.NotNil(t, variables)
	assert.Len(t, variables, 3)
	assert.Equal(t, "afile", variables[0].Key)
	assert.Equal(t, "example text", variables[0].Value)
	assert.Equal(t, "hello", variables[1].Key)
	assert.Equal(t, "world", variables[1].Value)
	assert.Equal(t, "name", variables[2].Key)
	assert.Equal(t, "my-name", variables[2].Value)
}
