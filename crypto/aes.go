package aescrypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"os"
)

var gcm cipher.AEAD

func Init() {
	aes, err := aes.NewCipher([]byte(os.Getenv("AES_KEY")))
	if err != nil {
		panic(err)
	}

	gcm, err = cipher.NewGCM(aes)
	if err != nil {
		panic(err)
	}
}

func EncryptField(field string) string {
	nonce := make([]byte, gcm.NonceSize())
	_, err := rand.Read(nonce)
	if err != nil {
		panic(err)
	}

	encryptedField := gcm.Seal(nonce, nonce, []byte(field), nil)

	encodedField := base64.StdEncoding.EncodeToString(encryptedField)

	return encodedField
}

func DecryptField(field string) string {
	decodedField, err := base64.StdEncoding.DecodeString(field)
	if err != nil {
		panic(err)
	}

	nonceSize := gcm.NonceSize()
	nonce, ciphertext := decodedField[:nonceSize], decodedField[nonceSize:]

	decryptedField, err := gcm.Open(nil, []byte(nonce), []byte(ciphertext), nil)
	if err != nil {
		panic(err)
	}

	return string(decryptedField)
}
