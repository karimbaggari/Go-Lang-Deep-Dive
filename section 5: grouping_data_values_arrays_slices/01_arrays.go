// Example of Arrays in Go

package main

import "fmt"

func main() {
	// Declare and initialize an array
	var arr [5]int
	arr[0] = 10
	arr[1] = 20

	arr2 := [5]int{10, 20}
	fmt.Println("Array:", arr)
	fmt.Println("Array 2:", arr2)

	// Declare and initialize with composite literals
	names := [3]string{"Alice", "Bob", "Charlie"}
	fmt.Println("Names:", names)

	// Iterate over an array
	for i, v := range names {
		fmt.Printf("Index: %d, Value: %s\n", i, v)
	}
}
