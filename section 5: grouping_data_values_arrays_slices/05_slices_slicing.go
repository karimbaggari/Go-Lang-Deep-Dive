// Example of Slicing a Slice

package main

import "fmt"

func main() {
	slice := []int{10, 20, 30, 40, 50}

	// Slice operations
	fmt.Println("Original Slice:", slice)
	fmt.Println("slice[1:4]:", slice[1:4]) // start from one and stop before 4 index
	fmt.Println("slice[:3]:", slice[:3]) // print everything and stop before 3 index
	fmt.Println("slice[2:]:", slice[2:]) // start from index 2 and print everything
}
