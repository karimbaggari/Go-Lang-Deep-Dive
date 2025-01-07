// Using the Comma-OK Idiom with Maps
package main

import "fmt"

func main() {
	// Declare and initialize a map
	cities := map[string]string{
		"NY": "New York",
		"SF": "San Francisco",
	}

	// Check if a key exists
	if city, ok := cities["LA"]; ok {
		fmt.Println("City Found:", city)
	} else {
		fmt.Println("City Not Found")
	}
}
