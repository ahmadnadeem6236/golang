package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/note/note"
	"example.com/note/todo"
)

type saver interface { //interface create contract b/w the assigned structs
	Save() error
}

type outputable interface {
	Display()
	saver
}

func main() {
	title, content := getNoteData()
	userNote, err := note.New(title, content)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = outputData(userNote)
	if err != nil {
		fmt.Println(err)
		return
	}

	text := getTodoData()
	userTodo, err := todo.New(text)
	if err != nil {
		fmt.Println(err)
		return
	}
	outputData(userTodo)

}

func outputData(data outputable) error {
	data.Display()
	return saveData(data)
}

func saveData(data saver) error {
	err := data.Save()

	if err != nil {
		fmt.Println("Saving todo is failed!")
		return err
	}
	fmt.Println("Saved")
	return nil
}

func getTodoData() string {
	text := getUserInput("Todo: ")
	return text

}

func getNoteData() (string, string) {
	tile := getUserInput("Note title: ")
	content := getUserInput("Note content: ")
	return tile, content
}

func getUserInput(prompt string) string {
	fmt.Print(prompt)
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}
