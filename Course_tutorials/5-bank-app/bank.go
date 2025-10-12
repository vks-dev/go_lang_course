package main

import "fmt"

var initialAmt float64 = 10000

func main() {

	fmt.Println("Welcome to World Bank")
	fmt.Println("What would you like to do?")
	fmt.Println("1. Check Balance")
	fmt.Println("2. Deposit Money")
	fmt.Println("3. Withdraw Money")
	fmt.Println("4. Exit")

	var choice int
	fmt.Print("Enter your choice: ")
	fmt.Scan(&choice)

	switch choice {
	case 1:
		balanceCheck()
	case 2:
		addMoney()
	case 3:
		withdrawMoney()
	case 4:
		return

	default:
		fmt.Printf("Your choice: %v is invalid, please try again.\n", choice)
	}

}

func balanceCheck() {
	fmt.Printf("Your account has %v/- available for use.\n", initialAmt)
}

func addMoney() {
	var deposit float64
	fmt.Print("How much would you like to deposit? ")
	fmt.Scan(&deposit)

	if deposit <= 0 {
		fmt.Println("You're depositing an invalid amount. Deposit amount must be greater than 0. Please try again")
		return
	}

	initialAmt += deposit
	fmt.Printf("You've successfully deposited %v/- into your account. Total Balance is: %v/- \n", deposit, initialAmt)
}

func withdrawMoney() {
	var withdrawAmt float64
	fmt.Print("How much would you like to withdraw? ")
	fmt.Scan(&withdrawAmt)

	if withdrawAmt <= 0 {
		fmt.Println("You're withdrawing an invalid amount. Withdraw amount must be greater than 0. Please try again")
		return
	}

	if withdrawAmt > initialAmt {
		fmt.Println("You're withdrawing more amount than what is available in your account. Please try again")
		return
	}

	initialAmt -= withdrawAmt
	fmt.Printf("You've successfully withdrawn %v/- from your account. Remaining balance is: %v/- \n", withdrawAmt, initialAmt)
}
