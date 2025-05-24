package main

import (
	"fmt"
	"log"

	"github.com/ChainsAre2Tight/streebog"
)

func main() {
	key := []byte("any-key")
	message := []byte("any-message")
	hash, err := streebog.HMAC256(key, message)
	if err != nil {
		log.Fatalf("Error during hmac computation: %s", err)
	}
	fmt.Printf("Hash: %x\n", hash)
}
