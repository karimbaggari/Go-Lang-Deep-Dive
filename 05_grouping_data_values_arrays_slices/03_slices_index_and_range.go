// Accessing Values by Index and Iterating with for range

package main

import "fmt"

func main() {
	// Declare a slice
	slice := []string{"Go", "Python", "Java"}

	// Access elements by index
	fmt.Println("First Element:", slice[0])

	// Iterate with for range
	for index, value := range slice {
		fmt.Printf("Index: %d, Value: %s\n", index, value)
	}
}
