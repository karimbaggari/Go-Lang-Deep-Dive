// Anonymous Functions

package main

import "fmt"

func main() {
	// Inline Anonymous Function
	func(msg string) {
		fmt.Println(msg)
	}("Hello, Go!")

	// Assigning Anonymous Function to a Variable
	add := func(a, b int) int {
		return a + b
	}

	fmt.Println("Sum:", add(5, 10))
}
