package main

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"log"
	"sort"
)

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

func solve(block []bytes, lblock int) {
	scores := make([]int, 128)
	scoremap := make(map [int] []byte)
	for i := 0; i <= 127; i++ {
		testslice := bytes.Repeat([]byte(string(i)), lblock)
		plaintext := fixedXOR(testslice, cipher)
		sc := score(plaintext)
		scores[i] = sc
		scoremap[sc] = plaintext
	}
	sort.Ints(scores)
	for i := 125; i <= 127; i++ {
		fmt.Printf("%d: %d: %s\n", i, scores[i], scoremap[scores[i]])
	}
}


func hamming_dist(s1 []byte, s2 []byte) int {
	dist := 0
	for i := 0; i < len(s1); i++ {
		xorstr := fmt.Sprintf("%08b", s1[i] ^ s2[i])
		for j := 0; j <= 7; j++ {
			if xorstr[j] == 0x31 {
				dist += 1
			} else if xorstr[j] != 0x30 {
				log.Fatal("Error in XOR output")
			}
		}
	}
	return dist
}

func main() {
	b64, err := ioutil.ReadFile("6.txt")
	if err != nil {
		log.Fatal(err)
	}

	dst := make([]byte, hex.DecodedLen(len(b64)))
	n, err := hex.Decode(dst, b64)
	if err != nil {
		log.Fatal(err)
	}

	norm_dist_array := make([]float64, 39)

	for keysize := 2; keysize <= 40; keysize++ {
		if 2 * keysize < n {
			s1 := dst[:keysize]
			s2 := dst[keysize+1:2*keysize]
			dist := hamming_dist(s1, s2)
			norm_dist := dist / float64(keysize)
			norm_dist_array[keysize - 2] = norm_dist
		} else {
			log.Fatal("Error: 2 * keysize >= n")
		}
	}

	sort.Float64s(norm_dist_array)

	// We test the 3 keysizes with the best scores
	for i := 37; i <= 39; i++ {
		keysize := norm_dist_array[i]
		fmt.Printf("\n===Keysize: %d===\n", keysize)
		// Create expdst, which is dst padded with zeroes so that
		// the dst string can be divided up into arrays of length
		// keysize
		expdst := make([]byte, n + (n % keysize))
		for j := 0; j < n; j++ {
			expdst[j] = dst[j]
		}
		for j := 1; j < n % keysize; j++ {
			expdst[keysize + j] = 0
		}
		for j := 0; j < keysize; j++ {
			fmt.Printf("\n====Block %d====\n", j)
			lblock := len(expdst) / keysize
			tblock := make([]bytes, lblock)
			for k := 0; k < lblock; k++ {
				tblock[k] = dst[j + (k * keysize)]
			}
			solve(tblock, lblock)
		}
	}
}
