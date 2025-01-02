
package main

import (
	"fmt"
	"sync"
)

var counter = 0
var mu sync.Mutex // Mutex to protect access to the counter

// increment safely increments the counter using a mutex
func increment() {
	mu.Lock() // Lock the mutex to ensure exclusive access to the counter
	defer mu.Unlock() // Ensure the mutex is unlocked when the function exits
	counter++ // Increment the shared counter
}

func main() {
	var wg sync.WaitGroup // WaitGroup to wait for goroutines to finish
	wg.Add(2) // Set the number of goroutines to wait for

	// First goroutine to increment the counter
	go func() {
		defer wg.Done() // Mark this goroutine as done when it finishes
		for i := 0; i < 1000; i++ {
			increment() // Call the increment function
		}
	}()

	// Second goroutine to increment the counter
	go func() {
		defer wg.Done() // Mark this goroutine as done when it finishes
		for i := 0; i < 1000; i++ {
			increment() // Call the increment function
		}
	}()

	wg.Wait() // Wait for all goroutines to finish
	fmt.Println("Final Counter:", counter) // Print the final value of the counter
}

// run with `go run -race 03_race_conditions.go` to see the race condition
// take out the mutex protection to see the race condition fail 


/*
The mutex protection prevents goroutines from accessing and modifying the shared counter variable simultaneously.
*/