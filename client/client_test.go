package client

import (
	"testing"
)

func TestEncoding(t *testing.T) {
	decrypted := &DecryptedPassword{
		Key:   "hello",
		Value: "world",
	}
	encrypted := decrypted.SimpleEncrypt()
	t.Log(encrypted)

	debyte := encrypted.SimpleDecrypt()
	t.Log(debyte)
}
