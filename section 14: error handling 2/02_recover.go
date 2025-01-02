package main

import "fmt"

func safeDivision(a, b int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()

	result := a / b
	fmt.Println("Result:", result)
}

func main() {
	safeDivision(10, 0) // Will cause a panic but will be recovered
	fmt.Println("Program continues...")
}
