package core

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/laetificat/sftpchanger/src/config"
)

type sftpConfig struct {
	Name         string   `json:"name"`
	Host         string   `json:"host"`
	Protocol     string   `json:"protocol"`
	Port         uint     `json:"port"`
	Username     string   `json:"username"`
	Password     string   `json:"password"`
	RemotePath   string   `json:"remotePath"`
	UploadOnSave bool     `json:"uploadOnSave"`
	SyncMode     string   `json:"syncMode"`
	Ignore       []string `json:"ignore"`
}

/*
CreateConfig creates a new config with the given name, if the name is empty it will
try to guess the name by using the current directory.
*/
// TODO: Allow to pass reader so testing will be easier
func CreateConfig(configName string) (string, error) {
	var err error
	if len(configName) == 0 {
		configName, err = guessProjectName()
	}
	if err != nil {
		return configName, err
	}

	reader := bufio.NewReader(os.Stdin)
	answers, err := AskUserForInfo(reader, []string{"name", "host", "username", "password", "remote path"})
	if err != nil {
		return configName, err
	}

	configPath, err := filepath.Abs(filepath.Join(config.Cfg.ForestAppLocation, ".vscode", configName+"-sftp.json"))
	if err != nil {
		return configName, err
	}

	if err := writeConfig(answers, configPath); err != nil {
		return configName, err
	}

	return fmt.Sprintf("Successfully created %s", configPath), nil
}

/*
writeConfig uses the given answers to create a new config file based on the template
*/
func writeConfig(answers map[string]string, configPath string) error {
	configStruct := sftpConfig{
		Name:         answers["name"],
		Host:         answers["host"],
		Protocol:     "sftp",
		Port:         22,
		Username:     answers["username"],
		Password:     answers["password"],
		RemotePath:   answers["remote path"],
		UploadOnSave: true,
		SyncMode:     "update",
		Ignore: []string{
			"**/.vscode/**",
			"**/.DS_Store/**",
			"project_details.txt",
			".eslintrc.json",
			"**/node_modules/**",
			"**/.git/**",
		},
	}

	b, err := json.MarshalIndent(configStruct, "", "\t")
	if err != nil {
		return err
	}

	return ioutil.WriteFile(configPath, b, 0644)
}
