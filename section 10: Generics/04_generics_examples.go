package main

import (
	"fmt"
)

// Define a constraint: T must be a type that supports the `+` operator (i.e., number types)
type Adder interface {
	int | int32 | int64 | float32 | float64 // Only numeric types are allowed
}

// A generic function that adds two values of type T, where T must satisfy the Adder interface
func Add[T Adder](a T, b T) T {
	fmt.Println("Adding", a, "and", b)
	return a + b
}

func main() {
	// Using the Add function with integers
	intResult := Add(10, 20)
	fmt.Println("Result of adding integers:", intResult)

	// Using the Add function with floating point numbers
	floatResult := Add(10.5, 20.3)
	fmt.Println("Result of adding floats:", floatResult)

	// Uncommenting the following line would cause a compile-time error because strings are not part of the Adder interface
	// stringResult := Add("Hello", "World")
	// fmt.Println(stringResult)
}
