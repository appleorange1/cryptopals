package main

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Not enough arguments")
	}
	input := os.Args[1]
	src := []byte(input)
	dst := make([]byte, hex.DecodedLen(len(src)))
	n, err := hex.Decode(dst, src)
	if err != nil {
		log.Fatal(err)
	}
	str := base64.StdEncoding.EncodeToString(dst[:n])
	fmt.Printf("%s\n", str)
}
