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
	errUser := di.ProviderUserService.PutUserColumnData(c, user, column)

	if errUser != nil {
		return echo.NewHTTPError(errUser.Code, errUser)
	}

	return c.JSON(http.StatusOK, map[string]string {"message": "ok."})
}
