package projectdetails

import (
	"fmt"
	"net/url"
	"strings"
)

func GetHostFromGitRemote(input string) (string, error) {
	if !strings.HasPrefix(input, "http://") && !strings.HasPrefix(input, "https://") && strings.Contains(input, "@") {
		parts := strings.Split(input, "@")
		if len(parts) != 2 {
			return "", fmt.Errorf("error parsing git ssh remote: %v", input)
		}
		path := parts[1]
		hostParts := strings.Split(path, ":")
		return hostParts[0], nil
	}

	u, err := url.Parse(input)
	if err != nil {
		return "", fmt.Errorf("error parsing URL: %v", err)
	}

	return u.Host, nil
}

func repoPathFromRemote(remote string, host string) (string, error) {
	// ssh
	if strings.HasPrefix(remote, "git@") {
		return strings.TrimSuffix(strings.TrimSuffix(strings.TrimPrefix(remote, fmt.Sprintf("git@%s:", host)), ".git"), "/"), nil
	}

	// http (parse url and return path without .git suffix)
	u, err := url.Parse(remote)
	if err != nil {
		return "", fmt.Errorf("parsing remote url: %w", err)
	}

	return strings.TrimSuffix(strings.TrimPrefix(u.Path, "/"), ".git"), nil
}
