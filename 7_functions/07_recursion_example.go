// Recursion Example

package main

import "fmt"


// Factorial Function
func factorial(n int) int {
	if n == 0 {
		return 1
	}
	result := n * factorial(n-1)
	fmt.Println("Calculating result for n =", n , "multiplied by factorial of", factorial(n-1)) // Log when the result calculation starts
	
	fmt.Println("Result of factorial(", n, "):", result) // Log the result
	return result
}

func main() {
	fmt.Println("Factorial of 3:", factorial(3))
}
