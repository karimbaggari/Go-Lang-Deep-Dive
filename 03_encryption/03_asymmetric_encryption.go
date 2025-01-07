// Asymmetric Encryption Example
// Demonstrates RSA encryption and decryption using public and private keys.
package main

import (
	"crypto/rand"
	"crypto/rsa"
	// "crypto/x509"
	// "encoding/pem"
	"fmt"
)

func main() {
	message := "Hello, RSA Encryption!"

	// Generate RSA keys
	privateKey, publicKey := generateKeys()

	// Encrypt message using public key
	encryptedMessage := encryptRSA(publicKey, message)
	fmt.Printf("Encrypted Message: %x\n", encryptedMessage)

	// Decrypt message using private key
	decryptedMessage := decryptRSA(privateKey, encryptedMessage)
	fmt.Printf("Decrypted Message: %s\n", decryptedMessage)
}

func generateKeys() (*rsa.PrivateKey, *rsa.PublicKey) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		panic(err)
	}
	return privateKey, &privateKey.PublicKey
}

func encryptRSA(publicKey *rsa.PublicKey, message string) []byte {
	encryptedMessage, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey, []byte(message))
	if err != nil {
		panic(err)
	}
	return encryptedMessage
}

func decryptRSA(privateKey *rsa.PrivateKey, encryptedMessage []byte) string {
	decryptedMessage, err := rsa.DecryptPKCS1v15(rand.Reader, privateKey, encryptedMessage)
	if err != nil {
		panic(err)
	}
	return string(decryptedMessage)
}
