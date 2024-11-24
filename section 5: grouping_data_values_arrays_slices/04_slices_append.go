// Example of Appending to a Slice

package main

import "fmt"

func main() {
	// Initialize a slice
	slice := []int{1, 2, 3}

	// Append elements
	slice = append(slice, 4, 5)
	fmt.Println("After Append:", slice)

	// Append another slice
	extra := []int{6, 7}
	slice = append(slice, extra...) // Note this ...extra isnt allowed ur not in javascript!
	fmt.Println("After Appending Another Slice:", slice)
}
