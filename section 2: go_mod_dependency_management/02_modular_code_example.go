// Example of Modular Code
// Demonstrates package visibility and exported vs. unexported identifiers.

package main

import (
	"fmt"
	"mymodule/helpers"
	// another way to do it 
	// helpers "mymodule/helpers"
)

func main() {
	// Accessing the Exported Function
	result := helpers.Add(3, 5)
	fmt.Println("Sum:", result)

	// This will fail:
	// fmt.Println(helpers.subtract(10, 4)) // Unexported functions are not visible
	// ? check the helpers folder 
}
