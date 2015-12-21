package client

import (
	"encoding/base64"
	// "fmt"
)

// Password before encryption
type EncryptedPassword struct {
	Key   string
	Value string
}

// SimpleDecrypt accepts pointer receiver *EncryptedPassword and returns *DecryptedPassword
func (encrypted *EncryptedPassword) SimpleDecrypt() *DecryptedPassword {
	b, err := base64.StdEncoding.DecodeString(encrypted.Value)
	if err != nil {
		// fmt.Println(err)
	}
	var decrypted DecryptedPassword
	str := string(b[:])
	decrypted.Key = encrypted.Key
	decrypted.Value = string(str)
	return &decrypted
}

// Password before decryption
type DecryptedPassword struct {
	Key   string
	Value string
}

// SimpleDecrypt accepts pointer receiver *DecryptedPassword and returns *EncryptedPassword
func (notyet *DecryptedPassword) SimpleEncrypt() *EncryptedPassword {
	var encrypted EncryptedPassword

	b := base64.StdEncoding.EncodeToString([]byte(notyet.Value))
	encrypted.Key = notyet.Key
	encrypted.Value = string(b)
	return &encrypted
}

func (notyet *DecryptedPassword) String() string {
	return notyet.Key + " " + notyet.Value
}
