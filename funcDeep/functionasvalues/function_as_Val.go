package functionasvalues

import "fmt"

type transformfn func(int) int

func Functionasvalues() {
	numbers := []int{1, 2, 3, 4}
	doubled := transformNumbers(&numbers, double) //passing functions as value
	tripled := transformNumbers(&numbers, triple)

	fmt.Println(doubled)
	fmt.Println(tripled)

	transformN1 := getTransformed(&numbers)
	res := transformNumbers(&numbers, transformN1)

	fmt.Println(res)

}

func transformNumbers(numbers *[]int, transform transformfn) []int { //transform is function type argument
	newNumbers := []int{}

	for _, val := range *numbers {
		newNumbers = append(newNumbers, transform(val))
	}
	return newNumbers
}

func getTransformed(numbers *[]int) transformfn {
	if (*numbers)[0] == 1 {
		return double //returning fn as value
	} else {
		return triple //returning fn as value
	}
}

func double(val int) int {
	return val * 2
}

func triple(val int) int {
	return val * 3
}
