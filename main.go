package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.POST("/transactions/new", newTransaction)

	// Start server
	e.Logger.Fatal(e.Start(":1234"))
}

// Handler
func newTransaction(c echo.Context) error {
	body := new(Transaction)
	if err := c.Bind(body); err != nil {
		return err
	}
	if body.Amount == 0 || body.Recipient == "" || body.Sender == "" {
		return c.JSON(http.StatusBadRequest, body)
	}

	return c.JSON(http.StatusCreated, body)
}
