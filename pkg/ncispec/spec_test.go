package ncispec

import (
	"os"
	"testing"

	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	zerolog.SetGlobalLevel(zerolog.DebugLevel)
	code := m.Run()
	os.Exit(code)
}

func TestOfMap(t *testing.T) {
	data := make(map[string]string)
	data[NCI_VERSION] = "1.0.0"
	nci := OfMap(data)

	assert.Equal(t, "1.0.0", nci.Version)
}

func TestToMap(t *testing.T) {
	spec := NormalizeCISpec{
		Version: "1.0.0",
	}
	data := ToMap(spec)

	assert.Equal(t, "1.0.0", data[NCI_VERSION])
}
