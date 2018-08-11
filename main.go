package main

import (
	"github.com/labstack/echo"
	"github.com/tomoyane/grant-n-z/common"
	"github.com/tomoyane/grant-n-z/controller"
	"gopkg.in/go-playground/validator.v9"
	"github.com/tomoyane/grant-n-z/domain"
)

func main() {
	common.InitDB()

	e := echo.New()
	e.Validator = &domain.GrantValidator{Validator: validator.New()}
	e.POST("/v1/users", controller.UserController)
	e.Logger.Fatal(e.Start(":8080"))
}