// Example of Switch Statements

package main

import "fmt"

func main() {
	day := "Tuesday"

	switch day {
	case "Monday", "Tuesday", "Wednesday":
		fmt.Println("It's a weekday.")
	case "Saturday", "Sunday":
		fmt.Println("It's a weekend.")
	default:
		fmt.Println("Invalid day.")
	}

	// Switch without a condition
	number := 15
	switch {
	case number%2 == 0:
		fmt.Println("The number is even.")
	case number%2 != 0:
		fmt.Println("The number is odd.")
	}
}
