// Using go mod Commands
// This file explains various `go mod` commands.
package main

import "fmt"

func main() {
	fmt.Println("Manage your Go dependencies effectively with the following commands:")
	fmt.Println("1. Initialize a module: `go mod init <module-name>`")
	fmt.Println("2. Add dependencies: `go get <package>`")
	fmt.Println("3. Clean up unused dependencies: `go mod tidy`")
	fmt.Println("4. Verify dependencies: `go mod verify`")
	fmt.Println("5. List modules: `go list -m all`")
}
