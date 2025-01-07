// Wrapping Errors

package main

import (
	"errors"
	"fmt"
)

// Custom Error Definition
type FileError struct {
	Filename string
	Err      error
}

func (e *FileError) Error() string {
	return fmt.Sprintf("file error: %s: %v", e.Filename, e.Err)
}

// Function Wrapping Errors
func readFile(filename string) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", &FileError{Filename: filename, Err: err} // Use custom error
	}
	defer file.Close()

	return "", nil
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
