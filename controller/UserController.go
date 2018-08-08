package controller

import (
	"github.com/labstack/echo"
	"net/http"
)

func UserController (c echo.Context) error {
	return c.JSON(http.StatusOK, "")
}
