// Interfaces and Polymorphism

package main

import "fmt"

// Interface Definition
type Shape interface {
	Area() float64
}

// Structs Implementing the Interface
type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

type Rectangle struct {
	Width, Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func main() {
	shapes := []Shape{
		Circle{Radius: 5},
		Rectangle{Width: 10, Height: 5},
	}

	for _, shape := range shapes {
		fmt.Println("Area:", shape.Area())
	}
}
