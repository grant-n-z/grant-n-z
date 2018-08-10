package controller

import (
	"github.com/labstack/echo"
	"net/http"
	"github.com/tomoyane/grant-n-z/domain/entity"
	"github.com/tomoyane/grant-n-z/domain/service"
	"github.com/satori/go.uuid"
	"strings"
	"gopkg.in/go-playground/validator.v9"
	"github.com/tomoyane/grant-n-z/domain"
)

var validate = validator.New()

func UserController (c echo.Context) (err error) {
	user := new(entity.User)
	if err = c.Bind(user); err != nil {
		return
	}

	user.Uuid = uuid.NewV4()
	user.Password = service.EncryptPw(user.Password)

	validateErr := validate.Struct(user)
	if validateErr != nil {
		return c.JSON(http.StatusBadRequest,
			domain.ErrorResponse{}.Error(http.StatusBadRequest, "001"))
	}

	userData := service.GetUserByEmail(user.Email)
	if userData == nil {
		return c.JSON(http.StatusInternalServerError,
			domain.ErrorResponse{}.Error(http.StatusInternalServerError, "002"))
	}

	if !strings.EqualFold(userData.Email, "") {
		return c.JSON(http.StatusUnprocessableEntity,
			domain.ErrorResponse{}.Error(http.StatusUnprocessableEntity, "003"))
	}

	if service.InsertUser(*user) == nil {
		return c.JSON(http.StatusInternalServerError,
			domain.ErrorResponse{}.Error(http.StatusInternalServerError, "004"))
	}

	success := map[string]string {
		"message": "user creation succeeded.",
	}

	return c.JSON(http.StatusOK, success)
}