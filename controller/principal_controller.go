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
	principal := new(entity.Principal)

	if err = c.Bind(principal); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, handler.BadRequest(""))
	}

	if err = c.Validate(principal); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, handler.BadRequest(""))
	}

	principalData := di.ProviderPrincipalService.GetPrincipalByName(principal.Name)
	if principalData == nil {
		return echo.NewHTTPError(http.StatusInternalServerError, handler.InternalServerError(""))
	}

	if len(principalData.Name) > 0 {
		return echo.NewHTTPError(http.StatusConflict, handler.Conflict(""))
	}

	principalData = di.ProviderPrincipalService.InsertRole(*principal)
	if principalData == nil {
		return echo.NewHTTPError(http.StatusInternalServerError, handler.InternalServerError(""))
	}

	c.Response().Header().Add("Location", infra.GetHostName() + "/v1/roles/" + principalData.Uuid.String())
	return c.JSON(http.StatusCreated, principalData)
}