// Example of Using make with Slices
package main

import "fmt"

func main() {
	// Create a slice with make
	slice := make([]int, 5, 10)
	fmt.Println("Slice:", slice)
	fmt.Println("Length:", len(slice), "Capacity:", cap(slice))
}

/* 
   Length = 5 (number of elements currently in the slice)
   Capacity = 10 (maximum number of elements the slice can hold before resizing)
*/