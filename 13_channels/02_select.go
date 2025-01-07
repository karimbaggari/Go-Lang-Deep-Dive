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

	// Wait for the first message and exit
	select {
	case msg2 := <-ch2:
		fmt.Println("Received:", msg2)
	case msg1 := <-ch1:
		fmt.Println("Received:", msg1)
	}
    }


