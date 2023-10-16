package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {

	// instantiate echo
	e := echo.New()

	// Route
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello world!")
	})
	e.Logger.Fatal(e.Start(":1323"))
}
