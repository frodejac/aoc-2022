package io

import (
	"os"
	"testing"
)

func TestGetInputWithoutEnvVar(t *testing.T) {
	// Unset any existing session token
	os.Unsetenv("AOC_SESSION")

	_, ok := getExternalInput(1)

	if ok {
		t.Error("Expected GetInput to return false when no session token is set")
	}
}
