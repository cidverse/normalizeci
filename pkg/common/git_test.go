package common

import (
	"testing"
	"io/ioutil"
	"os"

	"gopkg.in/src-d/go-git.v4"
)

func TestOnInvalidGitDirectory(t *testing.T) {
	var scmArgs = GetSCMArguments(GetGitDirectory()+"/../invalidpath")

	// log all normalized values
	for _, envvar := range scmArgs {
		t.Log(envvar)
	}
}

func TestOnEmptyGitRepository(t *testing.T) {
	// create empty repo
	dir, dirErr := ioutil.TempDir("", "normalizeci")
	if dirErr != nil {
		t.Errorf(dirErr.Error())
	}

	// init empty repo
	_, gitErr := git.PlainInit(dir, false)
	if gitErr != nil {
		t.Errorf(gitErr.Error())
	}
	
	// call func
	var scmArgs = GetSCMArguments(dir)

	// log all normalized values
	for _, envvar := range scmArgs {
		t.Log("Args:"+envvar)
	}

	// clean up
	defer os.RemoveAll(dir)
}
