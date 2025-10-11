package main

import "fmt"

var initialAmt int = 10000

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
		fmt.Println("You've chosen option 1")
		balanceCheck()
	case 2:
		fmt.Println("You've chosen option 2")
		addMoney()
	case 3:
		fmt.Println("You've chosen option 3")
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
	var deposit int
	fmt.Print("How much would you like to deposit? ")
	fmt.Scan(&deposit)

	initialAmt = initialAmt + deposit
	fmt.Printf("You've successfully deposited %v/- into your account. Total Balance is: %v/- \n", deposit, initialAmt)
}

func withdrawMoney() {
	var withdrawAmt int
	fmt.Print("How much would you like to withdraw? ")
	fmt.Scan(&withdrawAmt)

	if withdrawAmt > initialAmt {
		fmt.Println("You're withdrawing more amount than what is available in your account. Please try again")
		return
	}

	initialAmt = initialAmt - withdrawAmt
	fmt.Printf("You've successfully withdrawn %v/- from your account. Remaining balance is: %v/- \n", withdrawAmt, initialAmt)
}
