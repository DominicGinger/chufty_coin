package main

import (
	"bytes"
	"encoding/gob"
	"log"

	bolt "github.com/coreos/bbolt"
)

type Db struct {
	bucketName []byte
	db         *bolt.DB
}

func serialize(b Block) []byte {
	var result bytes.Buffer

	if err := gob.NewEncoder(&result).Encode(b); err != nil {
		log.Fatal(err)
	}

	return result.Bytes()
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

func (d *Db) put(pairs map[string][]byte) {
	d.db.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists(d.bucketName)
		if err != nil {
			log.Fatal(err)
		}
		for key, value := range pairs {
			if err = bucket.Put([]byte(key), value); err != nil {
				log.Fatal(err)
			}
		}
		return nil
	})
}

func (d *Db) deepGet(key []byte, depth int) []byte {
	result := key
	d.db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(d.bucketName)
		for i := 0; i <= depth; i++ {
			result = bucket.Get(result)
		}
		return nil
	})
	return result
}

func (d *Db) get(keys []string) [][]byte {
	result := make([][]byte, len(keys))
	d.db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(d.bucketName)
		for _, key := range keys {
			result = append(result, bucket.Get([]byte(key)))
		}
		return nil
	})
	return result
}
