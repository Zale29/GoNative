package controllers

import (
	"Go/NativeQuery/helper"
	"Go/NativeQuery/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Checklogin(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	res, err := models.Checklogin(username, password)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"messages": err.Error(),
		})
	}

	if !res {
		return echo.ErrUnauthorized
	}
	return c.String(http.StatusOK, "Behasil Login")

	// Create token with claims
	// token := jwt.New(jwt.SigningMethodHS256)

	// claims := token.Claims.(jwt.MapClaims)
	// claims["username"] = username
	// claims["level"] = "application"
	// claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	// Generate encoded token and send it as response.
	// t, err := token.SignedString([]byte("secret"))
	// if err != nil {
	// 	return err
	// }

	// return c.JSON(http.StatusOK, echo.Map{
	// 	"token": t,
	// })
}

func GenerateFromPassword(c echo.Context) error {
	password := c.Param("password")

	hash, _ := helper.Hashpassword(password)

	return c.JSON(http.StatusOK, hash)
}
