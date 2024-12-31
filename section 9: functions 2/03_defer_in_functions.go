// Using Defer in Functions

package main

import "fmt"

func printMessages() {
	defer fmt.Println("Goodbye!") // Executes last
	fmt.Println("Hello!")         // Executes first
}

func main() {
	printMessages()
}
