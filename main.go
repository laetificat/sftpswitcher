package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/laetificat/sftpchanger/cmd"
	"github.com/laetificat/sftpchanger/src/config"
	"github.com/laetificat/sftpchanger/src/core"
	"github.com/spf13/viper"
)

func main() {
	loadViper()

	err := cmd.Execute()
	if err != nil {
		panic(err)
	}
}

/*
loadViper initiates viper and loads the config, if no config is found it will
create one.
*/
func loadViper() {
	viper.SetConfigName("sftpchangerconfig")
	viper.AddConfigPath(".")
	viper.AddConfigPath("$HOME/.sftpchanger")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			err := createNewConfig()
			if err != nil {
				panic(err)
			}
			loadViper()
		} else {
			panic(err)
		}
	}

	if err := viper.Unmarshal(&config.Cfg); err != nil {
		panic(err)
	}
}

/*
createNewConfig creates a new configuration file in the user's home directory
and asks the user where the location of the forest app is
*/
func createNewConfig() error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	configPath := filepath.Join(homeDir, ".sftpchanger")
	fmt.Printf("No config found, creating one in %s\n", configPath)
	err = os.MkdirAll(configPath, 0777) // 0777 because umask breaks stuff
	if err != nil {
		return err
	}

	err = os.Chmod(configPath, 0755) // chmod to set the right permissions
	if err != nil {
		return err
	}

	reader := bufio.NewReader(os.Stdin)
	answers, err := core.AskUserForInfo(reader, []string{"forest app location"})
	if err != nil {
		return err
	}

	configFile := filepath.Join(configPath, "sftpchangerconfig.toml")
	contents := []byte(fmt.Sprintf("forest_app_location = \"%s\"", answers["forest app location"]))
	err = ioutil.WriteFile(configFile, contents, 0644)
	if err != nil {
		return err
	}

	fmt.Printf("Successfully created %s\n", configFile)

	return nil
}
