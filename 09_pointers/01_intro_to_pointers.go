package main

import "fmt"

func main() {
    a := 42
    p := &a // Pointer to 'a'

    fmt.Println("Value of a:", a)
    fmt.Println("Address of a:", p)
}
