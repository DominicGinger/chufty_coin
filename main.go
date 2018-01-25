package main

import (
	"encoding/hex"
	"fmt"
)

func validHash(hash [32]byte, miningRule []byte) bool {
	for i, m := range miningRule {
		if m != hash[i] {
			return false
		}
	}
	return true
}

func compareHash(a, b [32]byte) bool {
	for i, x := range a {
		if b[i] != x {
			return false
		}
	}
	return true
}

func main() {
	data := []byte("Some Data")
	minigRule, _ := hex.DecodeString("beef")
	b := Block{id: 1, data: data}
	b.mineBlock(minigRule, 1)
	fmt.Printf("id: %v\n", b.id)
	fmt.Printf("nonce: %v\n", b.nonce)
	fmt.Printf("data: %s\n", b.data)
	fmt.Printf("hash: %x\n", b.hash)
	fmt.Printf("valid: %v\n", b.validate())
}
