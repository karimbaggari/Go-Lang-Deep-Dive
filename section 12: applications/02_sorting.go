package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{3, 1, 4, 1, 5}
	sort.Ints(nums)
	fmt.Println("Sorted:", nums)

	// Custom sorting
	words := []string{"apple", "banana", "cherry"}
	sort.Sort(sort.Reverse(sort.StringSlice(words)))
	fmt.Println("Reverse Sorted:", words)
}
