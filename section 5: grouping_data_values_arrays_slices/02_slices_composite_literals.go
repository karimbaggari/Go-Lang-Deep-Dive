// Example of Slices with Composite Literals

package main

import "fmt"

func main() {
	// Declare and initialize a slice
	numbers := []int{10, 20, 30}
	fmt.Println("Slice:", numbers)

	// Slices are dynamic
	numbers = append(numbers, 40)
	fmt.Println("Updated Slice:", numbers)
}
