package aes

import (
	"crypto/rand"
	"encoding/hex"
)

// Create random bytes
func RandomBytes(size int) ([]byte, error){
	// Creating random 32 byte to use as a key for AES256
	cryptobytes := make([]byte, size) 
	if _, err := rand.Read(cryptobytes); err != nil {
		return nil, err
	}
	return cryptobytes, nil
}

// Random encryption key may not be what you want in order to be able decrypt text. Use with caution
func RandomKey(randomBytes... []byte) (string, error) {
	Cryptobytes  :=	[]byte{}
	//cryptobytes := []byte{}

	if (len(randomBytes)==0){
		val, err := RandomBytes(32)
		if (err!=nil) {
			return "", err
		} else {
			Cryptobytes = val
		}		
	} else {
		Cryptobytes = randomBytes[0]
	}	
	key := hex.EncodeToString( Cryptobytes ) 
	return key, nil
}