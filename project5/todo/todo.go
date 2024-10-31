package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type Todo struct {
	Text string `json:"text"`
}

func New(text string) (Todo, error) {
	if text == "" {
		return Todo{}, errors.New("invalid note")
	}

	return Todo{
		Text: text,
	}, nil
}

func (todo Todo) Display() {
	fmt.Println(todo.Text)
}

func (todo Todo) Save() error {
	filename := "todo.json"

	json, err := json.Marshal(todo)
	if err != nil {
		return err
	}

	os.WriteFile(filename, json, 0644)
	return nil
}
