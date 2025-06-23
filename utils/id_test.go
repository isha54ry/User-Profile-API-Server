package utils

import (
	"testing"
)

func TestGenerateID(t *testing.T) {
	id := GenerateID()
	if len(id) == 0 {
		t.Errorf("Expected a non-empty ID, got '%s'", id)
	}
}
