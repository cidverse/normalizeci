package vcsapi

import (
	"time"
)

type Commit struct {
	ShortHash   string            `json:"hash_short"`
	Hash        string            `json:"hash"`
	Message     string            `json:"message"`
	Description string            `json:"description"`
	Author      CommitAuthor      `json:"author"`
	Committer   CommitAuthor      `json:"committer"`
	Changes     []CommitChange    `json:"changes,omitempty"`
	Tags        []VCSRef          `json:"tags"`
	AuthoredAt  time.Time         `json:"authored_at"`
	CommittedAt time.Time         `json:"committed_at"`
	Context     map[string]string `json:"context,omitempty"`
}

type CommitChange struct {
	Type     string     `json:"type"`
	FileFrom CommitFile `json:"file_from"`
	FileTo   CommitFile `json:"file_to"`
	Patch    string     `json:"patch"`
}

type CommitFile struct {
	Name string `json:"name"`
	Size int64  `json:"size"`
	Hash string `json:"hash"`
}

type CommitAuthor struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type VCSRef struct {
	Type  string `json:"type"`
	Value string `json:"value"`
	Hash  string `json:"hash,omitempty"`
}

type VCSRelease struct {
	Type    string `json:"type"`
	Value   string `json:"value"`
	Version string `json:"version"`
	Hash    string `json:"hash,omitempty"`
}
