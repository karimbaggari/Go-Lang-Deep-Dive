// Example of function expressions

package main

import "fmt"

// Function expression assigned to a variable
func main() {
    add := func(x, y int) int {
        return x + y
    }

    result := add(3, 4) // Call the function expression
    fmt.Println(result) // Output: 7
} 