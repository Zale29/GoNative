package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello This is Echo!")
	})

	e.Logger.Fatal(e.Start(":1234"))

}
