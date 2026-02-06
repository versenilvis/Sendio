// Package crypto - Crypto test
// Purpose of these tests:
// 1. Ensure data integrity after encryption and decryption cycle
// 2. Validate security constraints like key size requirements
// 3. Verify error handling for corrupted data or unauthorized keys
//
// Logic:
// - We use a mock 32-byte key and a secret string
// - We simulate various scenarios: happy path, invalid keys, and tampered data
//
// Expected Result:
// - Successful match for original and decrypted data
// - Explicit failure for any violation of security rules (wrong key or corrupted data)
package crypto

import (
	"testing"
)

func TestEncryptDecrypt(t *testing.T) {
	key := "0123456789abcdef0123456789abcdef" // 32 bytes
	originalText := "my-secret-smtp-password-123"

	// 1. Test successful encryption
	encrypted, err := Encrypt(originalText, key)
	if err != nil {
		t.Fatalf("Encrypt failed: %v", err)
	}
	if encrypted == "" {
		t.Fatal("Encrypted string is empty")
	}
	if encrypted == originalText {
		t.Fatal("Encrypted string is same as original text")
	}

	// 2. Test successful decryption
	decrypted, err := Decrypt(encrypted, key)
	if err != nil {
		t.Fatalf("Decrypt failed: %v", err)
	}
	if decrypted != originalText {
		t.Errorf("Decrypted text mismatch: got %v, want %v", decrypted, originalText)
	}

	// 3. Test with invalid key size
	shortKey := "too-short"
	_, err = Encrypt(originalText, shortKey)
	if err == nil || err != ErrInvalidKeySize {
		t.Errorf("Expected ErrInvalidKeySize for short key, got %v", err)
	}

	// 4. Test decryption with wrong key
	wrongKey := "wrong-32-byte-key-abcdefghijk-01"
	_, err = Decrypt(encrypted, wrongKey)
	if err == nil {
		t.Error("Expected error when decrypting with wrong key, got nil")
	}

	// 5. Test with corrupted ciphertext
	corrupted := encrypted[:len(encrypted)-5] + "XXXXX"
	_, err = Decrypt(corrupted, key)
	if err == nil {
		t.Error("Expected error for corrupted ciphertext, got nil")
	}
}
