// Higher-Order Functions

package main

import "fmt"

// Higher-Order Function
func applyOperation(a, b int, op func(int, int) int) int {
	return op(a, b)
}

func main() {
	add := func(x, y int) int { return x + y }
	multiply := func(x, y int) int { return x * y }

	fmt.Println("Addition:", applyOperation(2, 3, add))
	fmt.Println("Multiplication:", applyOperation(2, 3, multiply))
}
