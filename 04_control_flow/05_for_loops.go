// Example of For Loops

package main

import "fmt"

// there isnt an actual while loop in go but there is alternatives

func main() {
	// Classic for loop
	for i := 0; i < 5; i++ {
		fmt.Println("Iteration:", i)
	}

	// While-like loop
	counter := 0
	for counter < 3 {
		fmt.Println("Counter:", counter)
		counter++
	}

	// Infinite loop
	for {
		fmt.Println("Infinite loop")
		break // Break to prevent infinite execution
	}
}
