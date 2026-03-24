package tests

import (
	path_size "code"
	"testing"
)

func TestGetSize(t *testing.T) {
	size, err := path_size.GetSize("data/test.csv")
	if err != nil {
		t.Errorf("Error %v", err)
	}
	if size != 7 {
		t.Errorf("Size not match")
	}
}

func TestGetSizeError(t *testing.T) {
	_, err := path_size.GetSize("unknown")
	if err == nil {
		t.Errorf("Unknown dir, want err")
	}
}
