package vcsapi

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVCSRefEmpty(t *testing.T) {
	ref, err := NewVCSRefFromString("")

	assert.Nil(t, ref)
	assert.NoError(t, err)
}

func TestVCSRefTag(t *testing.T) {
	ref, err := NewVCSRefFromString("tag/v1.0.0")

	assert.Equal(t, "tag", ref.Type)
	assert.Equal(t, "v1.0.0", ref.Value)
	assert.NoError(t, err)
}

func TestVCSRefBranch(t *testing.T) {
	ref, err := NewVCSRefFromString("branch/main")

	assert.Equal(t, "branch", ref.Type)
	assert.Equal(t, "main", ref.Value)
	assert.NoError(t, err)
}

func TestVCSRefHash(t *testing.T) {
	ref, err := NewVCSRefFromString("hash/7793ba4898570d41baf4299accf950e517f76db9")

	assert.Equal(t, "hash", ref.Type)
	assert.Equal(t, "7793ba4898570d41baf4299accf950e517f76db9", ref.Hash)
	assert.NoError(t, err)
}

func TestVCSRefErr(t *testing.T) {
	ref, err := NewVCSRefFromString("invalid-ref")

	assert.Nil(t, ref)
	if assert.Error(t, err) {
		assert.Equal(t, errors.New("invalid-ref is not a valid vcs ref"), err)
	}
}
