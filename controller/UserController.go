package controller

import (
	"github.com/labstack/echo"
	"github.com/tomoyane/grant-n-z/domain/entity"
	"github.com/satori/go.uuid"
	"github.com/tomoyane/grant-n-z/domain/service"
	"net/http"
	"github.com/tomoyane/grant-n-z/domain"
	"strings"
)

func GenerateUser(c echo.Context) (err error) {
	user := new(entity.User)

	if err = c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest,
			domain.ErrorResponse{}.Error(http.StatusBadRequest, "001"))
	}

	user.Uuid = uuid.NewV4()
	user.Password = service.EncryptPw(user.Password)

	if err = c.Validate(user); err != nil {
		return c.JSON(http.StatusBadRequest,
			domain.ErrorResponse{}.Error(http.StatusBadRequest, "002"))
	}

	userData := service.GetUserByEmail(user.Email)
	if userData == nil {
		return c.JSON(http.StatusInternalServerError,
			domain.ErrorResponse{}.Error(http.StatusInternalServerError, "003"))
	}

	if !strings.EqualFold(userData.Email, "") {
		return c.JSON(http.StatusUnprocessableEntity,
			domain.ErrorResponse{}.Error(http.StatusUnprocessableEntity, "004"))
	}

	if service.InsertUser(*user) == nil {
		return c.JSON(http.StatusInternalServerError,
			domain.ErrorResponse{}.Error(http.StatusInternalServerError, "005"))
	}

	success := map[string]string {
		"message": "user creation succeeded.",
	}

	return c.JSON(http.StatusCreated, success)
}