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
	e.POST("/transactions/new", resc.newTransaction)

	// Start server
	e.Logger.Fatal(e.Start(":1234"))
}

// Handler
func (resc resource) newTransaction(c echo.Context) error {
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
