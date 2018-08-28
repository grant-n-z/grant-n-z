package controller

import (
	"github.com/labstack/echo"
	"net/http"
	"github.com/tomoyane/grant-n-z/di"
	"github.com/tomoyane/grant-n-z/domain"
	"github.com/tomoyane/grant-n-z/domain/entity"
)

func PostToken(c echo.Context) (err error) {
	user := new(entity.User)

	if err = c.Bind(user); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest,
			domain.ErrorResponse{}.Error(http.StatusBadRequest, "006"))
	}

	user.Username = user.Email
	if err = c.Validate(user); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest,
			domain.ErrorResponse{}.Error(http.StatusBadRequest, "007"))
	}

	userData := di.ProviderUserService.GetUserByEmail(user.Email)
	if userData == nil {
		return echo.NewHTTPError(http.StatusInternalServerError,
			domain.ErrorResponse{}.Error(http.StatusInternalServerError, "008"))
	}

	if len(userData.Email) == 0 {
		return echo.NewHTTPError(http.StatusUnprocessableEntity,
			domain.ErrorResponse{}.Error(http.StatusUnprocessableEntity, "009"))
	}

	if !di.ProviderUserService.ComparePw(userData.Password, user.Password) {
		return echo.NewHTTPError(http.StatusUnprocessableEntity,
			domain.ErrorResponse{}.Error(http.StatusUnprocessableEntity, "010"))
	}

	tokenStr := di.ProviderTokenService.GenerateJwt(userData.Username, userData.Uuid)
	refreshTokenStr := di.ProviderTokenService.GenerateJwt(userData.Username, userData.Uuid)

	token := di.ProviderTokenService.InsertToken(userData.Uuid, tokenStr, refreshTokenStr)
	if token == nil {
		return echo.NewHTTPError(http.StatusInternalServerError,
			domain.ErrorResponse{}.Error(http.StatusInternalServerError, "011"))
	}

	success := map[string]string {
		"token": token.Token,
	}

	return c.JSON(http.StatusOK, success)
}