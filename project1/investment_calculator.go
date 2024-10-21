package main

import (
	"fmt"
	"math"
)

func futureVal(investmentAmount float64, expectedReturn float64, years float64) float64 {
	result := investmentAmount * math.Pow(1+expectedReturn/100, years)
	return result
}

func futureRval(futureValue float64, inflationRate float64, years float64) float64 {
	result := futureValue / math.Pow(1+inflationRate/100, years)
	return result
}

func main() {
	const inflationRate = 2.5
	var investmentAmount float64
	var expectedReturn float64
	var years float64

	fmt.Print("Invested Amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Expected Return: ")
	fmt.Scan(&expectedReturn)

	fmt.Print("Years: ")
	fmt.Scan(&years)

	futureValue := futureVal(investmentAmount, expectedReturn, years)
	futureRealValue := futureRval(futureValue, inflationRate, years)

	fmt.Println()
	fmt.Printf("Future Amount you get: %0.2f\n", futureValue)
	fmt.Printf("Future Amount you get after inflation: %0.2f\n", futureRealValue)
}
