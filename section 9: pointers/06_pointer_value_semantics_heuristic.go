
  package main

  import "fmt"

  type Person struct {
      Name string
      Age  int
  }

  func updateAge(p *Person, newAge int) {
      p.Age = newAge
  }

  func main() {
      john := Person{Name: "John", Age: 30}
      fmt.Println("Before:", john.Age) // Output: Before: 30
      updateAge(&john, 31) // Passing pointer to modify original
      fmt.Println("After:", john.Age) // Output: After: 31
  }