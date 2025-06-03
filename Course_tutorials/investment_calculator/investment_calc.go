package main

import (
	"fmt"
	"math"
)

func main() {
	var principalAmount float64 = 10000
	var interestRate = 12.8
	// years := 10
	var years float64 = 10;

	var finalAmount = principalAmount * math.Pow(1+(interestRate/100), years)

	fmt.Printf("Final Amount is: %.2f\n", finalAmount)

	profit := finalAmount - principalAmount

	fmt.Printf("Profit is: %.2f\n", profit)

	// fmt.Printf("After investing %v for %v years at an interest rate of %v %, the final amount is : %v\n", principalAmount, years, interestRate, finalAmount);
}
