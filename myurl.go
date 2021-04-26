package main

import (
	"encoding/json"
	"fmt"
	"github.com/boltdb/bolt"
	"log"
	"math/rand"
)

type MyUrl struct {
	UniqueId string
	LongUrl  string
	ShortUrl string
}

func CreateId() string {
	bytes := make([]byte, 5)
	chars := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	for i := 0; i < 5; i++ {
		bytes[i] = chars[rand.Intn(len(chars))]
	}
	return string(bytes)
}

func AddToStorage(myUrl MyUrl) {
	db, err := bolt.Open("my.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte("URLs"))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		encoded, err := json.Marshal(myUrl)
		fmt.Printf("%v", myUrl)
		return b.Put([]byte(myUrl.UniqueId), encoded)
	})
}
