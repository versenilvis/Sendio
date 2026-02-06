// Package crypto - AES-256-GCM encryption and decryption
// encrypt sensitive data before storing in database (e.g., smtp config)
package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
)

var (
	ErrInvalidKeySize = errors.New("invalid key size: must be 32 bytes for AES-256")
	ErrDecrypt        = errors.New("decryption failed: check your key or ciphertext")
)

// Encrypt encrypts plain text string using AES-GCM with the provided 32-byte key
// returns base64 encoded string containing [number used once (nonce)][ciphertext]
// e.g.
// plaintext: mypass123
// key: 12345678901234567890123456789012
// with random.Reader, OS will give us 12 bytes of random data like:
// [A1, B2, C3, D4, E5, F6, G7, H8, I9, J10, K11, L12]
// mypass123 + Key + Nonce = [X9, Z8, Y7...]
// [A1, B2, C3, D4, ...] + [X9, Z8, Y7...]
// then base64 encode it: "oWV8mY...qW7p9"
func Encrypt(plaintext string, key string) (string, error) {
	keyBytes := []byte(key)
	if len(keyBytes) != 32 {
		return "", ErrInvalidKeySize
	}

	block, err := aes.NewCipher(keyBytes)
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	// random nonce
	// rand.Reader from crypto/rand connect to the random number generator from OS (CSPRNG)
	// it differs from math/rand
	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	// gcm.Seal will do 2 jobs:
	// encrypt the data + append the nonce to the data
	ciphertext := gcm.Seal(nonce, nonce, []byte(plaintext), nil)
	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

// Decrypt decrypts a base64 encoded ciphertext string using AES-GCM with the provided 32-byte key
func Decrypt(cryptoText string, key string) (string, error) {
	keyBytes := []byte(key)
	if len(keyBytes) != 32 {
		return "", ErrInvalidKeySize
	}

	ciphertext, err := base64.StdEncoding.DecodeString(cryptoText)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(keyBytes)
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonceSize := gcm.NonceSize()
	if len(ciphertext) < nonceSize {
		return "", fmt.Errorf("%w: ciphertext too short", ErrDecrypt)
	}

	// it's often 12 bytes
	// first 12 bytes is nonce
	// rest is ciphertext
	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]
	// convert to plain text
	// check if it was changed
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return "", fmt.Errorf("%w: %v", ErrDecrypt, err)
	}

	return string(plaintext), nil
}
