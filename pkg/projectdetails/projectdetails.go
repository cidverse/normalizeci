package projectdetails

func GetProjectDetails(repoType string, repoRemote string, hostType string, hostServer string) map[string]string {
	if repoType == "git" {
		if hostType == "github" {
			projectDetails, projectDetailsErr := GetProjectDetailsGitHub(hostServer, repoRemote)
			if projectDetailsErr == nil {
				return projectDetails
			}
		} else if hostType == "gitlab" {
			projectDetails, projectDetailsErr := GetProjectDetailsGitLab(hostServer, repoRemote)
			if projectDetailsErr == nil {
				return projectDetails
			}
		}
	}

	return make(map[string]string)
}
