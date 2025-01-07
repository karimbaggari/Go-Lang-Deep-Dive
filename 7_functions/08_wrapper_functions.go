// Wrapper Functions

package main

import "fmt"

// Wrapper Function
func logger(f func(a, b int) int) func(a, b int) int {
	return func(a, b int) int {
		fmt.Println("Input Values:", a, b)
		result := f(a, b)
		fmt.Println("Result:", result)
		return result
	}
}

func add(a, b int) int {
	return a + b
}

func main() {
	loggedAdd := logger(add)
	loggedAdd(5, 10)
}
