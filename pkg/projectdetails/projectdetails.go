package projectdetails

import (
	"fmt"

	v1 "github.com/cidverse/normalizeci/pkg/ncispec/v1"
)

var MockProjectDetails *v1.Project

func GetProjectDetails(repoType string, repoRemote string, hostType string, hostServer string) (v1.Project, error) {
	if MockProjectDetails != nil {
		return *MockProjectDetails, nil
	}

	if repoType == "git" {
		if hostType == "github" {
			projectDetails, err := GetProjectDetailsGitHub(hostServer, repoRemote)
			return projectDetails, err
		} else if hostType == "gitlab" {
			projectDetails, err := GetProjectDetailsGitLab(hostServer, repoRemote)
			return projectDetails, err
		}
	}

	return v1.Project{}, fmt.Errorf("repository host type %s is not supported", hostType)
}
