package controller

import (
	"github.com/labstack/echo"
	"net/http"
	"github.com/tomoyane/grant-n-z/di"
	"github.com/tomoyane/grant-n-z/domain/entity"
	"github.com/tomoyane/grant-n-z/handler"
)

func PostToken(c echo.Context) (err error) {
	user := new(entity.User)

	if err = c.Bind(user); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, handler.BadRequest("007"))
	}

	user.Username = user.Email
	if err = c.Validate(user); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, handler.BadRequest("008"))
	}

	token, errToken := di.ProviderTokenService.IssueToken(user)

	if errToken != nil {
		return echo.NewHTTPError(errToken.Code, errToken)
	}

	return c.JSON(http.StatusOK,  map[string]string {"token": token.Token})
}