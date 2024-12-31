// Basic Functions

package main

import "fmt"

// Function Definition
func add(a int, b int) int {
	return a + b
}

// Named Return Values
func multiply(a, b int) (result int) {
	result = a * b
	return
}

func main() {
	fmt.Println("Sum:", add(2, 3))
	fmt.Println("Product:", multiply(4, 5))
}
