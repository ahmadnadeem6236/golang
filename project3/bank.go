package main

import (
	"fmt"
	"os"
	"strconv"
)

func getBalFromFile() float64 {
	accountbal, _ := os.ReadFile("balance.txt")
	balTxt := string(accountbal)
	mainBal, _ := strconv.ParseFloat(balTxt, 64)
	return mainBal

}

func writeBalToFile(balance float64) {
	accountbalance := fmt.Sprint(balance)
	os.WriteFile("balance.txt", []byte(accountbalance), 0644)
}

func main() {
	accountBal := getBalFromFile()

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
			writeBalToFile(accountBal)
		} else if choice == 3 {
			fmt.Println("######### Withdraw Money #########")
			fmt.Print("Enter amount to withdraw: ")
			var withDrawAmount float64
			fmt.Scan(&withDrawAmount)
			accountBal -= withDrawAmount
			fmt.Println("Your new Balance:", accountBal)
			writeBalToFile(accountBal)
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
