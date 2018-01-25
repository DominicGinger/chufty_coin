package main

import (
	"bytes"
	"encoding/gob"
	"encoding/hex"
	"log"

	"github.com/boltdb/bolt"
)

func deserializeBlock(d []byte) (b *Block) {
	decoder := gob.NewDecoder(bytes.NewReader(d))

	if err := decoder.Decode(&b); err != nil {
		log.Fatal(err)
	}

	return
}

func main() {
	data := []byte("Some Data")
	minigRule, _ := hex.DecodeString("beef")
	b := Block{Id: 1, Data: data}
	b.mineBlock(minigRule, 1)
	b.print()

	db, err := bolt.Open("./blockchain.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	bucketName := []byte("blockchain")

	db.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists(bucketName)
		if err != nil {
			log.Fatal(err)
		}
		id, body := b.serialize()
		return bucket.Put(id, body)
	})

	db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(bucketName)
		id, _ := b.serialize()
		serialized := bucket.Get(id)
		b2 := deserializeBlock(serialized)
		b2.print()
		return nil
	})
}
