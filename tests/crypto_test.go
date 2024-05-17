package test

import (
	"os"
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	aescrypt "gocrypt-api/crypto"
	_ "gocrypt-api/routers"
)

// Testing if encryption/decryption works correctly
func TestEncryptAndDecryptField(t *testing.T) {
	os.Setenv("AES_KEY", "a very very very very secret key")
	aescrypt.Init()

	field := "teste"
	encryptedField := aescrypt.EncryptField(field)
	decryptedField := aescrypt.DecryptField(encryptedField)

	Convey("Subject: Test Encryption\n", t, func() {
		Convey("Encrypted field should be different than the plaintext", func() {
			So(encryptedField, ShouldNotEqual, field)
		})
		Convey("Decrypted field should be the same as the plaintext", func() {
			So(decryptedField, ShouldEqual, field)
		})
	})
}
