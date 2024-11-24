// Deleting Elements from a Map

package main

import "fmt"

func main() {
	// Declare and initialize a map
	products := map[string]float64{
		"Laptop":  999.99,
		"Tablet":  499.99,
		"Phone":   299.99,
	}

	// Delete an element
	delete(products, "Tablet")
	fmt.Println("After Deletion:", products)
}
