// package main

// import "fmt"

// func main() {
// 	prices := [4]float64{10.99, 4.55, 10.02, 29.9}
// 	fmt.Println(prices)
// }

package list

import "fmt"

type Product struct {
	title string
	id    int
	price float64
}

func List() {

	// hobbies := [3]string{"programming", "reading", "googling"}
	// fmt.Println(hobbies)
	// fmt.Println(hobbies[0])
	// newHobbies := hobbies[1:]
	// fmt.Println(newHobbies)
	// newHobbies = hobbies[1:3]
	// fmt.Println(newHobbies)
	// newHobbies = newHobbies[:1]
	// fmt.Println(newHobbies)
	// newHobbies = newHobbies[1:]
	// fmt.Println(newHobbies)

	// dynamicHobbies := []string{"searching", "sitting"}
	// dynamicHobbies[1] = "sleeping"
	// dynamicHobbies = append(dynamicHobbies, "listening")
	// fmt.Println(dynamicHobbies)

	dynamicProduct := []Product{{"iPhone", 1, 20.3}, {"Macbook", 2, 890.92}}
	fmt.Println(dynamicProduct)
	dynamicProduct = append(dynamicProduct, Product{"iPad", 3, 29.33})
	fmt.Println(dynamicProduct[2].title)

}

// Time to practice what you learned!

// 1) Create a new array (!) that contains three hobbies you have
// 		Output (print) that array in the command line.
// 2) Also output more data about that array:
//		- The first element (standalone)
//		- The second and third element combined as a new list
// 3) Create a slice based on the first element that contains
//		the first and second elements.
//		Create that slice in two different ways (i.e. create two slices in the end)
// 4) Re-slice the slice from (3) and change it to contain the second
//		and last element of the original array.
// 5) Create a "dynamic array" that contains your course goals (at least 2 goals)
// 6) Set the second goal to a different one AND then add a third goal to that existing dynamic array
// 7) Bonus: Create a "Product" struct with title, id, price and create a
//		dynamic list of products (at least 2 products).
//		Then add a third product to the existing list of products.
