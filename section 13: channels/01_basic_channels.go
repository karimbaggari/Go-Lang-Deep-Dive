package main

import (
    "fmt"
)

func sendMessage(ch chan string, msg string) {
    ch <- msg // Send the message to the channel
}

func main() {
    // Create two channels
    ch1 := make(chan string)
    ch2 := make(chan string)

    // Start two goroutines, each sending a message to its respective channel
    go sendMessage(ch1, "Message from channel 1")
    go sendMessage(ch2, "Message from channel 2")

    // Receive messages from both channels and log them
    msg1 := <-ch1
    fmt.Println("Received:", msg1)
    msg2 := <-ch2
    fmt.Println("Received:", msg2)
}

// this example will be more clear when you check the select.go file the completion is there