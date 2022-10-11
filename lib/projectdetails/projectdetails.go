package projectdetails

import (
	"strings"
)

func GetProjectDetails(repoType string, repoRemote string) map[string]string {
	if repoType == "git" {
		// github
		if strings.HasPrefix(repoRemote, "https://github.com") || strings.HasPrefix(repoRemote, "git@github.com:") {
			projectDetails, projectDetailsErr := GetProjectDetailsGitHub(repoRemote)
			if projectDetailsErr == nil {
				return projectDetails
			}
		}
		// gitlab
		if strings.HasPrefix(repoRemote, "https://gitlab.com") || strings.HasPrefix(repoRemote, "git@gitlab.com:") {
			projectDetails, projectDetailsErr := GetProjectDetailsGitLab(repoRemote)
			if projectDetailsErr == nil {
				return projectDetails
			}
		}
	}

	return nil
}
