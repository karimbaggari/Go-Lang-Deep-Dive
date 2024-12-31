package main

import "fmt"

func main() {
    stackAllocation()
    heapAllocation()
}

func stackAllocation() {
    a := 42
    b := 99
    fmt.Println("Stack Allocation - Variables a and b:", a, b)
}

func heapAllocation() {
    c := new(int)
    *c = 50
    fmt.Println("Heap Allocation - Value at c:", *c)
}
