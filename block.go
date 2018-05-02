package main

import (
	"crypto/sha256"
	"encoding/json"
	"time"
)

type Block struct {
	Index        int
	Timestamp    time.Time
	Transactions []*Transaction
	Proof        int
	PreviousHash []byte
}

func (b *Block) hash() []byte {
	json, err := json.Marshal(b)
	if err != nil {
		panic(err)
	}

	bytes := sha256.Sum256(json)
	return bytes[:]
}
