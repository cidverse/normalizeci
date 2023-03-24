package githubactions

import (
	"github.com/cidverse/normalizeci/pkg/ncispec"
)

func (n Normalizer) Denormalize(spec ncispec.NormalizeCISpec) map[string]string {
	return make(map[string]string)
}
