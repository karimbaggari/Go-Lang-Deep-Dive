package main

import "fmt"

// Using type alias for underlying types
type IntAlias int

func PrintAliasValue(val IntAlias) {
	fmt.Println("Value:", val)
}

func main() {
	var x IntAlias = 10
	PrintAliasValue(x)
}
