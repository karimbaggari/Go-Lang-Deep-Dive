package main

import "fmt"

var globalVar = "I am global"

func main() {
	// Type Conversion
	var a int = 10
	var b float64 = 20.5
	c := float64(a) + b

	fmt.Printf("Sum: %.2f\n", c)
	fmt.Printf("Type of c: %T\n", c)

	// Scope
	localVar := "I am local"
	fmt.Println(globalVar)
	fmt.Println(localVar)
}
