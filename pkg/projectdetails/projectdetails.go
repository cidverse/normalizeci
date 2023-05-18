package projectdetails

import (
	"fmt"

	v1 "github.com/cidverse/normalizeci/pkg/ncispec/v1"
)

func GetProjectDetails(repoType string, repoRemote string, hostType string, hostServer string) (v1.Project, error) {
	if repoType == "git" {
		if hostType == "github" {
			projectDetails, projectDetailsErr := GetProjectDetailsGitHub(hostServer, repoRemote)
			if projectDetailsErr == nil {
				return projectDetails, nil
			}
		} else if hostType == "gitlab" {
			projectDetails, projectDetailsErr := GetProjectDetailsGitLab(hostServer, repoRemote)
			if projectDetailsErr == nil {
				return projectDetails, nil
			}
		}
	}

	return v1.Project{}, fmt.Errorf("repository host type %s is not supported", hostType)
}
