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

	userData := di.ProviderUserService.GetUserByEmail(user.Email)
	if userData == nil {
		return echo.NewHTTPError(http.StatusInternalServerError, handler.InternalServerError("009"))
	}

	if len(userData.Email) == 0 {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, handler.NotFound("010"))
	}

	if !di.ProviderUserService.ComparePw(userData.Password, user.Password) {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, handler.UnProcessableEntity("011"))
	}

	tokenStr := di.ProviderTokenService.GenerateJwt(userData.Username, userData.Uuid, false)
	refreshTokenStr := di.ProviderTokenService.GenerateJwt(userData.Username, userData.Uuid, false)

	if tokenStr == "" || refreshTokenStr == ""{
		return echo.NewHTTPError(http.StatusInternalServerError, handler.InternalServerError("012"))
	}

	token := di.ProviderTokenService.InsertToken(userData.Uuid, tokenStr, refreshTokenStr)
	if token == nil {
		return echo.NewHTTPError(http.StatusInternalServerError, handler.InternalServerError("013"))
	}

	success := map[string]string {
		"token": token.Token,
	}

	return c.JSON(http.StatusOK, success)
}