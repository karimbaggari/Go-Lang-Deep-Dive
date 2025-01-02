package main

import "fmt"

// Generic function to print elements of any type
func PrintElements[T any](elements []T) {
    for _, element := range elements {
        fmt.Println(element)
    }
}

func main() {
    // Example with integers
    intSlice := []int{1, 2, 3, 4, 5}
    fmt.Println("Integer Slice:")
    PrintElements(intSlice)

    // Example with strings
    stringSlice := []string{"Hello", "Generics", "in", "Go"}
    fmt.Println("String Slice:")
    PrintElements(stringSlice)
}