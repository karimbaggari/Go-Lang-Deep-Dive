// Variadic Functions

package main

import "fmt"

// Function with Variadic Parameters
func sum(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func main() {
	fmt.Println("Sum of 1, 2, 3:", sum(1, 2, 3))
	fmt.Println("Sum of 4, 5, 6, 7:", sum(4, 5, 6, 7))
}
