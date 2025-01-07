// helpers/helpers.go
package helpers


// to Export a function you put the first letter Capital 

// Exported Function
func Add(a, b int) int {
	return a + b
}

// Unexported Function
func subtract(a, b int) int {
	return a - b
}
