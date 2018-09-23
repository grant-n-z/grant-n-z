package controller

import (
	"github.com/labstack/echo"
	"net/http"
	"github.com/tomoyane/grant-n-z/domain/entity"
	"github.com/tomoyane/grant-n-z/handler"
	"github.com/tomoyane/grant-n-z/di"
	"github.com/tomoyane/grant-n-z/infra"
)

func PostRole(c echo.Context) (err error) {
	role := new(entity.Role)

	if err = c.Bind(role); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, handler.BadRequest(""))
	}

	if err = c.Validate(role); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, handler.BadRequest(""))
	}

	roleData := di.ProviderRoleService.GetRoleByPermission(role.Permission)
	if roleData == nil {
		return echo.NewHTTPError(http.StatusInternalServerError, handler.InternalServerError(""))
	}

	if len(roleData.Permission) > 0 {
		return echo.NewHTTPError(http.StatusConflict, handler.Conflict(""))
	}

	roleData = di.ProviderRoleService.InsertRole(*role)
	if roleData == nil {
		return echo.NewHTTPError(http.StatusInternalServerError, handler.InternalServerError(""))
	}

	c.Response().Header().Add("Location", infra.GetHostName() + "/v1/roles/" + role.Uuid.String())
	return c.JSON(http.StatusCreated, roleData)
}