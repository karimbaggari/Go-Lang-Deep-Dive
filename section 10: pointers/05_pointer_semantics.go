package main

import "fmt"

type Rectangle struct {
    Width, Height int
}

func (r Rectangle) AreaByValue() int {
    return r.Width * r.Height
}

func (r *Rectangle) ScaleByPointer(factor int) {
    r.Width *= factor
    r.Height *= factor
}

func main() {
    rect := Rectangle{Width: 10, Height: 5}
    fmt.Println("Area (value):", rect.AreaByValue())

    rect.ScaleByPointer(2)
    fmt.Println("Scaled Rectangle:", rect)
}
