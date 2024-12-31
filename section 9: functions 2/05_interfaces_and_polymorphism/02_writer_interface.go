// Using the Writer Interface

package main

import (
	"fmt"
	"os"
)

// Writing to a File
func writeToFile(content string) {
	file, err := os.Create("output.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	file.WriteString(content)
}

func main() {
	writeToFile("Hello, Go!")
}
