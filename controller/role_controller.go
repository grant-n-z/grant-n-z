package controller

import (
	"github.com/labstack/echo"
	"github.com/tomoyane/grant-n-z/di"
	"github.com/tomoyane/grant-n-z/domain/entity"
	"github.com/tomoyane/grant-n-z/handler"
	"github.com/tomoyane/grant-n-z/infra"
	"net/http"
)

func PostRole(c echo.Context) (err error) {
	token := c.Request().Header.Get("Authorization")
	errAuth := di.ProvideTokenService.VerifyToken(c, token)

	if errAuth != nil {
		return echo.NewHTTPError(errAuth.Code, errAuth)
	}

	role := new(entity.Role)
	if err := c.Bind(role); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, handler.BadRequest(""))
	}

	if err := c.Validate(role); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, handler.BadRequest(""))
	}

	roleData, errRole := di.ProvideRoleService.PostRoleData(role)
	if errRole != nil {
		return echo.NewHTTPError(errRole.Code, errRole)
	}

	c.Response().Header().Add("Location", infra.GetHostName() + "/v1/roles/" + roleData.Uuid.String())
	return c.JSON(http.StatusOK, roleData)
}