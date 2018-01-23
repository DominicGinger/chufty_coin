package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	data := []byte("Some Data")
	hash := sha256.Sum256(data)
	fmt.Printf("Data: %s\n", data)
	fmt.Printf("Hash: %x\n", hash)
}
