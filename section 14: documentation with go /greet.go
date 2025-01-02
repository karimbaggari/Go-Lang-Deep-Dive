// greet.go
// Package greet provides a simple greeting function.
package greet

import "fmt"

// Greet returns a greeting message with the provided name.
// It takes a string name and returns a greeting message as a string.
// execute go doc greet.Greet to see the documentation
func Greet(name string) string {
    return fmt.Sprintf("Hello, %s!", name)
}
