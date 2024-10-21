package main

import (
	"fmt"
)

func main() {
	accountBal := 1000.00

	fmt.Println("Welcome to the bank!")

	for {
		fmt.Println("What do you want to do?")
		fmt.Println("1. Check Balance")
		fmt.Println("2. Deposit Money")
		fmt.Println("3. Withdraw Money")
		fmt.Println("4. Exit")

		var choice int

		fmt.Print("Enter your choice: ")
		fmt.Scan(&choice)
		fmt.Println("Your Choice:", choice)

		if choice == 1 {
			fmt.Println("######### Your Balance #########")
			fmt.Println("Account Balance", accountBal)
		} else if choice == 2 {
			fmt.Println("######### Deposit Money #########")
			fmt.Print("Enter amount for deposit: ")
			var depositAmount float64
			fmt.Scan(&depositAmount)

			if depositAmount < 0 {
				fmt.Println("Enter greater amount than 0!")
				continue

			}

			accountBal += depositAmount
			fmt.Println("New Update Balance:", accountBal)
		} else if choice == 3 {
			fmt.Println("######### Withdraw Money #########")
			fmt.Print("Enter amount to withdraw: ")
			var withDrawAmount float64
			fmt.Scan(&withDrawAmount)
			accountBal -= withDrawAmount
			fmt.Println("Your new Balance:", accountBal)
		} else if choice == 4 {
			fmt.Println("Goodbye!")
			break
		} else {
			fmt.Println("Invalid Choice! Enter Valid Choice")
			continue
		}
	}

	fmt.Println("Thank you for Visiting us!")
}
