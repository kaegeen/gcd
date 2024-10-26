package main

import "fmt"

// Function to calculate the GCD of two numbers
func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func main() {
	var num1, num2 int

	// Prompt the user to enter two numbers
	fmt.Print("Enter the first number: ")
	fmt.Scan(&num1)
	fmt.Print("Enter the second number: ")
	fmt.Scan(&num2)

	// Calculate and display the GCD
	result := gcd(num1, num2)
	fmt.Printf("The GCD of %d and %d is %d\n", num1, num2, result)
}
