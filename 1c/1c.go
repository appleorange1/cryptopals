package main

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"log"
	"sort"
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

func score(candidate []byte) int {
	frequencies := map [rune] int {
		'a': 8167,
		'b': 1492,
		'c': 2782,
		'd': 4253,
		'e': 12702,
		'f': 2228,
		'g': 2015,
		'h': 6094,
		'i': 6966,
		'j': 153,
		'k': 3872,
		'l': 4025,
		'm': 2406,
		'n': 6749,
		'o': 7507,
		'p': 1929,
		'q': 95,
		'r': 5987,
		's': 6327,
		't': 9256,
		'u': 2758,
		'v': 978,
		'w': 5370,
		'x': 150,
		'y': 3978,
		'z': 74,
	}
	lcandidate := bytes.ToLower(candidate)
	sc := 0
	for i := 0; i <= len(lcandidate) - 2; i++ {
		freq, ok := frequencies[rune(lcandidate[i])]
		if ok {
			sc += freq
		} else {
			//sc = -1
			//break
		}
	}
	return sc
}

func main() {
	hexcipher := []byte("1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736")
	cipher := make([]byte, hex.DecodedLen(len(hexcipher)))
	n, err := hex.Decode(cipher, hexcipher)
	if err != nil {
		log.Fatal(err)
	}
	scores := make([]int, 128)
	scoremap := make(map [int] []byte)
	for i := 0; i <= 127; i++ {
		testslice := bytes.Repeat([]byte(string(i)), n)
		plaintext := fixedXOR(testslice, cipher)
		sc := score(plaintext)
		scores[i] = sc
		scoremap[sc] = plaintext
	}
	sort.Ints(scores)
	for i := 117; i <= 127; i++ {
		fmt.Printf("%d: %d: %s\n", i, scores[i], scoremap[scores[i]])
	}
}
