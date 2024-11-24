package main

import "fmt"

func main() {
	// Array
	arr := [3]int{1, 2, 3}

	// Slice
	slc := []int{4, 5, 6}

	// Map
	m := map[string]int{"one": 1, "two": 2}

	// Struct
	type Person struct {
		Name string
		Age  int
	}
	p := Person{Name: "Alice", Age: 25}

	fmt.Println("Array:", arr)
	fmt.Println("Slice:", slc)
	fmt.Println("Map:", m)
	fmt.Println("Struct:", p)
}
