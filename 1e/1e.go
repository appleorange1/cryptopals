package main

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"log"
)

func hexDecode(input string) (int, []byte) {
	src := []byte(input)
	dst := make([]byte, hex.DecodedLen(len(src)))
	n, err := hex.Decode(dst, src)
	if err != nil {
		log.Fatal(err)
	}
	return n, dst
}

func fixedXOR(arg1 []byte, arg2 []byte) []byte {
	if len(arg1) != len(arg2) {
		log.Fatal("Length is not the same")
	}
	cipher := make([]byte, len(arg1))
	for i := 0; i <= len(arg1) - 1; i++ {
		cipher[i] = arg1[i] ^ arg2[i]
	}
	return cipher
}

func main() {
	plaintext := []byte("Burning 'em, if you ain't quick and nimble\nI go crazy when I hear a cymbal")
	iceslice := bytes.Repeat([]byte(string("ICE")), len(plaintext) / 3)
	if len(plaintext) != len(iceslice) {
		iceslice = append(iceslice, []byte("ICE")[:len(plaintext) - len(iceslice)]...)
	}
	cipher := fixedXOR(iceslice, plaintext)
	hexcipher := hex.EncodeToString(cipher)
	fmt.Printf("%s\n", hexcipher)
}
