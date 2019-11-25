package core

import (
	"os"
	"strings"
	"testing"
)

func Test_guessProjectName(t *testing.T) {
	currentDir, err := os.Getwd()
	if err != nil {
		t.Error(err)
	}

	result, err := guessProjectName()
	if err != nil {
		t.Error(err)
	}

	if !strings.Contains(currentDir, result) {
		t.Errorf("\nError (Equals):\nexpected: %s\nactual: %s", currentDir, result)
	}
}
