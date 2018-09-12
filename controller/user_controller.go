package controller

import (
	"github.com/labstack/echo"
	"github.com/tomoyane/grant-n-z/domain/entity"
	"net/http"
	"github.com/tomoyane/grant-n-z/domain"
	"github.com/tomoyane/grant-n-z/di"
	"github.com/tomoyane/grant-n-z/infra"
)

func PostUser(c echo.Context) (err error) {
	user := new(entity.User)

	if err = c.Bind(user); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest,
			domain.ErrorResponse{}.Error(http.StatusBadRequest, "001"))
	}

	if err = c.Validate(user); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest,
			domain.ErrorResponse{}.Error(http.StatusBadRequest, "002"))
	}

	userData := di.ProviderUserService.GetUserByEmail(user.Email)
	if userData == nil {
		return echo.NewHTTPError(http.StatusInternalServerError,
			domain.ErrorResponse{}.Error(http.StatusInternalServerError, "003"))
	}

	if len(userData.Email) > 0 {
		return echo.NewHTTPError(http.StatusUnprocessableEntity,
			domain.ErrorResponse{}.Error(http.StatusUnprocessableEntity, "004"))
	}

	userData = di.ProviderUserService.InsertUser(*user)
	if userData == nil {
		infra.Db.Rollback()
		return echo.NewHTTPError(http.StatusInternalServerError,
			domain.ErrorResponse{}.Error(http.StatusInternalServerError, "005"))
	}

	if di.ProviderRoleService.InsertRole(userData.Uuid) == nil {
		infra.Db.Rollback()
		return echo.NewHTTPError(http.StatusInternalServerError,
			domain.ErrorResponse{}.Error(http.StatusInternalServerError, "006"))
	}

	success := map[string]string {
		"message": "user creation succeeded.",
	}

	c.Response().Header().Add("Location", infra.GetHostName() + "/v1/users/" + user.Uuid.String())
	return c.JSON(http.StatusCreated, success)
}