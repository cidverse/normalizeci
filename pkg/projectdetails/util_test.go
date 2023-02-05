package projectdetails

import (
	"testing"
)

func TestGetHostFromGitRemote(t *testing.T) {
	testCases := []struct {
		input  string
		output string
		err    string
	}{
		{"https://github.com/user/repo.git", "github.com", ""},
		{"git@github.com:user/repo.git", "github.com", ""},
		{"git://github.com/user/repo.git", "github.com", ""},
		{"https://user:password@github.com/user/repo.git", "github.com", ""},
		{"", "", "Error parsing URL: parse : empty url"},
	}

	for _, testCase := range testCases {
		result, err := GetHostFromGitRemote(testCase.input)
		if err != nil {
			if err.Error() != testCase.err {
				t.Errorf("GetHostFromGitRemote(%q) error = %q, expected %q", testCase.input, err.Error(), testCase.err)
			}
		} else if result != testCase.output {
			t.Errorf("GetHostFromGitRemote(%q) = %q, expected %q", testCase.input, result, testCase.output)
		}
	}
}
