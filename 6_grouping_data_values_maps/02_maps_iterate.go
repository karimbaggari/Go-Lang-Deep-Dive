// Iterating Over a Map
package main

import "fmt"

func main() {
	// Declare and initialize a map
	grades := map[string]int{
		"Math":    10,
		"Science": 20,
		"History": 30,
	}

	// Iterate with for range
	for subject, grade := range grades {
		fmt.Printf("Subject: %s, Grade: %d\n", subject, grade)
	}
}
