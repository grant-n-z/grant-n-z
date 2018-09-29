package controller

import (
	"github.com/labstack/echo"
	"github.com/tomoyane/grant-n-z/di"
	"github.com/tomoyane/grant-n-z/domain/entity"
	"github.com/tomoyane/grant-n-z/handler"
	"github.com/tomoyane/grant-n-z/infra"
	"net/http"
	"strings"
)

func PostUser(c echo.Context) (err error) {
	user := new(entity.User)
	if err := c.Bind(user); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, handler.BadRequest(""))
	}

	if err := c.Validate(user); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, handler.BadRequest(""))
	}

	errUser := di.ProviderUserService.PostUserData(user)
	if errUser != nil {
		return echo.NewHTTPError(errUser.Code, errUser)
	}

	c.Response().Header().Add("Location", infra.GetHostName() + "/v1/users/" + user.Uuid.String())
	return c.JSON(http.StatusCreated, map[string]string {"message": "user creation succeeded."})
}

func PutUser(c echo.Context) (err error) {
	token := c.Request().Header.Get("Authorization")
	column := c.Param("column")
	errAuth := di.ProviderTokenService.VerifyToken(c, token)

	if errAuth != nil {
		return echo.NewHTTPError(errAuth.Code, errAuth)
	}

	user := new(entity.User)
	if !strings.Contains(column, "username") &&
		!strings.EqualFold(column, "email") && !strings.EqualFold(column, "password") {

		return echo.NewHTTPError(http.StatusBadRequest, handler.BadRequest(""))
	}

	if err := c.Bind(user); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, handler.BadRequest(""))
	}

	if err := c.Validate(user); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, handler.BadRequest(""))
	}

	errUser := di.ProviderUserService.PutUserColumnData(user, column)
	if errUser != nil {
		return echo.NewHTTPError(errUser.Code, errUser)
	}

	return c.JSON(http.StatusOK, map[string]string {"message": "ok."})
}
