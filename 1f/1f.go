package main

import (
	"fmt"
)

func main() {
	s1 := []byte("this is a test")
	s2 := []byte("wokka wokka!!!")
	xorsl := make([]byte, len(s1))
	for i := 0; i < len(s1); i++ {
		xorsl[i] = s1[i] ^ s2[i]
	}
	fmt.Sprintf("%b", xorsl)
	/*
	for keysize := 2; keysize <= 40; keysize++ {
	}
	*/
}
