// Example of Deleting from a Slice

package main

import "fmt"

func main() {
	slice := []int{10, 20, 30, 40, 50}

	// Delete the element at index 2
	slice = append(slice[:2], slice[3:]...)
	fmt.Println("After Deletion:", slice)
}
