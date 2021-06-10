package vcsrepository

import "time"

type Commit struct {
	ShortHash   string
	Hash        string
	Message     string
	Description string
	Author      CommitAuthor
	Committer   CommitAuthor
	Tags        []CommitTag
	AuthoredAt  time.Time
	CommittedAt time.Time
	Context     map[string]string
}

type CommitAuthor struct {
	Name  string
	Email string
}

type CommitTag struct {
	Name   string
	VCSRef string
}

type Release struct {
	Name      string
	Reference string
}
