package main

import (
	"crypto/hmac"
	"fmt"
	"hash"

	"github.com/ChainsAre2Tight/streebog"
)

func main() {
	key := []byte("any-key")
	message := []byte("any-message")
	h := hmac.New(func() hash.Hash { return streebog.New(32) }, key)
	h.Write(message)
	fmt.Printf("Hash: %x\n", h.Sum(nil))
}
