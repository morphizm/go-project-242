package tests

import (
	path_size "code"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetSizeSimple(t *testing.T) {
	size, err := path_size.GetSize("data/test.csv", false, false)
	assert.Nil(t, err)
	assert.Equal(t, 7, size)
}

func TestGetSizeError(t *testing.T) {
	_, err := path_size.GetSize("unknown", false, false)
	assert.NotNil(t, err)
}

func TestFormatSize(t *testing.T) {
	str := path_size.FormatSize(2000, false)
	assert.Regexp(t, "2000B", str)

	mb := path_size.FormatSize(7*1000*1000, true)
	assert.Regexp(t, "\\dMB", mb)
}

func TestHidden(t *testing.T) {
	size, _ := path_size.GetSize("data", true, false)
	assert.Equal(t, 15, size)
}

func TestRecursive(t *testing.T) {
	size, _ := path_size.GetSize("data", false, true)
	assert.Equal(t, 14, size)
}

func TestRecursiveAndHidden(t *testing.T) {
	size, _ := path_size.GetSize("data", true, true)
	assert.Equal(t, 22, size)
}
