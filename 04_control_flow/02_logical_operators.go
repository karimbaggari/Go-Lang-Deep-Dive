// Example of Logical Operators
package main

import "fmt"

func main() {
	age := 25
	hasLicense := true

	// AND operator
	if age >= 18 && hasLicense {
		fmt.Println("You are eligible to drive.")
	}

	// OR operator
	if age < 18 || !hasLicense {
		fmt.Println("You are not eligible to drive.")
	}

	// NOT operator
	if !hasLicense {
		fmt.Println("You need a license to drive.")
	}
}
