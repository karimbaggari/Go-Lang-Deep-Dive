package main

import (
	"fmt"
	"sync"
)

// printMessage prints a message and signals the WaitGroup when done
func printMessage(msg string, wg *sync.WaitGroup) {
	defer wg.Done() // Decrement the counter when the goroutine completes
	fmt.Println(msg) // Print the message
}

func main() {
	var wg sync.WaitGroup // Create a WaitGroup to manage goroutines
	wg.Add(2) // Set the number of goroutines to wait for
	go printMessage("Hello", &wg) // Start the first goroutine
	go printMessage("World", &wg) // Start the second goroutine
	wg.Wait() // Wait for all registered goroutines to finish
}