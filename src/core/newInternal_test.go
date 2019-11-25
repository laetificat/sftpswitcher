package core

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func Test_writeConfig(t *testing.T) {
	testFile := filepath.Join(vscodeDirPath, "test-sftp.json")

	answers := map[string]string{}
	answers["name"] = "testName"
	answers["host"] = "testHost"
	answers["username"] = "testUsername"
	answers["password"] = "testPassword"
	answers["remote path"] = "testRemotePath"

	err := writeConfig(answers, testFile)
	if err != nil {
		t.Error(err)
	}

	configBytes, err := ioutil.ReadFile(testFile)
	if err != nil {
		t.Error(err)
	}

	for _, v := range answers {
		if !strings.Contains(string(configBytes), v) {
			t.Errorf("\nError (Equals):\nexpected: %s\nactual: %s", v, string(configBytes))
		}
	}

	err = os.Remove(testFile)
	if err != nil && !os.IsNotExist(err) {
		t.Error(err)
	}
}
