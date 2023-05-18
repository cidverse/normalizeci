package common

// NCISpec is a common interface for all versions of the spec
type NCISpec interface {
	Validate() []ValidationError
}
