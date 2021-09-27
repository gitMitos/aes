package aes

import (
	"crypto/rand"
	"encoding/hex"
)

// Random encryption key may not be what you want in order to be able decrypt text. Use with caution
func RandomKey() (string, error) {

	// Creating random 32 byte to use as a key for AES256
	cryptobytes := make([]byte, 32) 
	if _, err := rand.Read(cryptobytes); err != nil {
		return "", err
	}

	key := hex.EncodeToString( cryptobytes ) 

	return key, nil
}