// Methods in Go

package main

import "fmt"

// Defining a Struct
type Rectangle struct {
	Length, Width int
}

// Method with Value Receiver
func (r Rectangle) Area() int {
	return r.Length * r.Width
}

// Method with Pointer Receiver
func (r *Rectangle) Scale(factor int) {
	r.Length *= factor
	r.Width *= factor
}

func main() {
	rect := Rectangle{Length: 5, Width: 3}
	fmt.Println("Area:", rect.Area())
	rect.Scale(2)
	fmt.Println("Scaled Area:", rect.Area())
}
