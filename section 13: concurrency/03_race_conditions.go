package main

import (
	"fmt"
	"sync"
)

var counter = 0
var mu sync.Mutex

func increment() {
	mu.Lock()
	defer mu.Unlock()
	counter++
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		for i := 0; i < 1000; i++ {
			increment()
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 1000; i++ {
			increment()
		}
	}()

	wg.Wait()
	fmt.Println("Final Counter:", counter)
}
