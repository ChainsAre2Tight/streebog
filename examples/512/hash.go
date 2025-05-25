package main

import (
	"fmt"

	"github.com/ChainsAre2Tight/streebog"
)

func main() {
	message := []byte("any-message")
	hash := streebog.New(64)
	hash.Write(message)
	fmt.Printf("Hash: %x\n", hash.Sum(nil))
}
