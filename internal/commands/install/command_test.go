package install

import (
	"testing"
)

func TestCommandCreation(t *testing.T) {
	command := Install{}

	result := command.Run(nil)

	if result != nil {
		t.Errorf("Expected nil, got %v", result)
	}
}
