package main

import "fmt"

func main() {
	// Zero Values
	var i int
	var f float64
	var b bool
	var s string

	fmt.Println("Integer zero value:", i)
	fmt.Println("Float zero value:", f)
	fmt.Println("Boolean zero value:", b)
	fmt.Println("String zero value:", s)

	// Blank Identifier
	quotient, _ := divide(10, 3) // Discard remainder
	fmt.Println("Quotient:", quotient)
}

// This function takes two integers, a and b, and returns their quotient and remainder.
func divide(a, b int) (int, int) {
	return a / b, a % b
}
