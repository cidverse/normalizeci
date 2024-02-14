package projectdetails

import (
	"fmt"
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

func TestRepoPathFromRemote(t *testing.T) {
	testCases := []struct {
		remote string
		host   string
		expect string
	}{
		{"https://gitlab.com/user/project.git", "gitlab.com", "user/project"},
		{"git@gitlab.com:user/project.git", "gitlab.com", "user/project"},
		{"https://example.com/user/repo.git", "example.com", "user/repo"},
		{"git@example.com:user/repo.git", "example.com", "user/repo"},
		{"https://username:password@gitlab.com/user/project.git", "gitlab.com", "user/project"},
		{"https://github.com/user/project.git", "github.com", "user/project"},
		{"git@github.com:user/project.git", "github.com", "user/project"},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Remote: %s, Host: %s", tc.remote, tc.host), func(t *testing.T) {
			result, _ := repoPathFromRemote(tc.remote, tc.host)

			if result != tc.expect {
				t.Errorf("Expected: %s, Got: %s", tc.expect, result)
			}
		})
	}
}
