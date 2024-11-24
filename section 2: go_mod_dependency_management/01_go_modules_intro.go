// Go Modules Introduction
// This file demonstrates creating a new module and importing external packages.
/* 
	go mod init mymodule
	go get github.com/google/uuid
	go run 01_go_modules_intro.go
*/
package main

import (
	"fmt"
	"github.com/google/uuid" // Example of an external dependency
)

func main() {
	// Generating a UUID using the google/uuid package
	newUUID := uuid.New()
	fmt.Println("Generated UUID:", newUUID)
}
