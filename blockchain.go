package main

type Blockchain struct {
	// empty list for store transactions
	currentTrunsactions []Transaction

	// first empty list for store blockchains
	chain []Block
}

// createBlock create new block and add to chain
func (bc *Blockchain) createBlock(proof int, previousHash []byte) {
}

// createTransaction create new transaction and add to list
func (bc *Blockchain) createTransaction(sender string, recipient string, amount int) {
}

// lastBlock is return a first block
func (bc *Blockchain) lastBlock() {

}
