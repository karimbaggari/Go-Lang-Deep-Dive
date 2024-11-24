package main

import "fmt"

// Composition Example
type Engine struct {
	Horsepower int
}

type Car struct {
	Brand  string
	Engine // Embedded struct
}

func main() {
	myCar := Car{
		Brand:  "Tesla",
		Engine: Engine{Horsepower: 670},
	}

	fmt.Printf("Car: %s, Horsepower: %d\n", myCar.Brand, myCar.Horsepower)
}


// for the inheritance part, Sorry this isn't JAVA.