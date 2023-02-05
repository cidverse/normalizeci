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

func ToEnvName(input string) string {
	return strings.Replace(strings.ToUpper(input), ".", "_", -1)
}
