package common

import (
	"bufio"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	log "github.com/sirupsen/logrus"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing/object"
)

// GetGitDirectory finds the first git directory from the current working directory upwards
func GetGitDirectory() string {
	currentDirectory, _ := os.Getwd()
	var projectDirectory = ""
	directoryParts := strings.Split(currentDirectory, string(os.PathSeparator))

	for projectDirectory == "" {
		if _, err := os.Stat(filepath.Join(currentDirectory, "/.git")); err == nil {
			return currentDirectory
		}

		if directoryParts[0]+"\\" == currentDirectory || currentDirectory == "/" {
			return ""
		}

		currentDirectory = filepath.Dir(currentDirectory)
	}

	return ""
}

// GetSCMArguments gets the common git values
func GetSCMArguments(projectDir string) []string {
	var info []string

	// open repository from local path
	log.Debug("Using git repository at " + projectDir)
	repository, repositoryErr := git.PlainOpen(projectDir)
	if repositoryErr != nil {
		log.Warn("No repository!")
		return getDefaultInfo()
	}

	// get current reference
	ref, refErr := repository.Head()
	if refErr != nil {
		log.Warn("Empty repository!")
		return getDefaultInfo()
	}
	log.Debug("Git Ref " + ref.String())

	// repository kind and remote
	info = append(info, "NCI_REPOSITORY_KIND=git")
	remote, remoteErr := repository.Remote("origin")
	log.Debug("Git Remote " + remote.String())
	if remoteErr == nil && remote != nil && remote.Config() != nil && len(remote.Config().URLs) > 0 {
		info = append(info, "NCI_REPOSITORY_REMOTE="+remote.Config().URLs[0])
	} else {
		info = append(info, "NCI_REPOSITORY_REMOTE=local")
	}

	// pass
	if strings.HasPrefix(ref.Name().String(), "refs/heads/") {
		// branch
		branchName := ref.Name().String()[11:]
		info = append(info, "NCI_COMMIT_REF_TYPE=branch")
		info = append(info, "NCI_COMMIT_REF_NAME="+branchName)
		info = append(info, "NCI_COMMIT_REF_SLUG="+GetSlug(branchName))
	} else if ref.Name().String() == "HEAD" {
		// detached HEAD, look into  the reflog to determinate the true branch
		gitRefLogFile := projectDir + "/.git/logs/HEAD"
		lastLine := readLastLine(gitRefLogFile)
		log.Debug("RefLog - LastLine: " + lastLine)

		pattern := regexp.MustCompile(`.*checkout: moving from (?P<FROM>.*) to (?P<TO>.*)$`)
		match := pattern.FindStringSubmatch(lastLine)
		log.Debug("Found a reflog entry showing that there was a checkout based on " + match[1] + " to " + match[2])

		if len(match[2]) == 40 {
			// checkout out a specific commit, use origin branch as reference
			info = append(info, "NCI_COMMIT_REF_TYPE=branch")
			info = append(info, "NCI_COMMIT_REF_NAME="+match[1])
			info = append(info, "NCI_COMMIT_REF_SLUG="+GetSlug(match[1]))
		} else {
			// checkout of a tag or other named reference
			info = append(info, "NCI_COMMIT_REF_TYPE=tag")
			info = append(info, "NCI_COMMIT_REF_NAME="+match[2])
			info = append(info, "NCI_COMMIT_REF_SLUG="+GetSlug(match[2]))
		}
	} else {
		panic("Unsupported!")
	}

	// release name (=slug, but without leading v for tags)
	info = append(info, "NCI_COMMIT_REF_RELEASE="+strings.TrimLeft(GetEnvironment(info, "NCI_COMMIT_REF_SLUG"), "v"))

	// commit info
	info = append(info, "NCI_COMMIT_SHA="+ref.Hash().String())
	info = append(info, "NCI_COMMIT_SHA_SHORT="+ref.Hash().String()[0:8])

	cIter, _ := repository.Log(&git.LogOptions{From: ref.Hash()})
	firstCommit := true
	cIter.ForEach(func(commit *object.Commit) error {
		commitinfo := strings.Split(commit.Message, "\n")

		// only set for first commit
		if firstCommit {
			info = append(info, "NCI_COMMIT_TITLE="+commitinfo[0])
			if len(commitinfo) >= 3 {
				info = append(info, "NCI_COMMIT_DESCRIPTION="+strings.Join(commitinfo[2:], "\n"))
			} else {
				info = append(info, "NCI_COMMIT_DESCRIPTION=")
			}

			firstCommit = false
		}

		return nil
	})

	return info
}

// readLastLine gets the last line from a file
func readLastLine(filename string) string {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	defer file.Close()
	reader := bufio.NewReader(file)

	lastLine := ""
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		lastLine = string(line)
	}

	return lastLine
}

func getDefaultInfo() []string {
	var info []string

	info = append(info, "NCI_REPOSITORY_KIND=none")
	info = append(info, "NCI_REPOSITORY_REMOTE=local")
	info = append(info, "NCI_COMMIT_REF_TYPE=unknown")
	info = append(info, "NCI_COMMIT_REF_NAME=unknown")
	info = append(info, "NCI_COMMIT_REF_SLUG=unknown")
	info = append(info, "NCI_COMMIT_REF_RELEASE=unknown")
	info = append(info, "NCI_COMMIT_SHA=")
	info = append(info, "NCI_COMMIT_SHA_SHORT=")
	info = append(info, "NCI_COMMIT_TITLE=")
	info = append(info, "NCI_COMMIT_DESCRIPTION=")

	return info
}
