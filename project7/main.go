package main

import "fmt"

func main() {
	prices := []float64{10, 20, 30}
	taxRates := []float64{0, 0.07, 0.1, 0.15}

	result := make(map[float64][]float64)

	for _, taxRate := range taxRates {
		taxIncludeRates := make([]float64, len(prices))

		for priceIndex, priceValue := range prices {
			taxIncludeRates[priceIndex] = priceValue * (taxRate + 1)
		}
		result[taxRate] = taxIncludeRates
	}

	fmt.Println(result)
}
