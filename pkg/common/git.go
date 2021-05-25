package common

import (
	"bufio"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/rs/zerolog/log"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

// GetGitDirectory finds the first git directory from the current working directory upwards
func GetGitDirectory() string {
	currentDirectory, _ := os.Getwd()
	var projectDirectory = ""
	directoryParts := strings.Split(currentDirectory, string(os.PathSeparator))

	for projectDirectory == "" {
		if _, err := os.Stat(filepath.Join(currentDirectory, ".git")); err == nil {
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
	log.Debug().Msg("Using git repository at " + projectDir)
	repository, repositoryErr := git.PlainOpen(projectDir)
	if repositoryErr != nil {
		info = append(info, "NCI_REPOSITORY_KIND=none")
		info = append(info, "NCI_REPOSITORY_REMOTE=local")
		info = append(info, "NCI_COMMIT_REF_TYPE=unknown")
		info = append(info, "NCI_COMMIT_REF_NAME=unknown")
		info = append(info, "NCI_COMMIT_REF_SLUG=unknown")
		info = append(info, "NCI_COMMIT_REF_RELEASE=")
		info = append(info, "NCI_COMMIT_SHA=")
		info = append(info, "NCI_COMMIT_SHA_SHORT=")
		info = append(info, "NCI_COMMIT_TITLE=")
		info = append(info, "NCI_COMMIT_DESCRIPTION=")
		return info
	}
	isShallowClone := FileExists(filepath.Join(projectDir, ".git", "shallow "))

	ref, refErr := repository.Head()
	if refErr != nil {
		info = append(info, "NCI_REPOSITORY_KIND=none")
		info = append(info, "NCI_REPOSITORY_REMOTE=local")
		info = append(info, "NCI_COMMIT_REF_TYPE=unknown")
		info = append(info, "NCI_COMMIT_REF_NAME=unknown")
		info = append(info, "NCI_COMMIT_REF_SLUG=unknown")
		info = append(info, "NCI_COMMIT_REF_RELEASE=")
		info = append(info, "NCI_COMMIT_SHA=")
		info = append(info, "NCI_COMMIT_SHA_SHORT=")
		info = append(info, "NCI_COMMIT_TITLE=")
		info = append(info, "NCI_COMMIT_DESCRIPTION=")
		return info
	}
	log.Debug().Msg("Git Ref " + ref.String())

	// repository kind and remote
	info = append(info, "NCI_REPOSITORY_KIND=git")
	remote, remoteErr := repository.Remote("origin")
	if remoteErr == nil && remote != nil && remote.Config() != nil && len(remote.Config().URLs) > 0 {
		log.Debug().Msg("Git Remote " + remote.String())
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
		log.Debug().Msg("RefLog - LastLine: " + lastLine)

		pattern := regexp.MustCompile(`.*checkout: moving from (?P<FROM>.*) to (?P<TO>.*)$`)
		match := pattern.FindStringSubmatch(lastLine)
		log.Debug().Msg("Found a reflog entry showing that there was a checkout based on " + match[1] + " to " + match[2])

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
	commitCount := 0
	cIter.ForEach(func(commit *object.Commit) error {
		commitInfo := strings.Split(commit.Message, "\n")
		commitCount++

		// only set for first commit
		if firstCommit {
			info = append(info, "NCI_COMMIT_TITLE="+commitInfo[0])
			if len(commitInfo) >= 3 {
				info = append(info, "NCI_COMMIT_DESCRIPTION="+strings.Join(commitInfo[2:], "\n"))
			} else {
				info = append(info, "NCI_COMMIT_DESCRIPTION=")
			}

			firstCommit = false
		}

		return nil
	})

	// commit count
	if !isShallowClone {
		// can only be set, if the clone isn't shallow
		info = append(info, "NCI_COMMIT_COUNT=" + strconv.Itoa(commitCount))
	}


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
