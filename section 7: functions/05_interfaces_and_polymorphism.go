// Interfaces and Polymorphism Example

package main

import "fmt"



/*
In Go, polymorphism allows different types to be treated as 
instances of a common interface, enabling flexible and reusable
code. In the provided example, the Engine interface defines
a method Start(), which is implemented by both ElectricEngine
and GasEngine. This allows instances of these types 
to be assigned to the Engine field in the Car struct,
enabling interchangeable use. When the Start() method 
is called on an Engine, the actual method executed depends
on the concrete type of the engine, demonstrating dynamic 
method resolution. This design promotes extensibility,
as new types can be added without modifying existing code,
enhancing maintainability and reducing complexity.
*/

// Engine interface with a single method
type Engine interface {
	Start() string // Method to start the engine
}

// ElectricEngine struct defined to implement the Engine interface
type ElectricEngine struct{} // Empty struct


// Start method for ElectricEngine
func (e ElectricEngine) Start() string {
	return "Electric engine started silently." // Message for electric engine
}

// GasEngine struct implementing the Engine interface
type GasEngine struct{}

// Start method for GasEngine
func (g GasEngine) Start() string {
	return "Gas engine started with a roar!" // Message for gas engine
}

// Car struct that uses an Engine
type Car struct {
	Brand  string // Brand of the car
	Engine Engine // Engine type
}

func main() {
	// Create instances of Car with different engines
	cars := []Car{
		{Brand: "Tesla", Engine: ElectricEngine{}}, // Car with electric engine
		{Brand: "Ford", Engine: GasEngine{}},       // Car with gas engine
	}

	// Iterate over cars and start their engines
	for _, car := range cars {
		fmt.Printf("%s: %s\n", car.Brand, car.Engine.Start()) // Call Start method for each engine
	}
}