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

	if income < expenses {

		OutputText("Your expenses are greater than your income. You're losing money!", "println")
		return

	} else if income > expenses && income >= 0 && expenses >= 0 {

		EarningsBeforetax, EarningsAfterTax := Calculations(income, expenses, taxRate)

		OutputText("Your earnings before tax: %.2f/-\n", "printf", EarningsBeforetax)
		OutputText("Tax deducted is: %.2f/-\n", "printf", EarningsBeforetax-EarningsAfterTax)
		OutputText("Your profit after deducting tax: %.2f/-\n", "printf", EarningsAfterTax)

		ratio := EarningsAfterTax / income * 100
		OutputText("Your profit after tax as a percentage of your total income: %.2f%%\n", "printf", ratio)

		return

	} else {

		OutputText("Please enter valid non-negative numbers for income and expenses.", "println")
		return

	}
}

func OutputText(text string, printType string, args ...any) {

	if printType == "printf" {
		fmt.Printf(text, args...)
		return
	} else {
		fmt.Println(text)
		return
	}
}

func Calculations(income, expenses, taxRate float64) (EarningsBeforetax float64, EarningsAfterTax float64) {

	EarningsBeforetax = income - expenses
	EarningsAfterTax = EarningsBeforetax * (1 - taxRate/100)

	return EarningsBeforetax, EarningsAfterTax
	// we can also use only the return keyword here, as the named return values will be returned automatically

}