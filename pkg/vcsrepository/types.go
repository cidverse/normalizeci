package vcsrepository

type Commit struct {
	Message string
	Description string
	// TODO: add author details
}

type Release struct {
	Name string
	Reference string
}
