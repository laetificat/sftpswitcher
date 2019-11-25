package core

import "testing"

import "bufio"

import "strings"

func Test_CreateConfig(t *testing.T) {
	configName := "test-name"
	answers := []string{
		"testname",
		"testhost",
		"testusername",
		"testpassword",
		"testremotepath",
	}
	reader := bufio.NewReader(strings.NewReader(strings.Join(answers, "\n") + "\n"))

	answer, err := CreateConfig(reader, configName)
	if err != nil {
		t.Error(err)
	}

	if !strings.Contains(answer, "Successfully created") || !strings.Contains(answer, configName) {
		t.Errorf("\nError (Equals):\nexpected: %s\nactual: %s", "Successfully created "+configName, answer)
	}

	if err := removeTestConfigs([]string{configName + "-sftp.json"}); err != nil {
		t.Error(err)
	}
}

func Test_CreateConfigNoName(t *testing.T) {
	answers := []string{
		"testname",
		"testhost",
		"testusername",
		"testpassword",
		"testremotepath",
	}
	reader := bufio.NewReader(strings.NewReader(strings.Join(answers, "\n") + "\n"))

	answer, err := CreateConfig(reader, "")
	if err != nil {
		t.Error(err)
	}

	if !strings.Contains(answer, "Successfully created") || !strings.Contains(answer, "core") {
		t.Errorf("\nError (Equals):\nexpected: %s\nactual: %s", "Successfully created "+"core", answer)
	}

	if err := removeTestConfigs([]string{"core" + "-sftp.json"}); err != nil {
		t.Error(err)
	}
}

func Test_CreateConfigInvalidReader(t *testing.T) {
	answers := []string{
		"testname",
		"testhost",
		"testusername",
		"testpassword",
		"testremotepath",
	}
	reader := bufio.NewReader(strings.NewReader(strings.Join(answers, "\n")))

	_, err := CreateConfig(reader, "")
	if !strings.Contains(err.Error(), "EOF") {
		t.Errorf("\nError (Equals):\nexpected: %s\nactual: %s", "EOF", "")
	}
}
