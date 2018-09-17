package controller

import (
	"github.com/labstack/echo"
	"github.com/tomoyane/grant-n-z/di"
	"github.com/tomoyane/grant-n-z/test/stub"
	"github.com/tomoyane/grant-n-z/domain"
	"gopkg.in/go-playground/validator.v9"
	"fmt"
	"os"
	"testing"
)

var(
	e = echo.New()
)

func TestMain(m *testing.M) {
	os.Setenv("ENV", "test")

	di.InitUserService(stub.UserRepositoryStub{})
	di.InitTokenService(stub.TokenRepositoryStub{})
	di.InitRoleService(stub.RoleRepositoryStub{})

	e.Validator = &domain.GrantValidator{Validator: validator.New()}

	code := m.Run()

	fmt.Println("Done controller packge")

	os.Exit(code)
}



