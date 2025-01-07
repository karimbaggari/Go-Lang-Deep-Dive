// Example of Select Statements for Concurrency

package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		ch1 <- "Data from channel 1"
	}()

	go func() {
		time.Sleep(1 * time.Second)
		ch2 <- "Data from channel 2"
	}()

	select {
	case msg1 := <-ch1:
		fmt.Println(msg1)
	case msg2 := <-ch2:
		fmt.Println(msg2)
	case <-time.After(3 * time.Second):
		fmt.Println("Timeout occurred.")
	default:
		fmt.Println("No messages received.")
	}
	
}
