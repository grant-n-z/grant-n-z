package controller

import (
	"github.com/labstack/echo"
	"net/http"
	"github.com/tomoyane/grant-n-z/di"
)

func GrantToken(c echo.Context) (err error) {
	token := c.Request().Header.Get("Authorization")
	errAuth := di.ProviderTokenService.VerifyToken(c, token)
	if errAuth != nil {
		return echo.NewHTTPError(errAuth.Code, errAuth)
	}

	return c.JSON(http.StatusOK, map[string]bool {"authority": true})
}
