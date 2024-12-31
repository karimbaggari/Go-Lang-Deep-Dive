// Methods with Structs

package main

import "fmt"

type Rectangle struct {
	Width, Height float64
}

// Method with Receiver
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func main() {
	rect := Rectangle{Width: 10, Height: 5}
	fmt.Println("Area of Rectangle:", rect.Area())
}
