package main

import "fmt"

func main() {
	var n int16

	fmt.Print("Enter a number: ")
	fmt.Scan(&n)

	if n%2 == 0 {
		fmt.Printf("%v is an even number.\n", n)
	} else {
		fmt.Printf("%v is an odd number.\n", n)
	}
}