package main

import (
	"fmt"

	"example.com/structs/user"
)

func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	// ... do something awesome with that gathered data!
	appUser, err := user.NewUser(userFirstName, userLastName, userBirthdate)
	if err != nil {
		fmt.Println(err)
		return
	}

	appAdmin := user.NewAdmin("nadem@", "nad123", "Ahmad", "Nadeem")
	appAdmin.OutputAdminData()

	appUser.OutputUserData()
}

// func outputUserData(u *user) { // normal function

// 	fmt.Println(u.firstName, u.lastName, u.birthdate)

// }

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}
