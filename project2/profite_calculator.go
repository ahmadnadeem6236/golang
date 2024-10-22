package main

import (
	"errors"
	"fmt"
	"os"
)

func printScan(whatfor string, usefor float64) (float64, error) {
	fmt.Printf("%v: ", whatfor)
	fmt.Scan(&usefor)
	if usefor <= 0 {
		fmt.Println("Enter the amount greater than 0")
		return 0, errors.New("invalid number")
	}
	return usefor, nil
}

func calProfit(revenue float64, expense float64, taxRate float64) (ebt float64, profit float64, ratio float64) {
	ebt = revenue + expense
	profit = ebt * (1 - taxRate/100)
	ratio = ebt / profit
	return ebt, profit, ratio
}

func writeToFile(ebt float64, profit float64, ratio float64) {
	ebtStr, profitStr, ratioStr := fmt.Sprint(ebt), fmt.Sprint(profit), fmt.Sprint(ratio)
	result := []string{ebtStr, profitStr, ratioStr}
	os.WriteFile("report.txt", []byte(fmt.Sprint(result)), 0644)

}

func main() {

	var revenue float64
	var expense float64
	var taxRate float64

	revenue, errRev := printScan("Enter your Revenue", revenue)
	if errRev != nil {
		fmt.Println(errRev)
		return
	}

	expense, errExp := printScan("Enter your expense", expense)
	if errExp != nil {
		fmt.Println(errRev)
		return
	}
	taxRate, errTax := printScan("Enter your Tax Rate", taxRate)
	if errTax != nil {
		fmt.Println(errRev)
		return
	}

	ebt, profit, ratio := calProfit(revenue, expense, taxRate)
	writeToFile(ebt, profit, ratio)

	fmt.Println("EBT: ", ebt)

	fmt.Printf("Profit: %.2f\n", profit)

	fmt.Printf("Ratio: %.2f\n", ratio)

}
