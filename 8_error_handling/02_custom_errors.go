// Custom Errors

package main

import (
	"fmt"
)

// Custom Error Type
type DivisionError struct {
	Dividend int
	Divisor  int
}

func (e *DivisionError) Error() string {
	return fmt.Sprintf("cannot divide %d by %d", e.Dividend, e.Divisor)
}

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, &DivisionError{Dividend: a, Divisor: b}
	}
	return a / b, nil
}

func main() {
	_, err := divide(10, 0)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
