package main

import (
	"github.com/tomoyane/grant-n-z/common"
	"github.com/tomoyane/grant-n-z/domain/repository"
	"github.com/labstack/echo"
	"github.com/tomoyane/grant-n-z/domain"
	"gopkg.in/go-playground/validator.v9"
	"github.com/tomoyane/grant-n-z/controller"
	"github.com/tomoyane/grant-n-z/infra"
)

func main() {
	infra.InitDB()
	common.InitUserService(repository.UserRepositoryImpl{})

	e := echo.New()
	e.Validator = &domain.GrantValidator{Validator: validator.New()}
	e.POST("/v1/users", controller.GenerateUser)
	e.Logger.Fatal(e.Start(":8080"))
}