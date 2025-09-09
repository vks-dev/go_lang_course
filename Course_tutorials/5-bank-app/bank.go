package main

import "fmt"

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

	if choice >= 0 && choice <= 4 {
		fmt.Printf("This is your choice: %v\n", choice)
	} else {
		fmt.Printf("Your choice: %v is invalid, please try again.\n", choice)
	}

}
