package main

import (
	"crypto/sha256"
	"encoding/json"
	"time"
)

type Block struct {
	index        int
	timestamp    time.Time
	transactions []Transaction
	proof        int
	previousHash []byte
}

func (b *Block) hash() []byte {
	json, err := json.Marshal(b)
	if err != nil {
		panic(err)
	}

	bytes := sha256.Sum256(json)
	return bytes[:]
}
