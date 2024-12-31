// Implementing the Stringer Interface

package main

import "fmt"

// Custom Type
type Person struct {
	Name string
	Age  int
}

// Stringer Interface Implementation
func (p Person) String() string {
	return fmt.Sprintf("%s is %d years old.", p.Name, p.Age)
}

func main() {
	p := Person{Name: "Alice", Age: 30}
	fmt.Println(p)
}
