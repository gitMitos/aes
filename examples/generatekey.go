package main

import (
	"fmt"

	aes "github.com/gitMitos/aes"
)

func main() {
	key, err :=  aes.RandomKey()
	if (err!=nil){
		panic("Fail to create random encryption key")
	}

	fmt.Println( "Key is: ", key )
	fmt.Println(" len of key is: ", len(key))
}