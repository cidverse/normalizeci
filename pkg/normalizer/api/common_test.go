package api

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetHostFromURL(t *testing.T) {
	host, err := GetHostFromURL("http://github.com/user/repository")
	assert.NoError(t, err)
	assert.Equal(t, "github.com", host)
}

func TestGetDirectoryNameFromPath(t *testing.T) {
	var dirName string

	dirName = GetDirectoryNameFromPath("/home/arnie/amelia.jpg")
	if dirName != "arnie" {
		t.Errorf(dirName + "should be arnie!")
	}

	dirName = GetDirectoryNameFromPath("/mnt/photos/")
	if dirName != "photos" {
		t.Errorf(dirName + "should be photos!")
	}

	dirName = GetDirectoryNameFromPath("/usr/local//go")
	if dirName != "local" {
		t.Errorf(dirName + "should be local!")
	}
}
