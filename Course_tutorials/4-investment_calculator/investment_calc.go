package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate float64= 6.8
	var principalAmount, years, interestRate float64
	// interestRate := 12.8;  principalAmount, years, interestRate := 10000.0, 10.0, 12.8         Alternate Notiations for variable declaration


	// USER INPUT

	fmt.Print("Enter the principal amount: ")
	fmt.Scan(&principalAmount) // `&` is used to get the address of the variable, and since we use Scan() to modify a variable, we need to pass its address

	fmt.Print("Enter the time period in years: ")
	fmt.Scan(&years)

	fmt.Print("Enter the interest rate: ")
	fmt.Scan(&interestRate)


	// CALCULATIONS

	var finalAmount = principalAmount * math.Pow(1+(interestRate/100), years)
	profit := finalAmount - principalAmount
	inflationAdjustedValue := finalAmount / math.Pow(1 + (inflationRate / 100), years)

	
	// OUTPUT
	
	fmt.Printf("Final Amount is: %.2f\n", finalAmount)

	fmt.Printf("Profit is: %.2f\n", profit)

	fmt.Printf("Inflation adjusted final amount is : %.2f \n", inflationAdjustedValue)

	// fmt.Printf("After investing %v for %v years at an interest rate of %v %, the final amount is : %v\n", principalAmount, years, interestRate, finalAmount);
}
