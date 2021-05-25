package vcsrepository

// Conventional Commits - https://www.conventionalcommits.org/en/v1.0.0/
var ConventionalCommitPattern = `(?P<type>[A-Za-z]+)((?:\((?P<scope>[^()\r\n]*)\)|\()?(?P<breaking>!)?)(:\s?(?P<subject>.*))?`

type Commit struct {
	Message string
	Description string
	// TODO: add author details
}

type CommitVersionRule struct {
	Type string
	Scope string
	Release string // major / minor / patch
}

type Release struct {
	Name string
	Reference string
}

type ReleaseType int32
const (
	ReleaseNone  ReleaseType = 0
	ReleasePatch ReleaseType = 1
	ReleaseMinor ReleaseType = 2
	ReleaseMajor ReleaseType = 3
)

var DefaultReleaseVersionRules = []CommitVersionRule {
	{
		Type:    `feat`,
		Release: `minor`,
	},
	{
		Type:    `refactor`,
		Release: `minor`,
	},
	{
		Type:    `fix`,
		Release: `patch`,
	},
	{
		Type:    `ci`,
		Release: `patch`,
	},
	{
		Type:    `build`,
		Release: `patch`,
	},
	{
		Type:    `docs`,
		Release: `patch`,
	},
	{
		Type:    `perf`,
		Release: `patch`,
	},
	{
		Type:    `test`,
		Release: `patch`,
	},
	{
		Type:    `style`,
		Release: `patch`,
	},
}
