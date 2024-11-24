// Example of Maps in Go

package main

import "fmt"

func main() {
	// Declare and initialize a map
	ages := map[string]int{
		"Alice": 25,
		"Bob":   30,
	}

	// Access elements
	fmt.Println("Alice's Age:", ages["Alice"])

	// Add or update elements
	ages["Charlie"] = 35
	ages["Alice"] = 26
	fmt.Println("Updated Map:", ages)
}
