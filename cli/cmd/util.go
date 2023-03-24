package cmd

import (
	"io"
	"os"

	"github.com/rs/zerolog/log"
)

func fileOutput(file string, content string) {
	contentByteArray := []byte(content)
	err := os.WriteFile(file, contentByteArray, 0644)
	if err != nil {
		log.Err(err).Str("file", file).Msg("failed to generate file")
	}
}

func consoleOutput(content string) {
	_, err := io.WriteString(os.Stdout, content)
	if err != nil {
		log.Err(err).Msg("failed to write content to stdout")
	}
}
