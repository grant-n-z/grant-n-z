package controller

import (
	"github.com/labstack/echo"
	"net/http"
	"github.com/tomoyane/grant-n-z/di"
)

func GrantToken(c echo.Context) (err error) {
	token := c.Request().Header.Get("Authorization")

	result := di.ProviderTokenService.VerifyToken(c, di.ProviderUserService, di.ProviderRoleService, token)

	if result != nil {
		return result
	}

	success := map[string]bool {
		"authority": true,
	}

	return c.JSON(http.StatusOK, success)
}
