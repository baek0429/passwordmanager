package client

import (
	"testing"
)

func TestEncoding(t *testing.T) {
	decrypted := &DecryptedPassword{
		key:   "hello",
		value: "world",
	}
	encrypted := decrypted.SimpleEncrypt()
	t.Log(encrypted)

	debyte, _ := encrypted.SimpleDecrypt()
	t.Log(debyte)
}
