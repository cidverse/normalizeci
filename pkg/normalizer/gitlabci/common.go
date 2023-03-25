package gitlabci

// Normalizer is the implementation of the normalizer
type Normalizer struct {
	version string
	name    string
	slug    string
}

// GetName returns the name of the normalizer
func (n Normalizer) GetName() string {
	return n.name
}

// GetSlug returns the slug of the normalizer
func (n Normalizer) GetSlug() string {
	return n.slug
}

// Check if this package can handle the current environment
func (n Normalizer) Check(env map[string]string) bool {
	return env["GITLAB_CI"] == "true"
}

// NewNormalizer gets a instance of the normalizer
func NewNormalizer() Normalizer {
	entity := Normalizer{
		version: "0.3.0",
		name:    "GitLab CI",
		slug:    "gitlab-ci",
	}

	return entity
}
