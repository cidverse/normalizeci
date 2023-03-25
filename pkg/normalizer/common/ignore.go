package common

import (
	"os"
	"strings"

	ignore "github.com/sabhiram/go-gitignore"
)

func ProcessIgnoreFiles(files []string) *ignore.GitIgnore {
	var ignoreLines []string

	for _, file := range files {
		if _, err := os.Stat(file); err == nil {
			content, contentErr := os.ReadFile(file)
			if contentErr == nil {
				for _, l := range strings.Split(string(content), "\n") {
					ignoreLines = append(ignoreLines, l)
				}
			}
		}
	}

	return ignore.CompileIgnoreLines(ignoreLines...)
}
