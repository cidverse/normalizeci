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