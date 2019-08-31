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
	repository, _ := git.PlainOpen(projectDir)
	remote, remoteErr := repository.Remote("origin")
	ref, _ := repository.Head()

	// repository kind and remote
	info = append(info, "NCI_REPOSITORY_KIND=git")
	if remoteErr == nil {
		info = append(info, "NCI_REPOSITORY_REMOTE="+remote.Config().URLs[0])
	} else {
		info = append(info, "NCI_REPOSITORY_REMOTE=local")
	}

	// pass
	if strings.HasPrefix(ref.Name().String(), "refs/heads/") {
		// branch
		info = append(info, "NCI_COMMIT_REF_TYPE=branch")
		info = append(info, "NCI_COMMIT_REF_NAME="+strings.TrimLeft(ref.Name().String(), "refs/heads/"))
		info = append(info, "NCI_COMMIT_REF_SLUG="+GetSlug(strings.TrimLeft(ref.Name().String(), "refs/heads/")))
	} else if ref.Name().String() == "HEAD" {
		// detached HEAD, look into  the reflog to determinate the true branch
		gitRefLogFile := projectDir + "/.git/logs/HEAD"
		lastLine := readLastLine(gitRefLogFile)
		log.Debug(lastLine)

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
			info = append(info, "NCI_COMMIT_DESCRIPTION="+strings.Join(commitinfo[2:], "\n"))

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
