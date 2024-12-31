package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generator(numbers []int, ch chan<- int) {
	defer close(ch)
	for _, n := range numbers {
		ch <- n
	}
}

func worker(id int, in <-chan int, out chan<- string) {
	for n := range in {
		time.Sleep(time.Millisecond * time.Duration(rand.Intn(100)))
		out <- fmt.Sprintf("Worker %d processed %d", id, n)
	}
}

func fanOut(in <-chan int, workers []chan string) {
	for n := range in {
		for _, w := range workers {
			select {
			case w <- fmt.Sprintf("Processed: %d", n):
			default:
			}
		}
	}
}

func main() {
	in := make(chan int)
	out1 := make(chan string)
	out2 := make(chan string)

	go generator([]int{1, 2, 3, 4, 5}, in)
	go worker(1, in, out1)
	go worker(2, in, out2)

	for i := 0; i < 5; i++ {
		select {
		case msg := <-out1:
			fmt.Println(msg)
		case msg := <-out2:
			fmt.Println(msg)
		}
	}
}
