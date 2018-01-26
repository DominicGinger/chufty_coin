package main

import (
	"encoding/hex"
)

func main() {
	data := []byte("Some Data")
	minigRule, _ := hex.DecodeString("beef")
	b := Block{Id: 1, Data: data}
	b.mineBlock(minigRule, 1)
	b.print()

	db := Db{bucketName: []byte("blockchain")}
	db.open("./blockchain.db")
	defer db.db.Close()

	id, body := serialize(b)
	db.put(id, body)

	result := db.get(id)
	b2 := deserialize(result)
	b2.print()
}
