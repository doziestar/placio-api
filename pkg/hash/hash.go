package hash

import (
	"context"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"io"
	"placio-pkg/logger"
)

func encrypt(plainText []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}

	cipherText := gcm.Seal(nonce, nonce, plainText, nil)
	return cipherText, nil
}

func decrypt(cipherText []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonceSize := gcm.NonceSize()
	if len(cipherText) < nonceSize {
		return nil, errors.New("ciphertext too short")
	}

	nonce, cipherText := cipherText[:nonceSize], cipherText[nonceSize:]
	plainText, err := gcm.Open(nil, nonce, cipherText, nil)
	if err != nil {
		return nil, err
	}

	return plainText, nil

}

// EncryptString encrypts a string using the provided key and returns the encrypted result as a base64 encoded string.
func EncryptString(plainText string, key string) (string, error) {
	plainBytes := []byte(plainText)
	keyBytes := []byte(key)
	logger.Info(context.Background(), "encrypting string: "+plainText)
	cipherBytes, err := encrypt(plainBytes, keyBytes)
	if err != nil {
		logger.Error(context.Background(), err.Error())
		return "", err
	}
	logger.Info(context.Background(), "encrypted string: "+string(cipherBytes))

	return base64.StdEncoding.EncodeToString(cipherBytes), nil
}

// DecryptString decrypts a base64 encoded string using the provided key and returns the decrypted result as a string.
func DecryptString(cipherText string, key string) (string, error) {
	cipherBytes, err := base64.StdEncoding.DecodeString(cipherText)
	if err != nil {
		return "", err
	}
	keyBytes := []byte(key)
	plainBytes, err := decrypt(cipherBytes, keyBytes)
	if err != nil {
		return "", err
	}

	return string(plainBytes), nil

}
