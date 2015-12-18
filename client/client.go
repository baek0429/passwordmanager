package client

import (
	"encoding/base64"
	// "fmt"
)

type EncryptedPassword struct {
	key   string
	value string
}

func (encrypted *EncryptedPassword) SimpleDecrypt() *DecryptedPassword {
	b, err := base64.StdEncoding.DecodeString(encrypted.value)
	if err != nil {
		// fmt.Println(err)
	}
	var decrypted DecryptedPassword
	str := string(b[:])
	decrypted.Key = encrypted.key
	decrypted.Value = string(str)
	return &decrypted
}

type DecryptedPassword struct {
	Key   string
	Value string
}

func (notyet *DecryptedPassword) SimpleEncrypt() *EncryptedPassword {
	var encrypted EncryptedPassword

	b := base64.StdEncoding.EncodeToString([]byte(notyet.Value))
	encrypted.key = notyet.Key
	encrypted.value = string(b)
	return &encrypted
}
