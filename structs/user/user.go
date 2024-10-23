package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	FirstName string
	LastName  string
	Birthdate string
	CreatedAt time.Time
}

func NewUser(firstName, lastName, birthdate string) (*User, error) { // constructor and validations
	if firstName == "" || lastName == "" || birthdate == "" {
		return nil, errors.New("first name, last name and birthdate required")
	}
	return &User{
		FirstName: firstName,
		LastName:  lastName,
		Birthdate: birthdate,
		CreatedAt: time.Now(),
	}, nil
}

func (u User) OutputUserData() { // Method for struct

	fmt.Println(u.FirstName, u.LastName, u.Birthdate)
}
