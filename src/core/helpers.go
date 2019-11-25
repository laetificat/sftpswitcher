package core

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

/*
guessProjectName returns the name of the project.
*/
func guessProjectName() (string, error) {
	path, err := os.Getwd()
	if err != nil {
		return "", err
	}

	return filepath.Base(path), nil
}

/*
AskUserForInfo asks the user some questions to fill in the config template.
*/
func AskUserForInfo(reader *bufio.Reader, itemsToAsk []string) (map[string]string, error) {
	answers := map[string]string{}

	fmt.Println("Please fill in the following")
	for _, v := range itemsToAsk {
		fmt.Printf("%s: ", v)
		answer, err := reader.ReadString('\n')
		if err != nil {
			return nil, err
		}

		answers[v] = strings.TrimRight(answer, "\n")
	}

	return answers, nil
}
