package main

import (
	"crypto/sha256"
	"fmt"
	"time"
)

type Blockchain struct {
	// empty list for store transactions
	currentTrunsactions []Transaction

	// first empty list for store blockchains
	chain []*Block
}

func NewBlockchain() *Blockchain {
	bc := &Blockchain{}
	bc.createBlock(100, []byte("1"))
	return bc
}

// createBlock create new block and add to chain
func (bc *Blockchain) createBlock(proof int, previousHash []byte) Block {

	prevHash := previousHash
	if prevHash == nil {
		prevHash = bc.lastBlock().hash()
	}

	block := Block{
		index:        len(bc.chain) + 1,
		timestamp:    time.Now(),
		transactions: bc.currentTrunsactions,
		proof:        proof,
		previousHash: prevHash,
	}

	// reset currentTransactions
	bc.currentTrunsactions = []Transaction{}

	bc.chain = append(bc.chain, &block)
	return block
}

// createTransaction create new transaction and add to list
func (bc *Blockchain) createTransaction(sender string, recipient string, amount int) int {

	// create transaction to add to next mined block
	transaction := Transaction{
		sender:    sender,
		recipient: recipient,
		amount:    amount,
	}

	bc.currentTrunsactions = append(bc.currentTrunsactions, transaction)

	// return address of the block contain this transaction
	return bc.lastBlock().index + 1
}

// lastBlock is return a first block
func (bc *Blockchain) lastBlock() *Block {
	return bc.chain[len(bc.chain)-1]
}

func (bc *Blockchain) proofOfWork(lastProof int) int {
	proof := 0
	for !bc.validProof(lastProof, proof) {
		proof += 1
	}
	return proof
}

// validProof valid proof
func (bc *Blockchain) validProof(lastProof int, proof int) bool {
	guess := fmt.Sprintf("%d%d", lastProof, proof)
	guessHash := sha256.Sum256([]byte(guess))
	return string(guessHash[:2]) == "12"
}
