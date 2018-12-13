package main

import (
	"encoding/hex"
	"fmt"
	"log"
	"os"
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

func main() {
	if len(os.Args) >= 3 {
		l1, s1 := hexDecode(os.Args[1])
		l2, s2 := hexDecode(os.Args[2])
		if l1 != l2 {
			log.Fatal("String length is not the same")
		}
		cipher := make([]byte, l1)
		for i := 0; i <= l1 - 1; i++ {
			cipher[i] = s1[i] ^ s2[i]
		}
		hexcipher := hex.EncodeToString(cipher)
		fmt.Printf("%s\n", hexcipher)
	} else {
		log.Fatal("Not enough arguments")
	}
}
