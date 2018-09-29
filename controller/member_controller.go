package controller

import (
	"github.com/labstack/echo"
	"github.com/tomoyane/grant-n-z/di"
	"github.com/tomoyane/grant-n-z/domain/entity"
	"github.com/tomoyane/grant-n-z/handler"
	"github.com/tomoyane/grant-n-z/infra"
	"net/http"
)

func PostMember(c echo.Context) (err error) {
	token := c.Request().Header.Get("Authorization")
	errAuth := di.ProvideTokenService.VerifyToken(c, token)
	if errAuth != nil {
		return echo.NewHTTPError(errAuth.Code, errAuth)
	}

	member := new(entity.Member)
	if err = c.Bind(member); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, handler.BadRequest(""))
	}

	if err = c.Validate(member); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, handler.BadRequest(""))
	}

	memberData := di.ProvideMemberService.Insert(*member)

	c.Response().Header().Add("Location", infra.GetHostName() + "/v1/members/" + memberData.Uuid.String())
	return c.JSON(http.StatusCreated, memberData)
}
