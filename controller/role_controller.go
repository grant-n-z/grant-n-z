package controller

import (
	"github.com/labstack/echo"
	"net/http"
	"github.com/tomoyane/grant-n-z/domain/entity"
	"github.com/tomoyane/grant-n-z/di"
	"github.com/tomoyane/grant-n-z/infra"
)

func PostRole(c echo.Context) (err error) {
	token := c.Request().Header.Get("Authorization")

	errAuth := di.ProviderTokenService.VerifyToken(c, token)

	if errAuth != nil {
		return echo.NewHTTPError(errAuth.Code, errAuth)
	}

	role := new(entity.Role)
	roleData, errRole := di.ProviderRoleService.PostRoleData(c, role, token)

	if errRole != nil {
		return echo.NewHTTPError(errRole.Code, errRole)
	}

	c.Response().Header().Add("Location", infra.GetHostName() + "/v1/roles/" + role.Uuid.String())
	return c.JSON(http.StatusCreated, roleData)
}