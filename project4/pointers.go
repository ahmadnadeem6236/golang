package main

import "fmt"

func main() {
	age := 32

	fmt.Println("Age:", age)

	agePointer := &age

	fmt.Println(agePointer)

	// adultYears := getAdultYears(age)
	// fmt.Println(adultYears)
}

// func getAdultYears(age int) int {
// 	return age - 18
// }
