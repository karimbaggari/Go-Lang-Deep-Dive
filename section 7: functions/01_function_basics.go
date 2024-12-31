// Function Basics

package main

import "fmt"

// Simple Function
func add(a int, b int) int {
	return a + b
}

// Function with Multiple Return Values
func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero")
	}
	return a / b, nil
}

func main() {
	// Calling Simple Function
	sum := add(10, 20)
	fmt.Println("Sum:", sum)

	// Calling Function with Multiple Returns
	quotient, err := divide(10, 2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Quotient:", quotient)
	}
}
