package main

import (
    "errors"
    "fmt"
)

func main() {
    err := doSomething()
    if err != nil {
        fmt.Println("Error:", err)
    }

    // Example of using errors.Is
    if errors.Is(err, errors.New("an example error occurred")) {
        fmt.Println("Caught the specific error!")
    }
}

func doSomething() error {
    return errors.New("an example error occurred")
}