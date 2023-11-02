package routes

import (
	"net/http"
	"Go/NativeQuery/controlles"
	"github.com/labstack/echo/v4"
)

func Init() *echo.Echo {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello This is Echo!")
	})

	e.GET("/pegawai", controllers.FetchallPegawai)

	return e

}
