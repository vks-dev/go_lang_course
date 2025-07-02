package main

import "fmt"

func main() {
	var n1, n2 int
	fmt.Print("Enter first number: ")
	fmt.Scan(&n1)

	fmt.Print("Enter second number: ")
	fmt.Scan(&n2)

	sum := n1 + n2
	fmt.Printf("The sum of %d and %d is: %d\n", n1, n2, sum)
}