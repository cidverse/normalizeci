package gitclient

import (
	"errors"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/Masterminds/semver/v3"
	"github.com/cidverse/normalizeci/pkg/common"
	"github.com/cidverse/normalizeci/pkg/vcsrepository/vcsapi"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/go-git/go-git/v5/utils/merkletrie"
	"github.com/rs/zerolog/log"
)

type GitClient struct {
	dir        string
	repo       *git.Repository
	isShallow  bool
	tags       []vcsapi.VCSRef
	tagsByHash map[string][]vcsapi.VCSRef
}

func NewGitClient(dir string) (vcsapi.Client, error) {
	c := GitClient{dir: dir}
	if !c.Check() {
		return nil, errors.New("is not a git repository")
	}

	repo, err := git.PlainOpen(dir)
	if err != nil {
		return nil, errors.New("failed to open git repository at " + dir + ": " + err.Error())
	}
	c.repo = repo
	c.isShallow = fileExists(filepath.Join(dir, ".git", "shallow"))

	return c, nil
}

func (c GitClient) Check() bool {
	if _, err := os.Stat(path.Join(c.dir, ".git")); !os.IsNotExist(err) {
		return true
	}

	return false
}

func (c GitClient) VCSType() string {
	return "git"
}

func (c GitClient) VCSRemote() string {
	remote, remoteErr := c.repo.Remote("origin")
	if remoteErr == nil && remote != nil && remote.Config() != nil && len(remote.Config().URLs) > 0 {
		return remote.Config().URLs[0]
	}

	return "local"
}

func (c GitClient) VCSHostServer(remote string) string {
	if remote != "local" {
		// git over ssh
		if strings.HasPrefix(remote, "git@") {
			re := regexp.MustCompile(`(?i)^git@([^:]+):`)
			host := re.FindStringSubmatch(remote)[1]
			return host
		}

		u, err := url.Parse(remote)
		if err != nil {
			log.Warn().Err(err).Msg("error parsing URL")
			return ""
		}

		return u.Host
	}

	return ""
}

func (c GitClient) VCSHostType(server string) string {
	if server == "github.com" {
		return "github"
	} else if server == "gitlab.com" || strings.Contains(server, "gitlab.") {
		return "gitlab"
	} else if len(os.Getenv(common.ToEnvName(server)+"_TYPE")) > 0 {
		os.Getenv(common.ToEnvName(server) + "_TYPE")
	}

	return ""
}

func (c GitClient) VCSRefToInternalRef(ref vcsapi.VCSRef) string {
	if ref.Type == "branch" {
		return `refs/heads/` + ref.Value
	} else if ref.Type == "tag" {
		return `refs/tags/` + strings.TrimPrefix(ref.Value, "tags/")
	}

	return ref.Hash
}

func (c GitClient) VCSHead() (vcsHead vcsapi.VCSRef, err error) {
	// head reference
	ref, err := c.repo.Head()
	if err != nil {
		return vcsapi.VCSRef{}, err
	}

	if strings.HasPrefix(ref.Name().String(), "refs/heads/") {
		branchName := ref.Name().String()[11:]
		return vcsapi.VCSRef{Type: "branch", Value: branchName, Hash: ref.Hash().String()}, nil
	} else if strings.HasPrefix(ref.Name().String(), "refs/tags/") {
		tagName := ref.Name().String()[10:]
		return vcsapi.VCSRef{Type: "tag", Value: tagName, Hash: ref.Hash().String()}, nil
	} else if ref.Name().String() == "HEAD" {
		// detached HEAD, check git reflog for the true reference
		gitRefLogFile := filepath.Join(c.dir, ".git", "logs", "HEAD")
		lastLine := readLastLine(gitRefLogFile)

		pattern := regexp.MustCompile(`.*checkout: moving from (?P<FROM>.*) to (?P<TO>.*)$`)
		match := pattern.FindStringSubmatch(lastLine)

		if strings.HasPrefix(match[2], "refs/remotes/pull") {
			// handle github merge request as virtual branch
			return vcsapi.VCSRef{Type: "branch", Value: match[2][13:], Hash: ref.Hash().String()}, nil
		} else if len(match[2]) == 40 {
			return vcsapi.VCSRef{Type: "branch", Value: match[1], Hash: ref.Hash().String()}, nil
		} else {
			return vcsapi.VCSRef{Type: "tag", Value: match[2], Hash: ref.Hash().String()}, nil
		}
	}

	return vcsapi.VCSRef{}, errors.New("can't determinate repo head")
}

func (c GitClient) GetTags() []vcsapi.VCSRef {
	if c.tags == nil {
		var tags []vcsapi.VCSRef

		iter, err := c.repo.Tags()
		if err == nil {
			iter.ForEach(func(r *plumbing.Reference) error {
				if r.Name().IsTag() {
					t := vcsapi.VCSRef{
						Type:  "tag",
						Value: strings.TrimPrefix(r.Name().String(), "refs/tags/"),
						Hash:  r.Hash().String(),
					}
					tags = append(tags, t)
				}

				return nil
			})
		}

		c.tags = tags
	}

	return c.tags
}

func (c GitClient) GetTagsByHash(hash string) []vcsapi.VCSRef {
	if c.tagsByHash == nil {
		c.tagsByHash = make(map[string][]vcsapi.VCSRef)

		iter, err := c.repo.Tags()
		if err == nil {
			iter.ForEach(func(r *plumbing.Reference) error {
				if r.Name().IsTag() {
					t := vcsapi.VCSRef{
						Type:  "tag",
						Value: strings.TrimPrefix(r.Name().String(), "refs/tags/"),
						Hash:  r.Hash().String(),
					}

					c.tagsByHash[r.Hash().String()] = append(c.tagsByHash[r.Hash().String()], t)
				}

				return nil
			})
		}
	}

	return c.tagsByHash[hash]
}

func (c GitClient) FindCommitByHash(hash string, includeChanges bool) (vcsapi.Commit, error) {
	cIter, _ := c.repo.Log(&git.LogOptions{All: true})
	for {
		commit, commitErr := cIter.Next()
		if commitErr != nil {
			return vcsapi.Commit{}, commitErr
		}

		// check
		if hash == commit.Hash.String() || hash == commit.Hash.String()[:7] {
			// return commit
			return gitCommitToVCSCommit(commit, c.GetTagsByHash(commit.Hash.String()), includeChanges), nil
		}
	}
}

func (c GitClient) FindCommitsBetween(from *vcsapi.VCSRef, to *vcsapi.VCSRef, includeChanges bool, limit int) ([]vcsapi.Commit, error) {
	// from reference
	var fromHash plumbing.Hash
	if from == nil {
		head, headErr := c.repo.Head()
		if headErr != nil {
			return nil, headErr
		}
		fromHash = head.Hash()
	} else {
		fromHash, _ = refToHash(c.repo, c.VCSRefToInternalRef(*from))
	}

	// to reference
	var toHash plumbing.Hash
	if to != nil {
		toHash, _ = refToHash(c.repo, c.VCSRefToInternalRef(*to))
	}

	// commit iterator
	cIter, _ := c.repo.Log(&git.LogOptions{From: fromHash})
	var commits []vcsapi.Commit
	for {
		commit, commitErr := cIter.Next()
		if commitErr != nil {
			break
		}

		// check
		if to != nil && toHash.String() == commit.Hash.String() {
			break
		}

		// limit
		if limit != 0 && len(commits) >= limit {
			break
		}

		commits = append(commits, gitCommitToVCSCommit(commit, c.GetTagsByHash(commit.Hash.String()), includeChanges))
	}

	return commits, nil
}

func (c GitClient) FindLatestRelease(stable bool) (vcsapi.VCSRelease, error) {
	var latestVersion, _ = semver.NewVersion("0.0.0")
	var latest vcsapi.VCSRelease

	tags := c.GetTags()
	for _, tag := range tags {
		version, versionErr := semver.NewVersion(tag.Value)
		if versionErr == nil {
			if version.Compare(latestVersion) > 0 {
				if (stable && len(version.Prerelease()) == 0) || !stable {
					latestVersion = version
					latest = vcsapi.VCSRelease{
						Type:    tag.Type,
						Value:   tag.Value,
						Hash:    tag.Hash,
						Version: version.String(),
					}
				}
			}
		}
	}

	return latest, nil
}

func gitFileActionToText(input merkletrie.Action) string {
	if input == merkletrie.Insert {
		return "create"
	} else if input == merkletrie.Modify {
		return "update"
	} else if input == merkletrie.Delete {
		return "delete"
	}

	return ""
}

func refToHash(repo *git.Repository, ref string) (hash plumbing.Hash, err error) {
	if strings.HasPrefix(ref, "refs/") {
		var pRef *plumbing.Reference
		pRef, err = repo.Reference(plumbing.ReferenceName(ref), true)
		if err != nil {
			return hash, err
		}
		hash = pRef.Hash()
	} else {
		var commit *object.Commit
		commit, err = repo.CommitObject(plumbing.NewHash(ref))
		if err != nil {
			return hash, err
		}
		hash = commit.Hash
	}

	return hash, err
}
