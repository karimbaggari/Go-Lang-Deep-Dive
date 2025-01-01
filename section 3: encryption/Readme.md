# **Deep Dive into Go: Section 3 - Encryption**

## **1. Hash Algorithms**
- **File:** [01_hash_algorithms.go](encryption/01_hash_algorithms.go)  
- **Description:** This file demonstrates creating cryptographic hashes using the `crypto` package. It showcases the use of SHA-256 and SHA-512 algorithms to generate hash values for a given input string. Hashing is useful when you want to represent data in a fixed-size format and ensure data integrity.

## **2. Symmetric Encryption**
- **File:** [02_symmetric_encryption.go](encryption/02_symmetric_encryption.go)  
- **Description:** This example illustrates AES encryption and decryption using a shared key. It demonstrates how to encrypt a plaintext message and then decrypt it back to its original form using the same key. Symmetric encryption is fast and efficient, but the challenge lies in securely sharing the key between parties.

## **3. Asymmetric Encryption**
- **File:** [03_asymmetric_encryption.go](encryption/03_asymmetric_encryption.go)  
- **Description:** This file demonstrates RSA encryption and decryption using public and private keys. It shows how to generate RSA keys, encrypt a message with the public key, and decrypt it with the private key. Asymmetric encryption is ideal for securely transmitting data between parties without needing to share a secret key.

## **Conclusion**
Understanding encryption is crucial for securing data in applications. This section covers both symmetric and asymmetric encryption methods, along with hash algorithms, providing a solid foundation for implementing security in Go applications. It will help you ensure data confidentiality, integrity, and authenticity in your Go projects.
