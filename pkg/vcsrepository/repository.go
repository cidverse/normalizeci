package vcsrepository

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
)

func FindRepositoryDirectory(currentDirectory string) string {
	var projectDirectory = ""
	directoryParts := strings.Split(currentDirectory, string(os.PathSeparator))

	for projectDirectory == "" {
		// GIT
		if IsGitRepository(currentDirectory) {
			return currentDirectory
		}

		// abort when we reach the root directory, no repository found
		if directoryParts[0]+"\\" == currentDirectory || currentDirectory == "/" {
			return ""
		}

		// check parent directory in next iteration
		currentDirectory = filepath.Dir(currentDirectory)
	}

	return ""
}

// GetVCSRepositoryInformation detects the repository type and gathers normalized information from the repository
func GetVCSRepositoryInformation(dir string) (map[string]string, error) {
	// init with default values
	data := make(map[string]string)
	data["NCI_REPOSITORY_KIND"] = "none"
	data["NCI_REPOSITORY_REMOTE"] = "local"
	data["NCI_COMMIT_REF_TYPE"] = "unknown"
	data["NCI_COMMIT_REF_NAME"] = "unknown"
	data["NCI_COMMIT_REF_SLUG"] = ""
	data["NCI_COMMIT_REF_RELEASE"] = ""
	data["NCI_COMMIT_SHA"] = ""
	data["NCI_COMMIT_SHA_SHORT"] = ""
	data["NCI_COMMIT_TITLE"] = ""
	data["NCI_COMMIT_DESCRIPTION"] = ""

	if IsGitRepository(dir) {
		return CollectGitRepositoryInformation(dir, data)
	}

	return data, nil
}

// GetVCSRepositoryType returns a short name of the repository type inside the directory (git, svn, unknown)
func GetVCSRepositoryType(dir string) string {
	if IsGitRepository(dir) {
		return "git"
	}

	return "unknown"
}

// GetReferenceByName returns a unique identifier for a refType (tag, branch) / reFName (branch name / tag name)
func GetReferenceByName(projectDirectory string, refType string, refName string) (string, error) {
	if IsGitRepository(projectDirectory) {
		return FindGitReference(refType, refName), nil
	}

	return "", errors.New(projectDirectory + " is not a supported repository type!")
}

// FindCommitsBetweenRefs finds all git commits between two references
func FindCommitsBetweenRefs(dir string, startRef string, endRef string) ([]Commit, error) {
	if IsGitRepository(dir) {
		return FindGitCommitsBetweenRefs(dir, startRef, endRef)
	}

	return nil, errors.New(dir + " is not a supported repository type!")
}

// FindLatestRelease will retrieve the latest release version in the current branch
func FindLatestRelease(dir string, startAtReference string, stable bool, skipFrom bool) (Release, error) {
	if IsGitRepository(dir) {
		return FindLatestGitRelease(dir, startAtReference, stable, skipFrom)
	}

	return Release{}, errors.New(dir + " is not a supported repository type!")
}
