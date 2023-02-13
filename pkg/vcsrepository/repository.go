package vcsrepository

import (
	"errors"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/cidverse/normalizeci/pkg/ncispec"
	gitclient "github.com/cidverse/normalizeci/pkg/vcsrepository/git"
	"github.com/cidverse/normalizeci/pkg/vcsrepository/vcsapi"
	"github.com/gosimple/slug"
)

func GetVCSClient(dir string) (vcsapi.Client, error) {
	// git
	cg, _ := gitclient.NewGitClient(dir)
	if cg.Check() {
		return cg, nil
	}

	return nil, errors.New("directory is not a vcs repository")
}

func FindRepositoryDirectory(currentDirectory string) string {
	var projectDirectory = ""
	directoryParts := strings.Split(currentDirectory, string(os.PathSeparator))

	for projectDirectory == "" {
		// GIT

		if _, err := os.Stat(path.Join(currentDirectory, ".git")); !os.IsNotExist(err) {
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
func GetVCSRepositoryInformation(dir string) (data map[string]string, err error) {
	// init with default values
	data = make(map[string]string)
	data[ncispec.NCI_REPOSITORY_KIND] = "none"
	data[ncispec.NCI_REPOSITORY_REMOTE] = "local"
	data[ncispec.NCI_REPOSITORY_HOST_SERVER] = ""
	data[ncispec.NCI_REPOSITORY_HOST_TYPE] = ""
	data[ncispec.NCI_COMMIT_REF_TYPE] = "unknown"
	data[ncispec.NCI_COMMIT_REF_NAME] = "unknown"
	data[ncispec.NCI_COMMIT_REF_SLUG] = ""
	data[ncispec.NCI_COMMIT_REF_PATH] = ""
	data[ncispec.NCI_COMMIT_REF_VCS] = ""
	data[ncispec.NCI_COMMIT_REF_RELEASE] = ""
	data[ncispec.NCI_COMMIT_SHA] = ""
	data[ncispec.NCI_COMMIT_SHA_SHORT] = ""
	data[ncispec.NCI_COMMIT_TITLE] = ""
	data[ncispec.NCI_COMMIT_DESCRIPTION] = ""
	data[ncispec.NCI_REPOSITORY_STATUS] = "clean"

	// supported repository type
	client, clientErr := GetVCSClient(dir)
	if client == nil {
		return data, clientErr
	}

	// repository type and remote
	data[ncispec.NCI_REPOSITORY_KIND] = client.VCSType()
	data[ncispec.NCI_REPOSITORY_REMOTE] = client.VCSRemote()
	data[ncispec.NCI_REPOSITORY_HOST_SERVER] = client.VCSHostServer(data[ncispec.NCI_REPOSITORY_REMOTE])
	data[ncispec.NCI_REPOSITORY_HOST_TYPE] = client.VCSHostType(data[ncispec.NCI_REPOSITORY_HOST_SERVER])

	// repository head
	head, err := client.VCSHead()
	if err != nil {
		return data, err
	}
	refName := head.Value
	if head.Type == "tag" {
		refName = strings.TrimPrefix(strings.TrimPrefix(refName, "tags/"), "refs/tags/")
	}
	data[ncispec.NCI_COMMIT_REF_TYPE] = head.Type
	data[ncispec.NCI_COMMIT_REF_NAME] = refName
	data[ncispec.NCI_COMMIT_REF_SLUG] = slug.Make(refName)
	data[ncispec.NCI_COMMIT_REF_PATH] = head.Type + "/" + refName
	data[ncispec.NCI_COMMIT_REF_VCS] = client.VCSRefToInternalRef(head)

	// release name (=name, but without leading v, without slash)
	data[ncispec.NCI_COMMIT_REF_RELEASE] = getReleaseName(data[ncispec.NCI_COMMIT_REF_NAME])

	// repository status (data[ncispec.NCI_REPOSITORY_STATUS])
	// TODO: current isClean by go-git detects newlines as change, see https://github.com/go-git/go-git/issues/436
	/*
		workTree, workTreeErr := repository.Worktree()
		if workTreeErr == nil {
			workTreeStatus, workTreeStatusErr := workTree.Status()
			if workTreeStatusErr == nil {
				for file, fileStatus := range workTreeStatus {
					if ignoreMatcher.MatchesPath(file) {
						continue
					}

					// check for "dirty" files in the local repository
					if fileStatus.Worktree != git.Unmodified || fileStatus.Staging != git.Unmodified {
						// data[ncispec.NCI_REPOSITORY_STATUS] = "dirty"
						// break
					}
				}
			}
		}
	*/

	// commit info
	commit, err := client.FindCommitByHash(head.Hash, false)
	if err != nil {
		return data, err
	}
	data[ncispec.NCI_COMMIT_SHA] = commit.Hash
	data[ncispec.NCI_COMMIT_SHA_SHORT] = commit.ShortHash
	data[ncispec.NCI_COMMIT_TITLE] = commit.Message
	data[ncispec.NCI_COMMIT_DESCRIPTION] = commit.Description
	data[ncispec.NCI_COMMIT_AUTHOR_NAME] = commit.Author.Name
	data[ncispec.NCI_COMMIT_AUTHOR_EMAIL] = commit.Author.Email
	data[ncispec.NCI_COMMIT_COMMITTER_NAME] = commit.Committer.Name
	data[ncispec.NCI_COMMIT_COMMITTER_EMAIL] = commit.Committer.Email
	data[ncispec.NCI_COMMIT_COUNT] = strconv.Itoa(0)

	// commit count (only if clone is not shallow)
	if true == false {
		// can only be set, if the clone isn't shallow
		data[ncispec.NCI_COMMIT_COUNT] = strconv.Itoa(50)
	}

	latest, latestErr := client.FindLatestRelease(false)
	if latestErr == nil {
		data[ncispec.NCI_LASTRELEASE_REF_NAME] = latest.Version
		data[ncispec.NCI_LASTRELEASE_REF_SLUG] = slug.Make(latest.Version)
		data[ncispec.NCI_LASTRELEASE_REF_VCS] = client.VCSRefToInternalRef(vcsapi.VCSRef{Type: latest.Type, Value: latest.Value, Hash: latest.Hash})
		data[ncispec.NCI_LASTRELEASE_COMMIT_AFTER_COUNT] = "0"

		commits, commitsErr := client.FindCommitsBetween(&head, &vcsapi.VCSRef{Type: latest.Type, Value: latest.Value, Hash: latest.Hash}, false, 0)
		if commitsErr == nil {
			data[ncispec.NCI_LASTRELEASE_COMMIT_AFTER_COUNT] = strconv.Itoa(len(commits))
		}
	}

	return data, nil
}
