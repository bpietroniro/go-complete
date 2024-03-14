package main

import (
	"fmt"
	"math"
)

const inflationRate = 6.5

func main() {
	var investmentAmount, years float64
	expectedReturnRate := 5.5

	fmt.Print("Investment Amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)

	fmt.Print("Number of Years: ")
	fmt.Scan(&years)

	futureValue, futureRealValue := calculateFutureValue(investmentAmount, expectedReturnRate, years)

	formattedFV := fmt.Sprintf("Future Value: %.1f\n", futureValue)
	formattedRFV := fmt.Sprintf("Future Value (adjusted for inflation): %.1f\n", futureRealValue)

	fmt.Print(formattedFV)
	fmt.Print(formattedRFV)
}

func calculateFutureValue(investmentAmount, expectedReturnRate, years float64) (fv float64, rfv float64) {
	fv = investmentAmount * math.Pow((1+expectedReturnRate/100), years)
	rfv = fv / math.Pow(1+inflationRate/100, years)

	return
}
