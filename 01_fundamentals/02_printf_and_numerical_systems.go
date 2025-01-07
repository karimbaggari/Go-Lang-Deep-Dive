package main

import "fmt"

func main() {
	num := 255

	// Print in Decimal, Hexadecimal, Binary
	fmt.Printf("Decimal: %d\n", num)
	fmt.Printf("Hexadecimal: %#x\n", num)
	fmt.Printf("Binary: %b\n", num)

	// Numerical Systems
	dec := 42
	bin := 0b101010 // Binary
	hex := 0x2A     // Hexadecimal

	fmt.Printf("Decimal: %d, Binary: %b, Hexadecimal: %#x\n", dec, bin, hex)
}
