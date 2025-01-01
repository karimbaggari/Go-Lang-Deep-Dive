# **Deep Dive into Go: Section 3 - Encryption**

## **1. Hash Algorithms**
- **File:** [01_hash_algorithms.go](encryption/01_hash_algorithms.go)  
- **Description:** This file demonstrates creating cryptographic hashes using the `crypto` package. It showcases the use of SHA-256 and SHA-512 algorithms to generate hash values for a given input string.

## **2. Symmetric Encryption**
- **File:** [02_symmetric_encryption.go](encryption/02_symmetric_encryption.go)  
- **Description:** This example illustrates AES encryption and decryption using a shared key. It demonstrates how to encrypt a plaintext message and then decrypt it back to its original form using the same key.

## **3. Asymmetric Encryption**
- **File:** [03_asymmetric_encryption.go](encryption/03_asymmetric_encryption.go)  
- **Description:** This file demonstrates RSA encryption and decryption using public and private keys. It shows how to generate RSA keys, encrypt a message with the public key, and decrypt it with the private key.

## **4. Practical Applications**
- **File:** [03_cryptography.go](applications/03_cryptography.go)  
- **Description:** This file provides practical examples of using cryptographic functions in Go, including generating SHA256 hashes for given data.

## **Conclusion**
Understanding encryption is crucial for securing data in applications. This section covers both symmetric and asymmetric encryption methods, along with hash algorithms, providing a solid foundation for implementing security in Go applications.