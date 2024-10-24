package main

import (
	"errors"
	"fmt"
)

func main() {
	getNoteData()
}

func getNoteData() (string, string, error) {
	tile, err := getUserInput("Note title: ")
	if err != nil {
		return "", "", err
	}
	content, err := getUserInput("Note content: ")
	if err != nil {
		return "", "", err
	}
	return tile, content, nil
}

func getUserInput(prompt string) (string, error) {
	fmt.Print(prompt)
	var value string
	fmt.Scanln(&value)

	if value == "" {
		return "", errors.New("please enter the values")
	}
	return value, nil
}
