package core

import (
	"fmt"
	"io"
	"io/ioutil"
	"strings"
)

/*
GetAll returns a list of config file names that are found in the project's .vscode directory.
*/
func GetAll(vscodeDirPath string) ([]string, error) {
	vscodeDir, err := ioutil.ReadDir(vscodeDirPath)
	if err != nil {
		return nil, err
	}

	sftpConfigSlice := []string{}
	for _, v := range vscodeDir {
		if strings.Contains(v.Name(), "-sftp.json") {
			sftpConfigSlice = append(sftpConfigSlice, v.Name())
		}
	}

	return sftpConfigSlice, nil
}

/*
ShowList outputs the list of strings to each line to the terminal with the -sftp.json part trimmed off.
*/
func ShowList(list []string, writer io.Writer) {
	if len(list) < 1 {
		_, err := writer.Write([]byte("There are no configurations."))
		if err != nil {
			fmt.Println(err)
		}
		return
	}

	for _, v := range list {
		_, err := writer.Write([]byte(fmt.Sprintln(strings.Replace(v, "-sftp.json", "", -1))))
		if err != nil {
			fmt.Println(err)
		}
	}
}
