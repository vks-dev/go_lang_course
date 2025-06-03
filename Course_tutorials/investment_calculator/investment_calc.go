package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate float64= 6.8
	var principalAmount, years float64 = 10000, 10
	interestRate := 12.8
	// principalAmount, years, interestRate := 10000.0, 10.0, 12.8         Alternate Notiation for variable declaration

	var finalAmount = principalAmount * math.Pow(1+(interestRate/100), years)
	fmt.Printf("Final Amount is: %.2f\n", finalAmount)

	profit := finalAmount - principalAmount
	fmt.Printf("Profit is: %.2f\n", profit)

	inflationAdjustedValue := finalAmount / math.Pow(1 + (inflationRate / 100), years)
	fmt.Printf("Inflation adjusted final amount is : %.2f \n", inflationAdjustedValue)

	// fmt.Printf("After investing %v for %v years at an interest rate of %v %, the final amount is : %v\n", principalAmount, years, interestRate, finalAmount);
}
