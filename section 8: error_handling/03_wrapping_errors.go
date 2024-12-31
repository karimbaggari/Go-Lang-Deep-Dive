// Wrapping Errors

package main

import (
	"errors"
	"fmt"
)

// Function Wrapping Errors
func readFile(filename string) error {
	return fmt.Errorf("failed to open file %s: %w", filename, errors.New("file not found"))
}

func main() {
	err := readFile("data.txt")
	if err != nil {
		fmt.Println("Error:", err)

		// Unwrapping the Error
		if errors.Is(err, errors.New("file not found")) {
			fmt.Println("The file does not exist.")
		}
	}
}
