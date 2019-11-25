package core

import (
	"github.com/laetificat/sftpchanger/src/config"
	"bufio"
	"bytes"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"testing"
)

var vscodeDirPath string

func TestMain(m *testing.M) {
	var err error

	defer tearDown()

	vscodeDirPath, err = filepath.Abs(filepath.Join(os.TempDir(), "/sftpchanger/testing/.vscode"))
	config.Cfg.ForestAppLocation = strings.Replace(vscodeDirPath, "/.vscode", "", -1)

	if err != nil {
		panic(err)
	}

	err = os.MkdirAll(vscodeDirPath, 0777)
	if err != nil {
		panic(err)
	}

	os.Exit(m.Run())
}

/*
tearDown cleans up after tests, think about disconnecting from DB's or cleaning
up files
*/
func tearDown() {
	if err := os.RemoveAll(vscodeDirPath); err != nil {
		panic(err)
	}
}

func createTestConfigs(files []string) error {
	for _, v := range files {
		err := ioutil.WriteFile(filepath.Join(vscodeDirPath, v), []byte(""), 0777)
		if err != nil {
			return err
		}
	}

	return nil
}

func removeTestConfigs(files []string) error {
	for _, v := range files {
		err := os.Remove(filepath.Join(vscodeDirPath, v))
		if err != nil {
			return err
		}
	}

	return nil
}

func Test_GetAllNoConfigs(t *testing.T) {
	configs, err := GetAll(vscodeDirPath)
	if err != nil {
		t.Error(err)
	}

	if len(configs) != 0 {
		t.Errorf("\nError (Equals):\nexpected: %s\nactual: %s", "0", string(len(configs)))
	}
}

func Test_GetAllErrorReadDir(t *testing.T) {
	_, err := GetAll("/fly/me/to/the/moon")
	if err != nil {
		result := strings.Contains(err.Error(), "open /fly/me/to/the/moon: no such file or directory")
		if !result {
			t.Errorf("\nError (Equals):\nexpected: %s\nactual: %s", "open /fly/me/to/the/moon: no such file or directory", strconv.FormatBool(result))
		}
	} else {
		t.Errorf("\nError (Equals):\nexpected: %s\nactual: %s", "open /fly/me/to/the/moon: no such file or directory", "")
	}
}

func Test_GetAll(t *testing.T) {
	files := []string{"google-sftp.json", "apple-sftp.json", "amazon-sftp.json"}

	err := createTestConfigs(files)
	if err != nil {
		t.Error(err)
	}

	results, err := GetAll(vscodeDirPath)
	if err != nil {
		t.Error(err)
	}

	if len(results) < 2 {
		t.Errorf("\nError (Equals):\nexpected: %s\nactual: %s", "2", string(len(results)))
	}

	err = removeTestConfigs(files)
	if err != nil {
		t.Error(err)
	}
}

func Test_ShowListEmpty(t *testing.T) {
	expected := "There are no configurations."

	buffer := bytes.Buffer{}
	writer := bufio.NewWriter(&buffer)

	ShowList([]string{}, writer)
	writer.Flush()

	actual := buffer.String()

	if !strings.Contains(actual, expected) {
		t.Errorf("\nError (Equals):\nexpected: %s\nactual: %s", expected, actual)
	}
}

func Test_ShowListThree(t *testing.T) {
	files := []string{"google-sftp.json", "apple-sftp.json", "amazon-sftp.json"}

	err := createTestConfigs(files)
	if err != nil {
		t.Error(err)
	}

	buffer := bytes.Buffer{}
	writer := bufio.NewWriter(&buffer)

	ShowList(files, writer)
	writer.Flush()

	output := buffer.String()

	for _, v := range files {
		name := strings.Replace(v, "-sftp.json", "", -1)
		if !strings.Contains(output, name) {
			t.Errorf("\nError (Equals):\nexpected: %s\nactual: %s", name, output)
		}
	}

	err = removeTestConfigs(files)
	if err != nil {
		t.Error(err)
	}
}
