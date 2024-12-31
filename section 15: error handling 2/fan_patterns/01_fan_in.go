package main

import (
	"fmt"
	"time"
)

func producer(name string, ch chan<- string) {
	for i := 1; i <= 5; i++ {
		ch <- fmt.Sprintf("%s: %d", name, i)
		time.Sleep(time.Millisecond * 500)
	}
	close(ch)
}

func fanIn(ch1, ch2 <-chan string, out chan<- string) {
	defer close(out)
	for ch1 != nil || ch2 != nil {
		select {
		case msg, ok := <-ch1:
			if ok {
				out <- msg
			} else {
				ch1 = nil
			}
		case msg, ok := <-ch2:
			if ok {
				out <- msg
			} else {
				ch2 = nil
			}
		}
	}
}

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)
	out := make(chan string)

	go producer("Producer1", ch1)
	go producer("Producer2", ch2)
	go fanIn(ch1, ch2, out)

	for msg := range out {
		fmt.Println(msg)
	}
}
