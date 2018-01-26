package main

import (
	"bytes"
	"encoding/binary"
	"encoding/gob"
	"log"

	bolt "github.com/coreos/bbolt"
)

type Db struct {
	bucketName []byte
	db         *bolt.DB
}

func serialize(b Block) (id []byte, block []byte) {
	var result bytes.Buffer

	if err := gob.NewEncoder(&result).Encode(b); err != nil {
		log.Fatal(err)
	}

	id = make([]byte, 4)
	binary.LittleEndian.PutUint32(id, b.Id)

	return id, result.Bytes()
}

func deserialize(d []byte) (b *Block) {
	decoder := gob.NewDecoder(bytes.NewReader(d))

	if err := decoder.Decode(&b); err != nil {
		log.Fatal(err)
	}

	return
}

func (d *Db) open(fileName string) {
	db, err := bolt.Open(fileName, 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	d.db = db
}

func (d *Db) put(key, value []byte) {
	d.db.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists(d.bucketName)
		if err != nil {
			log.Fatal(err)
		}
		return bucket.Put(key, value)
	})
}

func (d *Db) get(key []byte) (result []byte) {
	d.db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(d.bucketName)
		result = bucket.Get(key)
		return nil
	})
	return
}
