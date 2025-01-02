// Example of a closure that maintains state

package main

import "fmt"

func makeCounter() func() int {
    count := 0 // This variable is captured by the closure
    return func() int {
        count++ // The closure can modify the captured variable
        return count
    }
}

// Usage of the closure
func main() {
    counter := makeCounter() // Create a new counter instance
    fmt.Println(counter())   // Output: 1
    fmt.Println(counter())   // Output: 2
    // Additional calls can continue to increment the count
    fmt.Println(counter())   // Output: 3
} 