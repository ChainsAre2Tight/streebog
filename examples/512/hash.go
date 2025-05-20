package main

import (
	"fmt"
	"log"

	"github.com/ChainsAre2Tight/streebog"
)

func main() {
	message := []byte("any-message")
	hash, err := streebog.Streebog512(message)
	if err != nil {
		log.Fatalf("Error during hash computation: %s", err)
	}
	fmt.Printf("Hash: %x\n", hash)
}
