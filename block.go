package main

import (
	"crypto/sha256"
	"encoding/binary"
	"fmt"
)

// Block struct
type Block struct {
	Id       uint32
	Nonce    uint32
	Data     []byte
	PrevHash [32]byte
	Hash     [32]byte
}

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

func (b *Block) hashBlock() [32]byte {
	id := make([]byte, 4)
	nonce := make([]byte, 4)
	binary.LittleEndian.PutUint32(id, b.Id)
	binary.LittleEndian.PutUint32(nonce, b.Nonce)

	bs := append(id, append(nonce, b.Data...)...)

	return sha256.Sum256(bs)
}

func (b *Block) mineBlock(rule []byte, nonce uint32) {
	b.Nonce = nonce
	hash := b.hashBlock()
	for !validHash(hash, rule) {
		b.Nonce = b.Nonce + 1
		hash = b.hashBlock()
	}
	b.Hash = hash
}

func (b *Block) validate() bool {
	return compareHash(b.Hash, b.hashBlock())
}

func (b *Block) print() {
	fmt.Printf("id: %v\n", b.Id)
	fmt.Printf("nonce: %v\n", b.Nonce)
	fmt.Printf("data: %s\n", b.Data)
	fmt.Printf("hash: %x\n", b.Hash)
	fmt.Printf("valid: %v\n", b.validate())
}
