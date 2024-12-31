// Error Handling Best Practices

package main

import (
	"errors"
	"fmt"
	"log"
)

// Returning and Logging Errors
func connectToDB(connectionString string) error {
	if connectionString == "" {
		return errors.New("connection string is empty")
	}
	return nil
}

func main() {
	err := connectToDB("")
	if err != nil {
		log.Printf("Error: %v", err)
		// Optionally exit or return
	}
	fmt.Println("Application continues gracefully.")
}
