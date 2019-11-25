package core

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/laetificat/sftpchanger/src/config"
)

/*
ChangeConfig copies the content of the given config to the sftp.json file.
*/
func ChangeConfig(configName string) (string, error) {
	var err error

	if len(configName) < 1 {
		configName, err = guessProjectName()
	}
	if err != nil {
		return configName, err
	}

	vscodeDir := filepath.Join(config.Cfg.ForestAppLocation, ".vscode")
	configPath := filepath.Join(vscodeDir, configName+"-sftp.json")

	_, err = os.Stat(configPath)
	if err != nil && os.IsNotExist(err) {
		return configName, fmt.Errorf("could not find config %s", configPath)
	} else if err != nil {
		return configName, err
	}

	configFile, err := ioutil.ReadFile(filepath.Join(vscodeDir, configName+"-sftp.json"))
	if err != nil {
		return configName, err
	}

	return configName, ioutil.WriteFile(filepath.Join(config.Cfg.ForestAppLocation, ".vscode", "sftp.json"), configFile, 0644)
}
