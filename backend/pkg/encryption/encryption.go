package encryption

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"github.com/thanhpk/randstr"
	"io"
)

const checksum = "ashdfjDHJF98dsh8ehdfj"

// Encrypt encrypts plain string with a secret key and returns encrypt string with generated key.
func Encrypt(data string) (string, string, error) {
	key := randstr.String(32)

	data = checksum + data

	cipherBlock, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", "", err
	}

	aead, err := cipher.NewGCM(cipherBlock)
	if err != nil {
		return "", "", err
	}

	nonce := make([]byte, aead.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return "", "", err
	}

	return base64.URLEncoding.EncodeToString(aead.Seal(nonce, nonce, []byte(data), nil)), key, nil
}

// Decrypt decrypts encrypt string with a secret key and returns plain string.
func Decrypt(key, data string) (string, error) {
	encryptData, err := base64.URLEncoding.DecodeString(data)
	if err != nil {
		return "", err
	}

	cipherBlock, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}

	aead, err := cipher.NewGCM(cipherBlock)
	if err != nil {
		return "", err
	}

	nonceSize := aead.NonceSize()
	if len(encryptData) < nonceSize {
		return "", err
	}

	nonce, cipherText := encryptData[:nonceSize], encryptData[nonceSize:]
	plainData, err := aead.Open(nil, nonce, cipherText, nil)
	if err != nil {
		return "", err
	}

	result := string(plainData)

	if len(result) < len(checksum) || result[:len(checksum)] != checksum {
		return "", errors.New("checksum check failed")
	}

	return result[len(checksum):], nil
}
