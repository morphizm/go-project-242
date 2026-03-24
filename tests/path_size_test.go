package tests

import (
	"code"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetSizeSimple(t *testing.T) {
	size, err := code.GetPathSize("data/test.csv", false, false, false)
	assert.Nil(t, err)
	assert.Equal(t, "7B", size)
}

func TestGetSizeError(t *testing.T) {
	_, err := code.GetPathSize("unknown", false, false, false)
	assert.NotNil(t, err)
}

func TestFormatSize(t *testing.T) {
	str := code.FormatSize(2000, false)
	assert.Regexp(t, "2000B", str)

	mb := code.FormatSize(7*1000*1000, true)
	assert.Regexp(t, "\\dMB", mb)
}

func TestHidden(t *testing.T) {
	size, _ := code.GetPathSize("data", false, false, true)
	assert.Equal(t, "15B", size)
}

func TestRecursive(t *testing.T) {
	size, _ := code.GetPathSize("data", true, false, false)
	assert.Equal(t, "14B", size)
}

func TestRecursiveAndHidden(t *testing.T) {
	size, _ := code.GetPathSize("data", true, false, true)
	assert.Equal(t, "22B", size)
}
