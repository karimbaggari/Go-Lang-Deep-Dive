package main

import "fmt"

func passByValue(x int) {
    x = 100
}

func passByReference(x *int) {
    *x = 100
}

func main() {
    n := 10

    passByValue(n)
    fmt.Println("After passByValue, n:", n)

    passByReference(&n)
    fmt.Println("After passByReference, n:", n)
}
