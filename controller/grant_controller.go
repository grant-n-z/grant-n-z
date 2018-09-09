package controller

import (
	"github.com/labstack/echo"
	"net/http"
	"github.com/tomoyane/grant-n-z/di"
	"github.com/tomoyane/grant-n-z/domain"
	"strings"
)

func GrantToken(c echo.Context) (err error) {
	token := c.Request().Header.Get("Authorization")
	if token == "" {
		return echo.NewHTTPError(http.StatusUnauthorized,
			domain.ErrorResponse{}.Error(http.StatusUnauthorized, "013"))
	}

	resultMap, result := di.ProviderTokenService.ValidJwt(token)
	if !result {
		return echo.NewHTTPError(http.StatusUnauthorized,
			domain.ErrorResponse{}.Error(http.StatusUnauthorized, "014"))
	}

	user := di.ProviderUserService.GetUserByUuid(resultMap["username"], resultMap["user_uuid"])
	if user == nil {
		return echo.NewHTTPError(http.StatusInternalServerError,
			domain.ErrorResponse{}.Error(http.StatusInternalServerError, "015"))
	}

	if len(user.Email) > 0 {
		return echo.NewHTTPError(http.StatusUnauthorized,
			domain.ErrorResponse{}.Error(http.StatusUnauthorized, "016"))
	}

	role := di.ProviderRoleService.GetRoleByUserUuid(user.Uuid.String())
	if user == nil {
		return echo.NewHTTPError(http.StatusInternalServerError,
			domain.ErrorResponse{}.Error(http.StatusInternalServerError, "017"))
	}

	if len(role.UserUuid) > 0 {
		return echo.NewHTTPError(http.StatusForbidden,
			domain.ErrorResponse{}.Error(http.StatusForbidden, "018"))
	}

	if !strings.EqualFold(role.Type, "user") || !strings.EqualFold(role.Type, "admin") {
		return echo.NewHTTPError(http.StatusForbidden,
			domain.ErrorResponse{}.Error(http.StatusForbidden, "019"))
	}

	return c.JSON(http.StatusOK, true)
}
