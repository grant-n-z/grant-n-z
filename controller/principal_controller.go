package controller

import (
	"github.com/labstack/echo"
	"net/http"
	"github.com/tomoyane/grant-n-z/domain/entity"
	"github.com/tomoyane/grant-n-z/handler"
	"github.com/tomoyane/grant-n-z/di"
	"github.com/tomoyane/grant-n-z/infra"
)

func PostPrincipal(c echo.Context) (err error) {
	token := c.Request().Header.Get("Authorization")
	errAuth := di.ProviderTokenService.VerifyToken(c, token)

	if errAuth != nil {
		return echo.NewHTTPError(errAuth.Code, errAuth)
	}

	principal := new(entity.Principal)
	if err = c.Bind(principal); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, handler.BadRequest(""))
	}

	if err = c.Validate(principal); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, handler.BadRequest(""))
	}

	principalData, errPrincipal := di.ProviderPrincipalService.PostPrincipalData(principal)
	if errPrincipal != nil {
		return echo.NewHTTPError(errPrincipal.Code, errPrincipal)
	}

	c.Response().Header().Add("Location", infra.GetHostName() + "/v1/principals/" + principalData.Uuid.String())
	return c.JSON(http.StatusCreated, principalData)
}