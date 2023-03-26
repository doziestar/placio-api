package hash

import "testing"

func TestEncryptDecryptString(t *testing.T) {
	key := "a1b2c3d4e5f6g7h8i9j0k1l2m3n4o5p6"
	plainText := "This is a sample text to be encrypted and decrypted."

	// Test encryption
	cipherText, err := EncryptString(plainText, key)
	if err != nil {
		t.Errorf("Encryption failed: %v", err)
	}

	// Test decryption
	decryptedText, err := DecryptString(cipherText, key)
	if err != nil {
		t.Errorf("Decryption failed: %v", err)
	}

	// Check if the decrypted text matches the original plaintext
	if decryptedText != plainText {
		t.Errorf("Decrypted text does not match the original plaintext: got %s, want %s", decryptedText, plainText)
	}
}
