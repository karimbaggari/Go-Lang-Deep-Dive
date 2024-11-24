// Example of Nested Loops

package main

import "fmt"

func main() {
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 2; j++ {
			fmt.Printf("Outer: %d, Inner: %d\n", i, j)
		}
	}
}
