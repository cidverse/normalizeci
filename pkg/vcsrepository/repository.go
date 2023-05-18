package vcsrepository

import (
	"errors"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"strings"

	v1 "github.com/cidverse/normalizeci/pkg/ncispec/v1"
	gitclient "github.com/cidverse/normalizeci/pkg/vcsrepository/git"
	"github.com/cidverse/normalizeci/pkg/vcsrepository/vcsapi"
	"github.com/gosimple/slug"
)

var MockClient vcsapi.Client

func GetVCSClient(dir string) (vcsapi.Client, error) {
	// mocked client
	if MockClient != nil {
		return MockClient, nil
	}

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

type RepositoryInformation struct {
	Repository v1.Repository
	Commit     v1.Commit
}

// GetVCSRepositoryInformation detects the repository type and gathers normalized information from the repository
func GetVCSRepositoryInformation(dir string) (RepositoryInformation, error) {
	// init with default values
	result := RepositoryInformation{
		Repository: v1.Repository{
			Kind:   "none",
			Remote: "local",
			Status: "clean",
		},
		Commit: v1.Commit{
			RefType: "unknown",
			RefName: "unknown",
		},
	}

	// supported repository type
	client, clientErr := GetVCSClient(dir)
	if client == nil {
		return result, clientErr
	}

	// repository type and remote
	result.Repository.Kind = client.VCSType()
	result.Repository.Remote = client.VCSRemote()
	result.Repository.HostServer = client.VCSHostServer(result.Repository.Remote)
	result.Repository.HostType = client.VCSHostType(result.Repository.HostServer)

	// repository head
	head, err := client.VCSHead()
	if err != nil {
		return result, err
	}
	refName := head.Value
	if head.Type == "tag" {
		refName = strings.TrimPrefix(strings.TrimPrefix(refName, "tags/"), "refs/tags/")
	}
	result.Commit.RefType = head.Type
	result.Commit.RefName = refName
	result.Commit.RefSlug = slug.Make(refName)
	result.Commit.RefPath = head.Type + "/" + refName
	result.Commit.RefVCS = client.VCSRefToInternalRef(head)

	// release name (=name, but without leading v, without slash)
	result.Commit.RefRelease = getReleaseName(result.Commit.RefName)

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
		return result, err
	}

	result.Commit.Hash = commit.Hash
	result.Commit.HashShort = commit.ShortHash
	result.Commit.Title = commit.Message
	result.Commit.Description = commit.Description
	result.Commit.AuthorName = commit.Author.Name
	result.Commit.AuthorEmail = commit.Author.Email
	result.Commit.CommitterName = commit.Committer.Name
	result.Commit.CommitterEmail = commit.Committer.Email
	result.Commit.Count = strconv.Itoa(0)

	// TODO: commit count (only if clone is not shallow)

	return result, nil
}
