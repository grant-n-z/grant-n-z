package handler

import (
	"fmt"
	"github.com/labstack/echo"
	"github.com/tomoyane/grant-n-z/server/di"
	"github.com/tomoyane/grant-n-z/server/domain"
	"github.com/tomoyane/grant-n-z/server/domain/repository"
	"gopkg.in/go-playground/validator.v9"
	"os"
	"testing"
)

var(
	e = echo.New()
)

func TestMain(m *testing.M) {
	os.Setenv("ENV", "test")

	di.InitUserService(repository.UserRepositoryStub{})
	di.InitTokenService(repository.TokenRepositoryStub{})
	di.InitRoleService(repository.RoleRepositoryStub{})

	e.Validator = &domain.GrantValidator{Validator: validator.New()}

	code := m.Run()

	fmt.Println("Done controller packge")

	os.Exit(code)
}



