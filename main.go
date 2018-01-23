package main

import (
	"crypto/sha256"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"log"
)

type Block struct {
	id    uint32
	nonce uint32
	data  []byte
	hash  [32]byte
}

func createBlock(i uint32, data []byte) Block {
	b := Block{id: i, data: data}
	rule, err := hex.DecodeString("beef")
	if err != nil {
		log.Fatal(err)
	}
	return mineBlock(b, rule, 1)
}

func hashBlock(b Block) [32]byte {
	id := make([]byte, 4)
	nonce := make([]byte, 4)
	binary.LittleEndian.PutUint32(id, b.id)
	binary.LittleEndian.PutUint32(nonce, b.nonce)

	bs := append(id, append(nonce, b.data...)...)

	return sha256.Sum256(bs)
}

func validHash(hash [32]byte, miningRule []byte) bool {
	for i, m := range miningRule {
		if m != hash[i] {
			return false
		}
	}
	return true
}

func mineBlock(b Block, rule []byte, nonce uint32) Block {
	b.nonce = nonce
	hash := hashBlock(b)
	for !validHash(hash, rule) {
		b.nonce = b.nonce + 1
		hash = hashBlock(b)
	}
	b.hash = hash
	return b
}

func compareHash(a, b [32]byte) bool {
	for i, x := range a {
		if b[i] != x {
			return false
		}
	}
	return true
}

func validateBlock(b Block) bool {
	return compareHash(b.hash, hashBlock(b))
}

func main() {
	data := []byte("Some Data")
	b := createBlock(1, data)
	fmt.Printf("id: %v\n", b.id)
	fmt.Printf("nonce: %v\n", b.nonce)
	fmt.Printf("data: %s\n", b.data)
	fmt.Printf("hash: %x\n", b.hash)
	fmt.Printf("valid: %v\n", validateBlock(b))
}
