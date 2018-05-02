package main

import (
	"net/http"

	"encoding/hex"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/satori/go.uuid"
)

type resource struct {
	blockchain     *Blockchain
	nodeIdentifier string
}

func main() {
	uuid := uuid.NewV4()
	resc := resource{
		blockchain:     NewBlockchain(),
		nodeIdentifier: hex.EncodeToString(uuid[:]),
	}

	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.POST("/transactions/new", resc.postNewTransaction)
	e.GET("/mine", resc.getMine)
	e.GET("/chain", resc.getChain)

	// Start server
	e.Logger.Fatal(e.Start(":1234"))
}

// postNewTransaction post new transaction
func (resc resource) postNewTransaction(c echo.Context) error {
	body := new(Transaction)
	if err := c.Bind(body); err != nil {
		return err
	}
	if body.Amount == 0 || body.Recipient == "" || body.Sender == "" {
		return c.JSON(http.StatusBadRequest, body)
	}

	// create a new transaction
	index := resc.blockchain.createTransaction(body.Sender, body.Recipient, body.Amount)

	resp := struct {
		Index int `json:"index"`
	}{
		Index: index,
	}

	return c.JSON(http.StatusCreated, resp)
}

// getMine mine new block
func (resc resource) getMine(c echo.Context) error {

	// find next proof
	lastBlock := resc.blockchain.lastBlock()
	lastProof := lastBlock.Proof
	proof := resc.blockchain.proofOfWork(lastProof)

	//
	resc.blockchain.createTransaction(
		"0",
		resc.nodeIdentifier,
		1,
	)

	// mine new block by adding block to chain
	block := resc.blockchain.createBlock(proof, nil)

	resp := struct {
		Message      string
		Index        int
		Transactions []*Transaction
		Proof        int
		PreviousHash []byte
	}{
		Message:      "mine a new block",
		Index:        block.Index,
		Transactions: block.Transactions,
		Proof:        block.Proof,
		PreviousHash: block.PreviousHash,
	}

	return c.JSON(http.StatusOK, resp)
}

// getChain return full blockchain
func (resc resource) getChain(c echo.Context) error {

	resp := struct {
		Chain  []*Block
		Length int
	}{
		Chain:  resc.blockchain.Chain,
		Length: len(resc.blockchain.Chain),
	}

	return c.JSON(http.StatusCreated, resp)
}
