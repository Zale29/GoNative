package controllers

import (
	"Go/NativeQuery/models"
)

func FetchallPegawai(c echo.Context) error {
	result, err := models.FetchallPegawai()

	if err != {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}