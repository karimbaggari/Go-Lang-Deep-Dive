// Example of for range Loops
package main

import "fmt"

func main() {
	// Iterating over a slice
	numbers := []int{10, 20, 30}
	for index, value := range numbers {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}

	// Iterating over a map
	mapping := map[string]int{"Alice": 25, "Bob": 30}
	for key, value := range mapping {
		fmt.Printf("Key: %s, Value: %d\n", key, value)
	}

	// Iterating over a string
	str := "GoLang"
	for index, char := range str {
		fmt.Printf("Index: %d, Char: %c\n", index, char)
	}
}
