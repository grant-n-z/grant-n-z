package controller

import (
	"github.com/labstack/echo"
	"github.com/tomoyane/grant-n-z/domain/entity"
	"net/http"
	"github.com/tomoyane/grant-n-z/domain"
	"github.com/satori/go.uuid"
	"strings"
	"github.com/tomoyane/grant-n-z/di"
	"github.com/tomoyane/grant-n-z/infra"
)

func PostUser(c echo.Context) (err error) {
	user := new(entity.User)

	if err = c.Bind(user); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest,
			domain.ErrorResponse{}.Error(http.StatusBadRequest, "001"))
	}

	user.Uuid = uuid.NewV4()
	user.Password = di.ProviderUserService.EncryptPw(user.Password)

	if err = c.Validate(user); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest,
			domain.ErrorResponse{}.Error(http.StatusBadRequest, "002"))
	}

	userData := di.ProviderUserService.GetUserByEmail(user.Email)
	if userData == nil {
		return echo.NewHTTPError(http.StatusInternalServerError,
			domain.ErrorResponse{}.Error(http.StatusInternalServerError, "003"))
	}

	if !strings.EqualFold(userData.Email, "") {
		return echo.NewHTTPError(http.StatusUnprocessableEntity,
			domain.ErrorResponse{}.Error(http.StatusUnprocessableEntity, "004"))
	}

	if di.ProviderUserService.InsertUser(*user) == nil {
		return echo.NewHTTPError(http.StatusInternalServerError,
			domain.ErrorResponse{}.Error(http.StatusInternalServerError, "005"))
	}

	token := di.ProviderTokenService.GenerateJwt(user.Username, user.Uuid)
	refreshToken := di.ProviderTokenService.GenerateJwt(user.Username, user.Uuid)

	if di.ProviderTokenService.InsertToken(user.Uuid, token, refreshToken) == nil {
		return echo.NewHTTPError(http.StatusInternalServerError,
			domain.ErrorResponse{}.Error(http.StatusInternalServerError, "006"))
	}

	success := map[string]string {
		"message": "user creation succeeded.",
	}

	c.Response().Header().Add("Location", infra.GetHostName() + "/v1/users/" + user.Uuid.String())
	return c.JSON(http.StatusCreated, success)
}