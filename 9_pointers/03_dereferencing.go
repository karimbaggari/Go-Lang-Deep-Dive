package main

import "fmt"

func main() {
    var y int = 20
    var ptr *int = &y

    fmt.Println("Before modification, y:", y)
    *ptr = 50 // Modify value via pointer
    fmt.Println("After modification, y:", y)
}
