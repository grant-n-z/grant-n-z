package controller

import (
	"github.com/labstack/echo"
	"github.com/tomoyane/grant-n-z/domain/entity"
	"net/http"
	"github.com/tomoyane/grant-n-z/di"
	"github.com/tomoyane/grant-n-z/infra"
)

func PostUser(c echo.Context) (err error) {
	user := new(entity.User)

	result := di.ProviderUserService.PostUserData(c, user)

	if result != nil {
		return result
	}

	success := map[string]string {
		"message": "user creation succeeded.",
	}

	c.Response().Header().Add("Location", infra.GetHostName() + "/v1/users/" + user.Uuid.String())
	return c.JSON(http.StatusCreated, success)
}

func PutUser(c echo.Context) (err error) {
	token := c.Request().Header.Get("Authorization")
	column := c.Param("column")
	user := new(entity.User)

	result := di.ProviderUserService.PutUserColumnData(
		c, di.ProviderTokenService, di.ProviderRoleService, user, token, column)

	if result != nil {
		return result
	}

	success := map[string]string {
		"message": "ok.",
	}

	return c.JSON(http.StatusOK, success)
}
