package main

// Normalizer is a common interface to work with all normalizers
type Normalizer interface {
	GetName() string
	Check(env []string) bool
	Normalize(env []string) []string
}
