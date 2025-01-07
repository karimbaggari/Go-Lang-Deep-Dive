// Using Defer

package main

import "fmt"

func main() {
	defer fmt.Println("This will print last, even though it's defined first.")
	fmt.Println("This will print before the deferred statement.")
}
