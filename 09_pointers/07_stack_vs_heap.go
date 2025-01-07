package main

import "fmt"

func stackExample() {
    // `x` is allocated on the stack because it's a local variable.
    x := 42
    fmt.Println("Stack Example - x:", x)
    // `x` disappears when stackExample ends.
}

func heapExample() *int {
    num := new(int) // Allocate memory on the heap
    *num = 42       // Assign a value
    return num      // Return the pointer
}

func main() {
    fmt.Println("Stack Example:")
    stackExample()

    heap := heapExample()
    fmt.Println("Heap Example - *heap:", *heap)
    *heap = 43
    fmt.Println("Heap Example - *heap:", *heap)
}



/* 
Stack Memory
What is it?

The stack is a tightly managed, linear region of memory used for short-term storage.
It operates in a "Last In, First Out" (LIFO) manner, like stacking and unstacking plates.
What is allocated on the stack?

Local variables and function call data (like function parameters).
Data stored here is temporary and only exists while the function is running.
How does it work?

When a function is called, space is reserved on the stack for its variables.
When the function ends, this space is released automatically.
Advantages:

Fast memory allocation and deallocation.
No need for garbage collection.
Limitations:

Limited size: If you exceed the stack's size, you get a "stack overflow."
Data is lost as soon as the function exits.
Heap Memory
What is it?

The heap is a larger, more flexible area of memory used for long-term storage.
It is shared by the whole program.
What is allocated on the heap?

Data that must outlive the function it was created in.
Variables allocated with functions like new, make, or that involve pointers.
How does it work?

Data is manually allocated and garbage-collected when no longer needed.
Slower than stack allocation because the system must search for free space.
Advantages:

Can store large data that needs to persist longer.
Suitable for dynamic memory needs.
Limitations:

Slower allocation and deallocation.
Requires garbage collection to clean up unused memory.
*/