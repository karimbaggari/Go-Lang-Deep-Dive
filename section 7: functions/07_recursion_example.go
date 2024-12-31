// Recursion Example

package main

import "fmt"

// Factorial Function
func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func main() {
	fmt.Println("Factorial of 5:", factorial(5))
}
