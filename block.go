package main

import "time"

type Block struct {
	index        int
	timestamp    time.Time
	transactions []Transaction
	proof        int
	previousHash []byte
}
