package aes

import (
	aescipher "crypto/aes"
	"crypto/cipher"
	"crypto/rand"

	//	"encoding/hex"
	"io"
)

// Encrypting []bytes to AES256 using provided Key (make sure you save Key somewhere if you want to decrypt)
func Encrypt(text []byte, key []byte) ([]byte, error) {
	// []byte("")
	//key, _ := hex.DecodeString(keyString)
	//plaintext := []byte(text)

	//Create a new Cipher Block from the key
	block, err := aescipher.NewCipher(key)
	if err != nil {
		return nil, err
	}

	//Create a new GCM - https://en.wikipedia.org/wiki/Galois/Counter_Mode
	//https://golang.org/pkg/crypto/cipher/#NewGCM
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	//Create a nonce. Nonce should be from GCM
	nonce := make([]byte, aesGCM.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}

	//Encrypt the data using aesGCM.Seal
	//Since we don't want to save the nonce somewhere else in this case, we add it as a prefix to the encrypted data. The first nonce argument in Seal is the prefix.
	encrypted_text := aesGCM.Seal(nonce, nonce, text, nil)
	
	return encrypted_text, nil
}