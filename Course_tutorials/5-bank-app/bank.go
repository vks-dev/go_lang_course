package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

var accountFile string
var tries int64 = 0

func getBalance() (float64, error) {
	dataInByte, err := os.ReadFile(accountFile)
	if err != nil {
		if tries == 0 {
			return 0.0, errors.New("it seems we don't have an account under this name, So you'll be starting with 0.0 balance")
		} else {
			return 0.0, errors.New("unable to find account balance, starting with 0.0 balance")
		}
	}

	balanceInStr := string(dataInByte)
	balance, err := strconv.ParseFloat(balanceInStr, 64)
	if err != nil {
		return 0.0, errors.New("failed to parse data from file")
	}
	return balance, nil
}

func writeToFile(balance float64) {
	balanceStr := fmt.Sprintf("%.3f", balance)
	os.WriteFile(accountFile, []byte(balanceStr), 0644)
}

func main() {

	var userName string

	fmt.Print("Enter your Name: ")
	fmt.Scan(&userName)

	accountFile = userName + ".txt"

	initialAmt, err := getBalance()

	if err != nil {
		fmt.Printf("Error retrieving initial balance: %v\n", err)
		fmt.Println("--------------------------------------------")
	}
	tries++

	fmt.Println("Welcome to World Bank")
	for {
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
			balanceCheck(initialAmt)
		case 2:
			addMoney(initialAmt)
		case 3:
			withdrawMoney(initialAmt)
		case 4:
			fmt.Print("Thank you for choosing our bank. Have a great day!")
			return

		default:
			fmt.Printf("Your choice: %v is invalid, please try again.\n", choice)
		}
		writeToFile(initialAmt)
	}
}

func balanceCheck(initialAmt float64) {
	fmt.Printf("Your account has %v/- available for use.\n", initialAmt)
}

func addMoney(initialAmt float64) {
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

func withdrawMoney(initialAmt float64) {
	var withdrawAmt float64
	fmt.Print("How much would you like to withdraw? ")
	fmt.Scan(&withdrawAmt)

	if initialAmt <= 0 {
		fmt.Println("Your account has insufficient balance to withdraw money. Please try again after depositing money.")
		return
	}

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
