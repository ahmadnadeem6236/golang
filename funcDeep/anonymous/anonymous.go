package anonymous

import "fmt"

func Anonymous() {
	numbers := []int{1, 2, 3}

	transformed := transformNumbers(&numbers, func(i int) int { //anonymous function
		return i * 4
	})

	fmt.Println(transformed)

	double := createTranformer(2)
	triple := createTranformer(3)

	transformed1 := transformNumbers(&numbers, double)
	transformed2 := transformNumbers(&numbers, triple)

	fmt.Println(transformed1)
	fmt.Println(transformed2)
}

func transformNumbers(numbers *[]int, transform func(int) int) []int {
	dNumbers := []int{}

	for _, val := range *numbers {
		dNumbers = append(dNumbers, transform(val))
	}

	return dNumbers
}

func createTranformer(factor int) func(int) int { // function closures with anonymous function
	return func(i int) int {
		return i * factor
	}
}
