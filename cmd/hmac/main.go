package main

import (
	"bufio"
	"encoding/hex"
	"fmt"
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
	hash, err := streebog.HMAC256(key, text)
	if err != nil {
		log.Fatalf("Error during HMAC256 computation: %s", err)
	}
	fmt.Printf("HMAC 256: %s\n", hex.EncodeToString(hash))
	hash, err = streebog.HMAC512(key, text)
	if err != nil {
		log.Fatalf("Error during HAMC512 hash computation: %s", err)
	}
	fmt.Printf("HMAC 512: %s\n", hex.EncodeToString(hash))
}
