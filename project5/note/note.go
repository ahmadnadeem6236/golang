package note

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

type Note struct {
	Title     string
	Content   string
	CreatedAt time.Time
}

func New(title, content string) (Note, error) {
	if title == "" || content == "" {
		return Note{}, errors.New("invalid note")
	}

	return Note{
		Title:     title,
		Content:   content,
		CreatedAt: time.Now(),
	}, nil
}

func (note Note) Display() {
	fmt.Printf("You note titled %v has the following content:\n%v\n", note.Title, note.Content)
}

func (note Note) Save() error {
	filename := strings.Replace(note.Title, " ", "_", -1)
	filename = strings.ToLower(filename) + ".json"

	json, err := json.Marshal(note)
	if err != nil {
		fmt.Println(err)
		return err
	}

	os.WriteFile(filename, json, 0644)
	return nil
}
