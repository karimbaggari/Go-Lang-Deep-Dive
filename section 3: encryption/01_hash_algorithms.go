// Hash Algorithms Example
// Demonstrates creating cryptographic hashes using the `crypto` package.

package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
)

func main() {
	data := "GoEncryptionExample"

	// SHA-256 Hash
	sha256Hash := sha256.Sum256([]byte(data))
	fmt.Printf("SHA-256: %x\n", sha256Hash)

	// SHA-512 Hash
	sha512Hash := sha512.Sum512([]byte(data))
	fmt.Printf("SHA-512: %x\n", sha512Hash)

	// Comparing hashes
	// fmt.Println("Hashes match:", sha256Hash == sha512Hash) // This will print false
}
