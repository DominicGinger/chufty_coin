package main

import (
	"encoding/hex"
)

func main() {
	minigRule, _ := hex.DecodeString("beef")
	db := Db{bucketName: []byte("blockchain")}
	db.open("./blockchain.db")
	defer db.db.Close()

	data := []byte("Some Data")
	b := Block{Id: 1, Data: data}
	b.mineBlock(minigRule)

	body := serialize(b)
	db.put(map[string][]byte{string(b.Hash[:]): body, "tip": b.Hash[:]})

	secondBlock := Block{Id: 2, Data: []byte("data 2"), PrevHash: b.Hash}
	secondBlock.mineBlock(minigRule)
	body = serialize(secondBlock)
	db.put(map[string][]byte{string(secondBlock.Hash[:]): body, "tip": secondBlock.Hash[:]})

	result := db.deepGet([]byte("tip"), 1)
	block := deserialize(result)
	block.print()
	result = db.deepGet(block.PrevHash[:], 0)
	block = deserialize(result)
	block.print()
}
