package core

import "testing"

import "strings"

import "bufio"

func Test_AskUserForInfo(t *testing.T) {
	expected := map[string]string{"question 1": "answer 1", "question 2": "answer 2"}

	itemsToAsk := []string{}
	answers := []string{}
	for k, v := range expected {
		itemsToAsk = append(itemsToAsk, k)
		answers = append(answers, v)
	}

	stringReader := strings.NewReader(strings.Join(answers, "\n") + "\n")
	reader := bufio.NewReader(stringReader)

	actual, err := AskUserForInfo(reader, itemsToAsk)
	if err != nil {
		t.Error(err)
	}

	for k, v := range expected {
		if val, ok := actual[k]; ok {
			if val != v {
				t.Errorf("\nError (Equals):\nexpected: %s\nactual: %s", v, val)
			}
		} else {
			t.Errorf("\nError (Equals):\nexpected: %s\nactual: %s", k, "")
		}
	}
}

func Test_AskUserForInfoInvalidBuffer(t *testing.T) {
	q_and_a := map[string]string{"question 1": "answer 1", "question 2": "answer 2"}

	itemsToAsk := []string{}
	answers := []string{}
	for k, v := range q_and_a {
		itemsToAsk = append(itemsToAsk, k)
		answers = append(answers, v)
	}

	stringReader := strings.NewReader(strings.Join(answers, "\n"))
	reader := bufio.NewReader(stringReader)

	_, err := AskUserForInfo(reader, itemsToAsk)

	if !strings.Contains(err.Error(), "EOF") {
		t.Errorf("\nError (Equals):\nexpected: %s\nactual: %s", "EOF", err.Error())
	}
}

func Test_AskUserForInfoNoQuestions(t *testing.T) {
	expected := map[string]string{}

	itemsToAsk := []string{}
	answers := []string{}
	for k, v := range expected {
		itemsToAsk = append(itemsToAsk, k)
		answers = append(answers, v)
	}

	stringReader := strings.NewReader(strings.Join(answers, "\n") + "\n")
	reader := bufio.NewReader(stringReader)

	actual, err := AskUserForInfo(reader, itemsToAsk)
	if err != nil {
		t.Error(err)
	}

	if len(actual) > 0 {
		t.Errorf("\nError (Equals):\nexpected: %s\nactual: %s", string(0), string(len(actual)))
	}
}
