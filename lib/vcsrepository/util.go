package vcsrepository

import (
	"bufio"
	"github.com/Masterminds/semver/v3"
	"github.com/gosimple/slug"
	"io"
	"os"
	"strings"
)

// readLastLine gets the last line from a file, used to parse the git reflog
func readLastLine(filename string) string {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	defer file.Close()
	reader := bufio.NewReader(file)

	lastLine := ""
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		lastLine = string(line)
	}

	return lastLine
}

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

// isVersionStable checks if the specified version is a stable release version (semver)
func isVersionStable(versionStr string) bool {
	version, err := semver.NewVersion(versionStr)

	// no unparsable versions
	if err != nil {
		return false
	}

	// no prereleases
	if len(version.Prerelease()) > 0 {
		return false
	}

	return true
}
