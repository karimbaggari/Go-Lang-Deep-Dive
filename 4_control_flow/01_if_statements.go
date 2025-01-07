// Example of if Statements and Comparisons

package main

import "fmt"

func main() {
	number := 15

	if number%2 == 0 {
		fmt.Println("The number is even.")
	} else {
		fmt.Println("The number is odd.")
	}

	// Using an if initializer
	if value := number * 2; value > 20 {
		fmt.Println("The value is greater than 20.")
	}
}
