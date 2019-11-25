package core

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/laetificat/sftpchanger/src/config"
)

func Test_ChangeConfigNoName(t *testing.T) {
	curDir, err := os.Getwd()
	expected := filepath.Base(curDir)
	if err != nil {
		t.Error(err)
	}

	err = ioutil.WriteFile(filepath.Join(config.Cfg.ForestAppLocation, ".vscode", expected+"-sftp.json"), []byte{}, 0777)
	if err != nil {
		t.Error(err)
	}

	actual, err := ChangeConfig("")
	if err != nil {
		t.Error(err)
	}

	if !strings.Contains(expected, actual) {
		t.Errorf("\nError (Equals):\nexpected: %s\nactual: %s", expected, actual)
	}

	err = os.Remove(filepath.Join(vscodeDirPath, expected+"-sftp.json"))
	if err != nil {
		t.Error(err)
	}
}

func Test_ChangeConfigWithName(t *testing.T) {
	expected := "sol"

	err := ioutil.WriteFile(filepath.Join(config.Cfg.ForestAppLocation, ".vscode", expected+"-sftp.json"), []byte{}, 0777)
	if err != nil {
		t.Error(err)
	}

	actual, err := ChangeConfig(expected)
	if err != nil {
		t.Error(err)
	}

	if !strings.Contains(expected, actual) {
		t.Errorf("\nError (Equals):\nexpected: %s\nactual: %s", expected, actual)
	}

	err = os.Remove(filepath.Join(vscodeDirPath, expected+"-sftp.json"))
	if err != nil {
		t.Error(err)
	}
}

func Test_ChangeConfigWithNameNotFound(t *testing.T) {
	configName := "notexist"
	expected := fmt.Sprintf("Could not find config %s", configName)
	err := os.Remove(filepath.Join(vscodeDirPath, configName+"-sftp.json"))
	if err != nil && !os.IsNotExist(err) {
		t.Error(err)
	}

	actual, err := ChangeConfig(configName)
	if err == nil {
		t.Errorf("\nError (Equals):\nexpected: %s\nactual: %s", expected, actual)
	} else {
		if !strings.Contains(err.Error(), "could not find config") {
			t.Errorf("\nError (Equals):\nexpected: %s\nactual: %s", expected, err.Error())
		}
	}
}
