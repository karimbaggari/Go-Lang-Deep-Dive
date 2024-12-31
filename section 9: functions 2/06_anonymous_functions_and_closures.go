// Anonymous Functions and Closures

package main

import "fmt"

func main() {
	// Anonymous Function
	add := func(a, b int) int {
		return a + b
	}
	fmt.Println("Sum:", add(3, 4))

	// Closure Example
	counter := func() func() int {
		count := 0
		return func() int {
			count++
			return count
		}
	}()

	fmt.Println(counter())
	fmt.Println(counter())
}
