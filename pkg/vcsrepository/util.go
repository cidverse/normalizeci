package vcsrepository

import (
	"github.com/gosimple/slug"
	"os"
	"strings"
)

// fileExists checks if the file exists and returns a boolean
func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if err != nil {
		return false
	}

	return !info.IsDir()
}

func getReleaseName(input string) string {
	input = slug.Substitute(input, map[string]string{
		"/": "-",
	})

	return strings.TrimLeft(input, "v")
}
