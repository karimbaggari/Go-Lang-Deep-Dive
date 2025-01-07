// Understanding Slice Internals and Underlying Arrays
package main

import "fmt"

func main() {
	original := []int{10, 20, 30}
	sub := original[:2]

	// Modify the original
	original[1] = 99

	fmt.Println("Original Slice:", original)
	fmt.Println("Sub Slice:", sub)
}
