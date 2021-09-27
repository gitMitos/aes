package aes

import "testing"

func Test_TwoKeysNotSame(t *testing.T) {
	key1, _ := RandomKey()
	key2, _ := RandomKey()
	if (key1==key2) {
		t.Error("2 random keys are same")
	}
}

func Test_KeyLength(t *testing.T) {
	key, _ := RandomKey()
	if (len(key)!=64) {
		t.Errorf("Key length is %v not 64", len(key))
	}
}

// func Test_KeyLengthMustFail(t *testing.T) {
// 	key, _ := RandomKey() 
// 	key += "Extra"

// 	if (len(key)!=64) {
// 		t.Errorf("Key length is %v not 64", len(key))
// 	}	

// }

func Benchmark_Randomkey(b *testing.B){
	for i:=0; i<b.N; i++ {
		_, _ = RandomKey()
	}
}