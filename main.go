package main

import (
	"github.com/labstack/echo"
	"github.com/tomoyane/grant-n-z/domain"
	"gopkg.in/go-playground/validator.v9"
	"github.com/tomoyane/grant-n-z/controller"
	"github.com/tomoyane/grant-n-z/infra"
	"github.com/tomoyane/grant-n-z/di"
	"github.com/tomoyane/grant-n-z/domain/repository"
)

func main() {
	di.InitUserService(repository.UserRepositoryImpl{})
	di.InitTokenService(repository.TokenRepositoryImpl{}, repository.UserRepositoryImpl{})
	di.InitRoleService(repository.RoleRepositoryImpl{})
	di.InitPrincipalService(repository.PrincipalRepositoryImpl{})

	infra.InitYaml()
	infra.InitDB()
	infra.MigrateDB()

	e := echo.New()
	e.Validator = &domain.GrantValidator{Validator: validator.New()}

	e.POST("/v1/users", controller.PostUser)
	e.PUT("/v1/users/:column", controller.PutUser)

	e.POST("/v1/roles", controller.PostRole)
	e.POST("/v1/principals", controller.PostPrincipal)

	e.POST("/v1/tokens", controller.PostToken)
	e.POST("/v1/grants", controller.GrantToken)
	e.Logger.Fatal(e.Start(":8080"))
}