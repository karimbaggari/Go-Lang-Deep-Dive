package main

import "fmt"

// Generic stack example
type Stack[T any] struct {
	data []T
}

func (s *Stack[T]) Push(item T) {
	s.data = append(s.data, item)
}

func (s *Stack[T]) Pop() T {
	if len(s.data) == 0 {
		var zero T
		return zero
	}
	item := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return item
}

func main() {
	intStack := Stack[int]{}
	intStack.Push(10)
	intStack.Push(20)
	fmt.Println(intStack.Pop()) // 20
	fmt.Println(intStack.Pop()) // 10

	stringStack := Stack[string]{}
	stringStack.Push("Hello")
	stringStack.Push("World")
	fmt.Println(stringStack.Pop()) // World
	fmt.Println(stringStack.Pop()) // Hello
}
