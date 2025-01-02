package main

import "fmt"

// Generic function with constraints
func Add[T int | float64](a, b T) T {
	return a + b
}

func main() {
	fmt.Println(Add(2, 3))       // int addition
	fmt.Println(Add(2.5, 3.5))   // float addition
}
