
  package main

  import "fmt"

  type Person struct {
      Name string
      Age  int
  }


// Heuristics for Pointer vs. Value Usage
// Use Pointers When:
// - Modifying Data: Pointers allow you to change the original data without creating a copy.
// - Large Structs: Passing pointers is more efficient for large structs to avoid the overhead of copying.
// - Nil Values: Pointers can represent the absence of a value (nil), which is useful for optional data.
//
// Use Values When:
// - Small Data Types: For small data types (like int, float, etc.), passing by value is usually more efficient.
// - Immutability: If you want to ensure that the data cannot be modified, use values.
// - Simplicity: Using values can make the code easier to understand and maintain, as it avoids pointer dereferencing.

  func updateAge(p *Person, newAge int) {
      p.Age = newAge
  }

  func main() {
      john := Person{Name: "John", Age: 30}
      fmt.Println("Before:", john.Age) // Output: Before: 30
      updateAge(&john, 31) // Passing pointer to modify original
      fmt.Println("After:", john.Age) // Output: After: 31
  }





