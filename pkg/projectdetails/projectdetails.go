package projectdetails

import (
	"os"
	"strings"
)

func GetProjectDetails(repoType string, repoRemote string) map[string]string {
	if repoType == "git" {
		// github
		if strings.HasPrefix(repoRemote, "https://github.com") || strings.HasPrefix(repoRemote, "git@github.com:") {
			projectDetails, projectDetailsErr := GetProjectDetailsGitHub("github.com", repoRemote)
			if projectDetailsErr == nil {
				return projectDetails
			}
		}

		// gitlab
		if strings.HasPrefix(repoRemote, "https://gitlab.com") || strings.HasPrefix(repoRemote, "git@gitlab.com:") {
			projectDetails, projectDetailsErr := GetProjectDetailsGitLab("gitlab.com", repoRemote)
			if projectDetailsErr == nil {
				return projectDetails
			}
		}

		// self-hosted instances
		host, hostErr := GetHostFromGitRemote(repoRemote)
		if hostErr == nil {
			if os.Getenv(ToEnvName(host)+"_TYPE") == "gitlab" {
				projectDetails, projectDetailsErr := GetProjectDetailsGitLab(host, repoRemote)
				if projectDetailsErr == nil {
					return projectDetails
				}
			}
		}
	}

	return nil
}
