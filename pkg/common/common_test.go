package common

import (
	"testing"
	"os"
)

func TestMain(m *testing.M) {
    SetupTestLogger()
    code := m.Run()
    os.Exit(code)
}

func TestGetHostFromURL(t *testing.T) {
	host := GetHostFromURL("http://github.com/user/repository")
	if host != "github.com" {
		t.Errorf("Host should be github.com, not "+host+"!")
	}
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
