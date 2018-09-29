package controller

import (
	"github.com/labstack/echo"
	"github.com/tomoyane/grant-n-z/di"
	"github.com/tomoyane/grant-n-z/domain/entity"
	"github.com/tomoyane/grant-n-z/handler"
	"github.com/tomoyane/grant-n-z/infra"
	"net/http"
)

func PostPrincipal(c echo.Context) (err error) {
	token := c.Request().Header.Get("Authorization")
	errAuth := di.ProvideTokenService.VerifyToken(c, token)

	if errAuth != nil {
		return echo.NewHTTPError(errAuth.Code, errAuth)
	}

	principalRequest := new(entity.PrincipalRequest)
	if err = c.Bind(principalRequest); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, handler.BadRequest(""))
	}

	if err = c.Validate(principalRequest); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, handler.BadRequest(""))
	}

	principalData, errPrincipal := di.ProvidePrincipalService.PostPrincipalData(*principalRequest)
	if errPrincipal != nil {
		return echo.NewHTTPError(errPrincipal.Code, errPrincipal)
	}

	c.Response().Header().Add("Location", infra.GetHostName() + "/v1/principals/" + principalData.Uuid.String())
	return c.JSON(http.StatusCreated, principalData)
}
