package main

import "fmt"

// factorial input 5 => 5*4*3*2*1

func main() {
	res := factorial(4)

	fmt.Println(res)

}

func factorial(number int) int {
	if number == 1 || number == 0 {
		return 1
	} else {
		return number * factorial(number-1)
	}

}
