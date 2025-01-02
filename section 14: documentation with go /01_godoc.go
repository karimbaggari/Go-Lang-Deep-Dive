package main

import "fmt"

// PrintMessage prints a given message.
// Example:
//     PrintMessage("Hello, World!")
// 
// This function demonstrates how to print a message to the console.
// It takes a single argument, msg, which is the message to be printed.
func PrintMessage(msg string) {
	fmt.Println(msg)
}

// main function is the entry point of the program
func main() {
	PrintMessage("Hello, World!") // Call the PrintMessage function
}