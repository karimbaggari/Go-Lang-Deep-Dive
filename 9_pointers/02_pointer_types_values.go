package main

import "fmt"

func main() {
    x := 100
    p := &x

    fmt.Printf("Type of p: %T\n", p)
    fmt.Println("Value at p:", *p) // Dereferencing
}
