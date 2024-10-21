package main

import (
	"fmt"
)

func printScan(whatfor string, usefor float64) float64 {
	fmt.Printf("%v: ", whatfor)
	fmt.Scan(&usefor)
	return usefor
}

func calProfit(revenue float64, expense float64, taxRate float64) (ebt float64, profit float64, ratio float64) {
	ebt = revenue + expense
	profit = ebt * (1 - taxRate/100)
	ratio = ebt / profit
	return ebt, profit, ratio
}

func main() {

	var revenue float64
	var expense float64
	var taxRate float64

	revenue = printScan("Enter your Revenue", revenue)
	expense = printScan("Enter your Expense", expense)
	taxRate = printScan("Enter your Tax Rate", taxRate)

	ebt, profit, ratio := calProfit(revenue, expense, taxRate)

	fmt.Println("EBT: ", ebt)

	fmt.Printf("Profit: %.2f\n", profit)

	fmt.Printf("Ratio: %.2f\n", ratio)

}
