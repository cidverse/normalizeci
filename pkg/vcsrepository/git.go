package vcsrepository

import (
	"errors"
	"fmt"
	"github.com/Masterminds/semver/v3"
	"github.com/cidverse/normalizeci/pkg/ncispec"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/gosimple/slug"
	"github.com/rs/zerolog/log"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

// IsGitRepository checks if the target directory contains a git repository
func IsGitRepository(directory string) bool {
	if _, err := os.Stat(path.Join(directory, ".git")); !os.IsNotExist(err) {
		return true
	}

	return false
}

// CollectGitRepositoryInformation retrieves normalized variables of the git repository
func CollectGitRepositoryInformation(dir string, data map[string]string) (map[string]string, error) {
	// open repository from local path
	log.Trace().Msg("Using git repository at " + dir)
	repository, repositoryErr := git.PlainOpen(dir)
	if repositoryErr != nil {
		return nil, errors.New("failed to open git repository at " + dir)
	}
	isShallowClone := fileExists(filepath.Join(dir, ".git", "shallow"))

	// head reference
	ref, refErr := repository.Head()
	if refErr != nil {
		return nil, errors.New("failed to read git repository head, repository might not be initialized")
	}
	log.Trace().Str("ref", ref.String()).Msg("repository head lookup")

	// repository kind and remote
	data["NCI_REPOSITORY_KIND"] = "git"
	remote, remoteErr := repository.Remote("origin")
	if remoteErr == nil && remote != nil && remote.Config() != nil && len(remote.Config().URLs) > 0 {
		log.Trace().Str("remote", remote.String()).Msg("git remote lookup")
		data["NCI_REPOSITORY_REMOTE"] = remote.Config().URLs[0]
	} else {
		data["NCI_REPOSITORY_REMOTE"] = "local"
	}

	// pass
	if strings.HasPrefix(ref.Name().String(), "refs/heads/") {
		// branch
		branchName := ref.Name().String()[11:]
		data["NCI_COMMIT_REF_TYPE"] = "branch"
		data["NCI_COMMIT_REF_NAME"] = branchName
		data["NCI_COMMIT_REF_SLUG"] = slug.Make(branchName)
	} else if ref.Name().String() == "HEAD" {
		// detached HEAD, look into  the reflog to determinate the true branch
		gitRefLogFile := filepath.Join(dir, ".git", "logs", "HEAD")
		lastLine := readLastLine(gitRefLogFile)
		log.Debug().Msg("RefLog - LastLine: " + lastLine)

		pattern := regexp.MustCompile(`.*checkout: moving from (?P<FROM>.*) to (?P<TO>.*)$`)
		match := pattern.FindStringSubmatch(lastLine)
		log.Debug().Msg("Found a reflog entry showing that there was a checkout based on " + match[1] + " to " + match[2])

		if len(match[2]) == 40 {
			// checkout out a specific commit, use origin branch as reference
			data["NCI_COMMIT_REF_TYPE"] = "branch"
			data["NCI_COMMIT_REF_NAME"] = match[1]
			data["NCI_COMMIT_REF_SLUG"] = slug.Make(match[1])
		} else {
			// checkout of a tag or other named reference
			data["NCI_COMMIT_REF_TYPE"] = "tag"
			data["NCI_COMMIT_REF_NAME"] = match[2]
			data["NCI_COMMIT_REF_SLUG"] = slug.Make(match[2])
		}
	} else {
		panic("Can't determinate git ref, unsupported type!")
	}

	// reference path
	data["NCI_COMMIT_REF_PATH"] = data["NCI_COMMIT_REF_TYPE"] + "/" + data["NCI_COMMIT_REF_NAME"]

	// vcs specific reference
	currentRef, currentRefErr := GetReferenceByName(dir, data["NCI_COMMIT_REF_TYPE"], data["NCI_COMMIT_REF_NAME"])
	if currentRefErr != nil {
		return nil, errors.New("can't find repository reference for " + data["NCI_COMMIT_REF_TYPE"] + " - " + data["NCI_COMMIT_REF_NAME"])
	}
	data["NCI_COMMIT_REF_VCS"] = currentRef

	// release name (=name, but without leading v, without slash)
	data[ncispec.NCI_COMMIT_REF_RELEASE] = getReleaseName(data[ncispec.NCI_COMMIT_REF_NAME])

	// commit info
	data[ncispec.NCI_COMMIT_SHA] = ref.Hash().String()
	data[ncispec.NCI_COMMIT_SHA_SHORT] = ref.Hash().String()[0:8]

	cIter, _ := repository.Log(&git.LogOptions{From: ref.Hash()})
	firstCommit := true
	commitCount := 0
	cIter.ForEach(func(commit *object.Commit) error {
		commitInfo := strings.Split(commit.Message, "\n")
		commitCount++

		// only set for first commit
		if firstCommit {
			data[ncispec.NCI_COMMIT_TITLE] = commitInfo[0]

			if len(commitInfo) >= 3 {
				data[ncispec.NCI_COMMIT_DESCRIPTION] = strings.Join(commitInfo[2:], "\n")
			} else {
				data[ncispec.NCI_COMMIT_DESCRIPTION] = ""
			}

			data[ncispec.NCI_COMMIT_AUTHOR_NAME] = commit.Author.Name
			data[ncispec.NCI_COMMIT_AUTHOR_EMAIL] = commit.Author.Email
			data[ncispec.NCI_COMMIT_COMMITTER_NAME] = commit.Committer.Name
			data[ncispec.NCI_COMMIT_COMMITTER_EMAIL] = commit.Committer.Email

			firstCommit = false
		}

		return nil
	})

	// commit count
	if !isShallowClone {
		// can only be set, if the clone isn't shallow
		data[ncispec.NCI_COMMIT_COUNT] = strconv.Itoa(commitCount)
	}

	// previous release
	isStableRelease := false
	if data[ncispec.NCI_COMMIT_REF_TYPE] == "tag" {
		isStableRelease = isVersionStable(data["NCI_COMMIT_REF_NAME"])
	}
	previousRelease, previousReleaseErr := FindLatestRelease(dir, currentRef, isStableRelease, true)
	if previousReleaseErr == nil {
		data["NCI_LASTRELEASE_REF_NAME"] = previousRelease.Name
		data["NCI_LASTRELEASE_REF_SLUG"] = slug.Make(previousRelease.Name)
		data["NCI_LASTRELEASE_REF_VCS"] = previousRelease.Reference

		commits, commitsErr := FindCommitsBetweenRefs(dir, currentRef, previousRelease.Reference)
		if commitsErr == nil {
			data["NCI_LASTRELEASE_COMMIT_AFTER_COUNT"] = strconv.Itoa(len(commits))
		}
	}

	return data, nil
}

// FindGitReference returns the unique reference
func FindGitReference(refType string, refName string) string {
	// reference
	var gitReference string
	if refType == "branch" {
		gitReference = `refs/heads/` + refName
	} else if refType == "tag" {
		gitReference = `refs/tags/` + refName
	}

	return gitReference
}

// FindGitCommitsBetweenRefs finds all git commits between two references, get references via FindGitReference
func FindGitCommitsBetweenRefs(projectDir string, from string, to string) ([]Commit, error) {
	log.Trace().Str("projectDirectory", projectDir).Str("from", from).Str("to", to).Msg("retrieving commits between 2 references from git repository")

	// open repository from local path
	repository, repositoryErr := git.PlainOpen(projectDir)
	if repositoryErr != nil {
		return nil, errors.New("failed to open git repository")
	}

	// git references
	// - from
	var fromHash plumbing.Hash
	if strings.HasPrefix(from, "refs/") {
		fromRef, fromRefErr := repository.Reference(plumbing.ReferenceName(from), true)
		if fromRefErr != nil {
			return nil, errors.New("can't resolve from reference")
		}
		fromHash = fromRef.Hash()
	} else {
		fromCommit, fromCommitErr := repository.CommitObject(plumbing.NewHash(from))
		if fromCommitErr != nil {
			return nil, errors.New("can't resolve commit hash [from] for " + from + ": " + fromCommitErr.Error())
		}
		fromHash = fromCommit.Hash
	}
	// - to
	var toHash plumbing.Hash
	if strings.HasPrefix(to, "refs/") {
		toRef, toRefErr := repository.Reference(plumbing.ReferenceName(to), true)
		if toRefErr != nil {
			return nil, errors.New("can't resolve to reference")
		}
		toHash = toRef.Hash()
	} else {
		toCommit, toCommitErr := repository.CommitObject(plumbing.NewHash(to))
		if toCommitErr != nil {
			return nil, errors.New("can't resolve commit hash [to] for " + to + ": " + toCommitErr.Error())
		}
		toHash = toCommit.Hash
	}

	// commit references
	var commitRefs = make(map[plumbing.Hash][]*plumbing.Reference)
	refIterator, refIteratorErr := repository.Tags()
	if refIteratorErr == nil {
		refIterator.ForEach(func(t *plumbing.Reference) error {
			commitRefs[t.Hash()] = append(commitRefs[t.Hash()], t)
			return nil
		})
	}

	// commit iterator
	cIter, _ := repository.Log(&git.LogOptions{From: fromHash})
	var commits []Commit
	for {
		commit, commitErr := cIter.Next()
		if commitErr != nil {
			break
		}

		// log
		log.Debug().Str("commit-hash", commit.Hash.String()).Str("commit-subject", commit.Message).Msg("checking commit")

		// check
		if commit.Hash.String() == toHash.String() {
			break
		}

		// refs?
		var commitTags []CommitTag
		if refs, hasRefs := commitRefs[commit.Hash]; hasRefs {
			for _, ref := range refs {
				commitTags = append(commitTags, CommitTag{Name: strings.TrimLeft(ref.Name().String(), "refs/tags/"), VCSRef: ref.Name().String()})
			}
		}

		commitInfo := strings.SplitN(commit.Message, "\n", 2)
		var commitDescription string
		if len(commitInfo) == 2 {
			commitDescription = commitInfo[1]
		}
		commits = append(commits, Commit{ShortHash: commit.Hash.String()[:7], Hash: commit.Hash.String(), Message: commitInfo[0], Description: commitDescription, Author: CommitAuthor{Name: commit.Author.Name, Email: commit.Author.Email}, Committer: CommitAuthor{Name: commit.Committer.Name, Email: commit.Committer.Email}, AuthoredAt: commit.Author.When, CommittedAt: commit.Committer.When, Tags: commitTags})
	}

	return commits, nil
}

func FindLatestGitRelease(projectDir string, from string, stable bool, skipFrom bool) (Release, error) {
	// open repository from local path
	log.Trace().Str("projectDirectory", projectDir).Msg("opening git repository")
	repository, repositoryErr := git.PlainOpen(projectDir)
	if repositoryErr != nil {
		return Release{}, errors.New("failed to open git repository")
	}

	// references
	fromRef, fromRefErr := repository.Reference(plumbing.ReferenceName(from), true)
	if fromRefErr != nil {
		return Release{}, errors.New("can't resolve from reference: " + from)
	}

	// commit references
	var commitRefs = make(map[plumbing.Hash][]*plumbing.Reference)
	refIterator, refIteratorErr := repository.Tags()
	if refIteratorErr == nil {
		refIterator.ForEach(func(t *plumbing.Reference) error {
			commitRefs[t.Hash()] = append(commitRefs[t.Hash()], t)
			return nil
		})
	}

	// commit iterator
	cIter, _ := repository.Log(&git.LogOptions{From: fromRef.Hash()})
	var lastCommit *object.Commit
	commitCount := 0
	for {
		commit, commitErr := cIter.Next()
		if commitErr != nil {
			break
		}
		lastCommit = commit
		commitCount++

		// log
		log.Debug().Str("commit-hash", commit.Hash.String()).Str("commit-subject", commit.Message).Str("commit-refs", fmt.Sprintf("%+v\n", commitRefs[commit.Hash])).Msg("checking commit")

		// skip first?
		if commitCount == 1 && skipFrom {
			continue
		}

		// refs
		if refs, hasRefs := commitRefs[commit.Hash]; hasRefs {
			for _, ref := range refs {
				if ref.Name().IsTag() {
					version, versionErr := semver.NewVersion(ref.Name().Short())
					if versionErr == nil {
						if stable && isVersionStable(ref.Name().Short()) {
							return Release{Name: version.String(), Reference: ref.Name().String()}, nil
						} else if !stable {
							return Release{Name: version.String(), Reference: ref.Name().String()}, nil
						}
					}
				}
			}
		}
	}

	if lastCommit == nil {
		return Release{}, errors.New("repository does not contain any commits")
	}
	return Release{Name: "0.0.0", Reference: lastCommit.Hash.String()}, nil
}
