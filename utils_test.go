package aes

import (
	"bytes"
	"encoding/hex"
	"testing"
)

func Test_IsBytesSame(t *testing.T){
	b, _ := RandomBytes(32)
	//text := string(b)
	h, _ := hex.DecodeString( hex.EncodeToString(b) )
	if ( !bytes.Equal(h , b) ){
		t.Error("Some problem with byte to hex conversion")
	}
	// 
	//t.Log(b,text, []byte(text), hex.EncodeToString(b), h )

}

func Test_TwoKeysNotSame(t *testing.T) {
	key1, _ := RandomKey()
	key2, _ := RandomKey()
	if (key1==key2) {
		t.Error("2 random keys are same")
	}
}

func Test_KeyLengthWithWrongSizeMustFail(t *testing.T){
	b := []byte{'0', '0', '0', '0', '0'}
	key, err := RandomKey(b)
	t.Log(key, err)
}

func Test_KeyLengthWithSize(t *testing.T) {
	//size := 64
	//
	key, _ := RandomKey()
	t.Log(key)	 
	if (len(key)!=64) {
		t.Errorf("Key length is %v not 64", len(key))
	}
}

func Test_RandomBytes64(t *testing.T){
	size := 64
	bytes, _ := RandomBytes( size )
	t.Log(bytes)
}
// func Test_KeyLengthMustFail(t *testing.T) {
// 	key, _ := RandomKey() 
// 	key += "Extra"

// 	if (len(key)!=64) {
// 		t.Errorf("Key length is %v not 64", len(key))
// 	}	

// }

func Benchmark_RandomkeyBytes(b *testing.B){
	for i:=0; i<b.N; i++ {
		_, _ = RandomBytes(32)
	}
}

func Benchmark_RandomkeyString(b *testing.B){
	for i:=0; i<b.N; i++ {
		_, _ = RandomKey()
	}
}