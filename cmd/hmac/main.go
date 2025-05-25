package main

import (
	"bufio"
	"crypto/hmac"
	"encoding/hex"
	"fmt"
	"hash"
	"log"
	"os"

	"github.com/ChainsAre2Tight/streebog"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter key: ")

	key, err := reader.ReadBytes('\n')
	if err != nil {
		log.Fatalf("Error during stdin decoding: %s", err)
	}
	key = key[:len(key)-1]

	fmt.Print("Enter message: ")
	text, err := reader.ReadBytes('\n')
	if err != nil {
		log.Fatalf("Error during stdin decoding: %s", err)
	}
	text = text[:len(text)-1]

	h := hmac.New(func() hash.Hash { return streebog.New(32) }, key)
	h.Write(text)
	fmt.Printf("HMAC 256: %s\n", hex.EncodeToString(h.Sum(nil)))

	h = hmac.New(func() hash.Hash { return streebog.New(64) }, key)
	h.Write(text)
	fmt.Printf("HMAC 512: %s\n", hex.EncodeToString(h.Sum(nil)))
}
