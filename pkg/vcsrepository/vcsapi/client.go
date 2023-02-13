package vcsapi

import (
	"fmt"
	"strings"
)

// Client is the common interface for all vcs repo implementations
type Client interface {
	// Check will verify that the repository type is supported by this implementation
	Check() bool

	// VCSType returns the version control system type (git, ...)
	VCSType() string

	// VCSRemote returns the primary remote (server)
	VCSRemote() string

	// VCSHostServer returns the host of the primary remote
	VCSHostServer(remote string) string

	// VCSHostType returns the type of the host
	VCSHostType(server string) string

	// VCSRefToInternalRef converts the reference to the internal notation used by the VCS
	VCSRefToInternalRef(ref VCSRef) string

	// VCSHead returns the current head of the repository (refType and refName)
	VCSHead() (ref VCSRef, err error)

	// GetTags returns all tags
	GetTags() (tags []VCSRef)

	// GetTagsByHash returns all tags of the given commit hash
	GetTagsByHash(hash string) []VCSRef

	// FindCommitByHash will query a commit by hash, additionally includes changes made by the commit
	FindCommitByHash(hash string, includeChanges bool) (Commit, error)

	// FindCommitsBetween finds all commits between two references (might need to use GetReference to get the proper ref name)
	FindCommitsBetween(from *VCSRef, to *VCSRef, includeChanges bool, limit int) ([]Commit, error)

	// FindLatestRelease finds the latest release starting from the current repo HEAD
	FindLatestRelease(stable bool) (VCSRelease, error)
}

// NewVCSRefFromString parses the input string and returns a VCSRef
func NewVCSRefFromString(input string) (*VCSRef, error) {
	if input == "" {
		return nil, nil
	}

	inputParts := strings.SplitN(input, "/", 2)
	if len(inputParts) != 2 {
		return nil, fmt.Errorf("%s is not a valid vcs ref", input)
	}

	if inputParts[0] == "tag" {
		return &VCSRef{Type: "tag", Value: inputParts[1]}, nil
	} else if inputParts[0] == "branch" {
		return &VCSRef{Type: "branch", Value: inputParts[1]}, nil
	} else if inputParts[0] == "hash" {
		return &VCSRef{Type: "hash", Hash: inputParts[1]}, nil
	}

	return nil, fmt.Errorf("%s is not a valid vcs ref", input)
}
