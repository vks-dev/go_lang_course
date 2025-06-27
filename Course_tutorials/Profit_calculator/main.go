package main

import "fmt"

func main() {

	var income, expenses, taxRate float64

	fmt.Print("Enter your total income: ")
	fmt.Scan(&income)

	fmt.Print("Enter your total expenses: ")
	fmt.Scan(&expenses)

	fmt.Print("Enter your tax rate (as a percentage): ")
	fmt.Scan(&taxRate)

	if income < expenses {

		fmt.Println("Your expenses are greater than your income. You're losing money!")
		return

	} else if income > expenses && income >= 0 && expenses >= 0 {

		EarningsBeforetax := income - expenses
		EarningsAfterTax := EarningsBeforetax * (1 - taxRate/100)

		fmt.Printf("Your earnings before tax: %.2f/-\n", EarningsBeforetax)
		fmt.Printf("Tax deducted is: %.2f/-\n", EarningsBeforetax-EarningsAfterTax)
		fmt.Printf("Your profit after deducting tax: %.2f/-\n", EarningsAfterTax)

		ratio := EarningsAfterTax / income * 100
		fmt.Printf("Your profit after tax as a percentage of your total income: %.2f%%\n", ratio)

		return

	} else {

		fmt.Println("Please enter valid non-negative numbers for income and expenses.")
		return

	}

}
