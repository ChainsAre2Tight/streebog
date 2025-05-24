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
	fmt.Print("Enter text: ")
	text, err := reader.ReadBytes('\n')
	if err != nil {
		log.Fatalf("Error during stdin decoding: %s", err)
	}
	text = text[:len(text)-1]
	hash, err := streebog.Streebog256(text)
	if err != nil {
		log.Fatalf("Error during 256 hash computation: %s", err)
	}
	fmt.Printf("Hash 256: %s\n", hex.EncodeToString(hash))
	hash, err = streebog.Streebog512(text)
	if err != nil {
		log.Fatalf("Error during 512 hash computation: %s", err)
	}
	fmt.Printf("Hash 512: %s\n", hex.EncodeToString(hash))
}
