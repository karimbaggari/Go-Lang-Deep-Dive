package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	data := "hello"
	hash := sha256.Sum256([]byte(data))
	fmt.Printf("SHA256 Hash: %x\n", hash)
}
