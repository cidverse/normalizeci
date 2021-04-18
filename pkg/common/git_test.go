package common

import (
	"gopkg.in/src-d/go-git.v4"
	"os"
	"testing"
)

func TestOnInvalidGitDirectory(t *testing.T) {
	// log all normalized values
	var scmArgs = GetSCMArguments(GetGitDirectory()+"/tmp/invaliddir")
	for _, envvar := range scmArgs {
		t.Log(envvar)
	}
}

func TestOnEmptyGitRepository(t *testing.T) {
	// create git repo
	var tmpDir = GetGitDirectory()+"/tmp/empty"
	os.RemoveAll(GetGitDirectory()+"/tmp/empty")
	_, err := git.PlainInit(tmpDir, true)
	if err != nil {
		t.Errorf(err.Error())
	}
	
	// log all normalized values
	var scmArgs = GetSCMArguments(tmpDir)
	for _, envvar := range scmArgs {
		t.Log(envvar)
	}
	os.RemoveAll(GetGitDirectory()+"/tmp/empty")
}
