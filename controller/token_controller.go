package controller

import (
	"github.com/labstack/echo"
	"net/http"
)

func PostToken(c echo.Context) (err error) {
	success := map[string]string {
		"message": "user creation succeeded.",
	}

	return c.JSON(http.StatusCreated, success)
}