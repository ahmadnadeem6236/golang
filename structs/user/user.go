package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct { //To export type we should Capatilize first letter otherwise we can use in methods or functions.
	firstName string
	lastName  string
	birthdate string
	createdAt time.Time
}

type Admin struct { //struct embedding
	email    string
	password string
	User
}

func NewAdmin(email, password, firstName, lastName string) Admin {
	return Admin{
		email:    email,
		password: password,
		User: User{
			firstName: firstName,
			lastName:  lastName,
		},
	}

}

func NewUser(firstName, lastName, birthdate string) (*User, error) { // constructor and validations
	if firstName == "" || lastName == "" || birthdate == "" {
		return nil, errors.New("first name, last name and birthdate required")
	}
	return &User{
		firstName: firstName,
		lastName:  lastName,
		birthdate: birthdate,
		createdAt: time.Now(),
	}, nil
}

func (u User) OutputUserData() { // Method for struct

	fmt.Println(u.firstName, u.lastName, u.birthdate)
}

func (a Admin) OutputAdminData() {
	fmt.Println(a.email, a.password)
}
