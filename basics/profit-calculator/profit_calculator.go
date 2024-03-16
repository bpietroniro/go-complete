package main

import (
	"errors"
	"fmt"
	"os"
)

const financialsFile = "financials.txt"

// Goals
// 1) Validate user input
// => Show error message & exit if invalid input is provided
// - No negative numbers
// - Not 0
// 2) Store calculated results into file

func main() {
	revenue, err := userInput("Revenue: ")
	if err != nil {
		fmt.Println(err)
		return
	}
	expenses, err := userInput("Expenses: ")
	if err != nil {
		fmt.Println(err)
		return
	}
	taxRate, err := userInput("Tax Rate (percentage points): ")
	if err != nil {
		fmt.Println(err)
		return
	}

	ebt, profit, ratio := calculateProfitData(revenue, expenses, taxRate)

	fmt.Printf("%.2f\n", ebt)
	fmt.Printf("%.2f\n", profit)
	fmt.Printf("%.4f\n", ratio)
	writeResultsToFile(ebt, profit, ratio)
}

func userInput(prompt string) (float64, error) {
	var input float64
	fmt.Print(prompt)
	fmt.Scan(&input)

	if input <= 0 {
		return input, errors.New("nonnegative numbers only")
	}

	return input, nil
}

func calculateProfitData(revenue, expenses, taxRate float64) (ebt, profit, ratio float64) {
	ebt = revenue - expenses
	profit = ebt * (1 - taxRate/100)
	ratio = ebt / profit
	return
}

func writeResultsToFile(ebt, profit, ratio float64) {
	results := fmt.Sprintf("EBT: %.2f\nProfit: %.2f\nRatio: %.2f\n", ebt, profit, ratio)
	os.WriteFile(financialsFile, []byte(results), 0644)
}
