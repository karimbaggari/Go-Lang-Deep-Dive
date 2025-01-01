// Example of Nested Loops

package main

import "fmt"

func main() {
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 2; j++ {
			if j == 1 {
				continue
			}
			if i == 2 {
				break
			}
			fmt.Printf("Outer: %d, Inner: %d\n", i, j)
		}
	}
}
