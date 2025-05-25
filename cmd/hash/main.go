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
	h256 := streebog.New(32)
	h512 := streebog.New(64)
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter text: ")
	text, err := reader.ReadBytes('\n')
	if err != nil {
		log.Fatalf("Error: %s", err)
	}
	text = text[:len(text)-1]

	h256.Write(text)
	h512.Write(text)
	fmt.Printf("Hash 256: %s\n", hex.EncodeToString(h256.Sum(nil)))
	fmt.Printf("Hash 512: %s\n", hex.EncodeToString(h512.Sum(nil)))
}
