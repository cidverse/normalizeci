package appveyor

import (
	v1 "github.com/cidverse/normalizeci/pkg/ncispec/v1"
)

func (n Normalizer) Denormalize(spec v1.Spec) (map[string]string, error) {
	return make(map[string]string), nil
}
