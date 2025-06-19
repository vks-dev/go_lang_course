package main

import "fmt"

func main () {

	var income, expenses, taxRate float64

	fmt.Print("Enter your total income: ")
	fmt.Scan(&income)

	fmt.Print("Enter your total expenses: ")
	fmt.Scan(&expenses)

	fmt.Print("Enter your tax rate (as a percentage): ")
	fmt.Scan(&taxRate)

	EarningsBeforetax := income - expenses
	EarningsAfterTax := EarningsBeforetax * (1 - taxRate/100)

	fmt.Printf("Your earnings before tax: %.2f\n", EarningsBeforetax)
	fmt.Printf("Tax deducted is: %.2f\n", EarningsBeforetax - EarningsAfterTax)
	fmt.Printf("Your profit after deducting tax: %.2f\n", EarningsAfterTax)

	ratio := EarningsAfterTax / income * 100
	fmt.Printf("Your profit after tax as a percentage of your total income: %.2f%%\n", ratio)
	fmt.Printf("you have earned this much amount this month, and this is your profit")

}
