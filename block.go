package main

import (
	"crypto/sha256"
	"encoding/binary"
)

// Block struct
type Block struct {
	id       uint32
	nonce    uint32
	data     []byte
	prevHash [32]byte
	hash     [32]byte
}

func (b *Block) hashBlock() [32]byte {
	id := make([]byte, 4)
	nonce := make([]byte, 4)
	binary.LittleEndian.PutUint32(id, b.id)
	binary.LittleEndian.PutUint32(nonce, b.nonce)

	bs := append(id, append(nonce, b.data...)...)

	return sha256.Sum256(bs)
}

func (b *Block) mineBlock(rule []byte, nonce uint32) {
	b.nonce = nonce
	hash := b.hashBlock()
	for !validHash(hash, rule) {
		b.nonce = b.nonce + 1
		hash = b.hashBlock()
	}
	b.hash = hash
}

func (b *Block) validate() bool {
	return compareHash(b.hash, b.hashBlock())
}
