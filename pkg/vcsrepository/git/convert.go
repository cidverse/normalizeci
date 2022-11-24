package gitclient

import (
	"github.com/cidverse/normalizeci/pkg/vcsrepository/vcsapi"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/rs/zerolog/log"
	"strings"
)

func gitCommitToVCSCommit(commit *object.Commit, tags []vcsapi.VCSRef, withContent bool) vcsapi.Commit {
	// commit title and description
	commitInfo := strings.SplitN(commit.Message, "\n", 2)
	var commitDescription string
	if len(commitInfo) == 2 {
		commitDescription = commitInfo[1]
		commitDescription = strings.Trim(strings.Trim(commitDescription, "\r\n"), "\n")
	}

	// changes
	var changes []vcsapi.CommitChange
	if withContent {
		var err error
		changes, err = gitCommitChangeToVCSCommitChange(commit)
		log.Debug().Err(err).Str("hash", commit.Hash.String()).Msg("failed to get changes for commit")
	}

	return vcsapi.Commit{
		ShortHash:   commit.Hash.String()[:7],
		Hash:        commit.Hash.String(),
		Message:     commitInfo[0],
		Description: commitDescription,
		Author: vcsapi.CommitAuthor{
			Name:  commit.Author.Name,
			Email: commit.Author.Email,
		},
		Committer: vcsapi.CommitAuthor{
			Name:  commit.Committer.Name,
			Email: commit.Committer.Email,
		},
		Tags:        tags,
		AuthoredAt:  commit.Author.When,
		CommittedAt: commit.Committer.When,
		Context:     nil,
		Changes:     changes,
	}
}

func gitCommitChangeToVCSCommitChange(commit *object.Commit) (result []vcsapi.CommitChange, err error) {
	currentDirState, err := commit.Tree()
	if err != nil {
		return nil, err
	}

	// get Previous Directory state
	prevCommitObject, err := commit.Parents().Next()
	if err != nil {
		return nil, err
	}

	prevDirState, err := prevCommitObject.Tree()
	if err != nil {
		return nil, err
	}

	changes, err := prevDirState.Diff(currentDirState)
	if err != nil {
		return nil, err
	}

	for _, change := range changes {
		// Ignore deleted files
		action, cErr := change.Action()
		if cErr != nil {
			return nil, cErr
		}

		// file change
		from, to, err := change.Files()
		if err != nil {
			return nil, err
		}
		var fromFile vcsapi.CommitFile
		if from != nil {
			fromFile = vcsapi.CommitFile{
				Name: from.Name,
				Size: from.Size,
				Hash: from.Hash.String(),
			}
		}
		var toFile vcsapi.CommitFile
		if to != nil {
			toFile = vcsapi.CommitFile{
				Name: to.Name,
				Size: to.Size,
				Hash: to.Hash.String(),
			}
		}

		// content patch
		patch, err := change.Patch()
		if err != nil {
			return nil, err
		}

		// append to change list
		result = append(result, vcsapi.CommitChange{
			Type:     gitFileActionToText(action),
			FileFrom: fromFile,
			FileTo:   toFile,
			Patch:    patch.String(),
		})
	}

	return result, err
}
