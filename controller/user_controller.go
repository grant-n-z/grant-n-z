package controller

import (
	"github.com/labstack/echo"
	"github.com/tomoyane/grant-n-z/domain/entity"
	"net/http"
	"github.com/tomoyane/grant-n-z/domain"
	"github.com/satori/go.uuid"
	"strings"
	"github.com/tomoyane/grant-n-z/common"
)

func GenerateUser(c echo.Context) (err error) {
	user := new(entity.User)

	if err = c.Bind(user); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest,
			domain.ErrorResponse{}.Error(http.StatusBadRequest, "001"))
	}

	user.Uuid = uuid.NewV4()
	user.Password = common.ProvideUserService.EncryptPw(user.Password)

	if err = c.Validate(user); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest,
			domain.ErrorResponse{}.Error(http.StatusBadRequest, "002"))
	}

	userData := common.ProvideUserService.GetUserByEmail(user.Email)
	if userData == nil {
		return echo.NewHTTPError(http.StatusInternalServerError,
			domain.ErrorResponse{}.Error(http.StatusInternalServerError, "003"))
	}

	if !strings.EqualFold(userData.Email, "") {
		return echo.NewHTTPError(http.StatusUnprocessableEntity,
			domain.ErrorResponse{}.Error(http.StatusUnprocessableEntity, "004"))
	}

	if common.ProvideUserService.InsertUser(*user) == nil {
		return echo.NewHTTPError(http.StatusInternalServerError,
			domain.ErrorResponse{}.Error(http.StatusInternalServerError, "005"))
	}

	success := map[string]string {
		"message": "user creation succeeded.",
	}

	return c.JSON(http.StatusCreated, success)
}