// Example of a higher-order function

package main

import "fmt"

// Function that takes another function as an argument
func applyFunction(fn func(int) int, value int) int {
    return fn(value)
}

// A sample function to be passed
func square(x int) int {
    return x * x
}

// Usage of the higher-order function
func main() {
    result := applyFunction(square, 5) // Pass the square function
    fmt.Println(result) // Output: 25
} 