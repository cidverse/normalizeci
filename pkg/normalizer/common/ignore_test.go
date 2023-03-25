package common

import (
	"os"
	"testing"
)

func TestProcessIgnoreFiles(t *testing.T) {
	// Create temporary files with ignore patterns
	tempFile1, err := os.CreateTemp("", "ignore1")
	if err != nil {
		t.Fatalf("failed to create temporary file: %v", err)
	}
	defer os.Remove(tempFile1.Name())
	if _, err := tempFile1.WriteString("*.tmp\n"); err != nil {
		t.Fatalf("failed to write to temporary file: %v", err)
	}

	tempFile2, err := os.CreateTemp("", "ignore2")
	if err != nil {
		t.Fatalf("failed to create temporary file: %v", err)
	}
	defer os.Remove(tempFile2.Name())
	if _, err := tempFile2.WriteString("*.log\n"); err != nil {
		t.Fatalf("failed to write to temporary file: %v", err)
	}

	// Call the function with the temporary files
	files := []string{tempFile1.Name(), tempFile2.Name()}
	gitIgnore := ProcessIgnoreFiles(files)

	// Verify
	if !gitIgnore.MatchesPath("file.tmp") {
		t.Errorf("Expected 'file.tmp' to be ignored")
	}
	if !gitIgnore.MatchesPath("/path/to/file.log") {
		t.Errorf("Expected '/path/to/file.log' to be ignored")
	}
	if gitIgnore.MatchesPath("file.txt") {
		t.Errorf("Expected file.txt to not be ignored")
	}
}
