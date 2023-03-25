package gitlabci

import (
	"github.com/cidverse/normalizeci/pkg/ncispec"
)

func (n Normalizer) Denormalize(spec ncispec.NormalizeCISpec) map[string]string {
	data := make(map[string]string)

	data["GITLAB_CI"] = "true"

	return data
}
