package gitlabci

import (
	v1 "github.com/cidverse/normalizeci/pkg/ncispec/v1"
)

func (n Normalizer) Denormalize(spec v1.Spec) map[string]string {
	data := make(map[string]string)

	data["GITLAB_CI"] = "true"

	return data
}
