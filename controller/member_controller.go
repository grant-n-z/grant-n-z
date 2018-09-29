package controller

import (
	"github.com/labstack/echo"
	"github.com/tomoyane/grant-n-z/di"
	"net/http"
)

func PostMember(c echo.Context) (err error) {
	token := c.Request().Header.Get("Authorization")
	errAuth := di.ProviderTokenService.VerifyToken(c, token)
	if errAuth != nil {
		return echo.NewHTTPError(errAuth.Code, errAuth)
	}

	return c.JSON(http.StatusOK, map[string]bool {"authority": true})
}
