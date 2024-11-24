// Symmetric Encryption Example
// Demonstrates AES encryption and decryption using a shared key.

package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
)

func main() {
	key := []byte("myverystrongpasswordo32bitlength")
	plaintext := "Hello, Go Encryption!"

	// Encrypt
	ciphertext, nonce := encryptAES(key, plaintext)
	fmt.Printf("Ciphertext: %s\n", ciphertext)

	// Decrypt
	decryptedText := decryptAES(key, ciphertext, nonce)
	fmt.Printf("Decrypted: %s\n", decryptedText)
}

// aes stands for advanced encryption standard
func encryptAES(key []byte, plaintext string) (string, []byte) {
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	nonce := make([]byte, 12) // AES-GCM standard nonce size
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err)
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		panic(err)
	}

	ciphertext := aesGCM.Seal(nil, nonce, []byte(plaintext), nil)
	return hex.EncodeToString(ciphertext), nonce
}

func decryptAES(key []byte, ciphertext string, nonce []byte) string {
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		panic(err)
	}

	ciphertextBytes, _ := hex.DecodeString(ciphertext)
	plaintext, err := aesGCM.Open(nil, nonce, ciphertextBytes, nil)
	if err != nil {
		panic(err)
	}

	return string(plaintext)
}
