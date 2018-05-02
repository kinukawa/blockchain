package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type resource struct {
	blockchain *Blockchain
}

func main() {

	resc := resource{
		blockchain: NewBlockchain(),
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
	return c.JSON(http.StatusOK, "mine new block")
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
