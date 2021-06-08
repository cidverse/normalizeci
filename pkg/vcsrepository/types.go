package vcsrepository

type Commit struct {
	Hash string
	Message string
	Description string
	Author CommitAuthor
	Committer CommitAuthor
	Tags []CommitTag
}

type CommitAuthor struct {
	Name string
	Email string
}

type CommitTag struct {
	Name string
	VCSRef string
}

type Release struct {
	Name string
	Reference string
}
