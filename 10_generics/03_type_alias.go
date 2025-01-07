package main

import "fmt"

// Defining a type alias for different numeric types
type NumberAlias interface {
	int | float64
}

// Creating a function that accepts NumberAlias
func PrintNumberValue[T NumberAlias](val T) {
	fmt.Println("Value:", val)
}

// Updating main function to demonstrate NumberAlias with different types
func main() {
	var intValue NumberAlias = 10
	var floatValue NumberAlias = 10.5

	PrintNumberValue(intValue)
	PrintNumberValue(floatValue)
}