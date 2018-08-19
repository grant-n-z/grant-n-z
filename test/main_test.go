package test

import (
	"testing"
	"os"
	"github.com/tomoyane/grant-n-z/common"
	"github.com/tomoyane/grant-n-z/test/stub"
	"fmt"
	"github.com/labstack/echo"
	"github.com/tomoyane/grant-n-z/domain"
	"gopkg.in/go-playground/validator.v9"
)

func TestMain(m *testing.M) {
	common.InitUserService(stub.UserRepositoryStub{})
	e := echo.New()
	e.Validator = &domain.GrantValidator{Validator: validator.New()}

	code := m.Run()

	os.Exit(code)
	fmt.Print("Done")
}
