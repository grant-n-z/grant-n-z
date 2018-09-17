package domain

import (
	"github.com/tomoyane/grant-n-z/di"
	"github.com/tomoyane/grant-n-z/test/stub"
	"fmt"
	"os"
	"testing"
	"github.com/satori/go.uuid"
)

var (
	username = "test"
	userUuid, _ = uuid.FromString("52F6228E-9169-4563-ADE2-07ED697B67BA")
	token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9"
	correctUserName = "test"
	correctEmail = "test@gmail.com"
	correctPassword = "21312abcdefg"
	correctUserUuid = "52F6228E-9169-4563-ADE2-07ED697B67BA"
)

func TestMain(m *testing.M) {
	os.Setenv("ENV", "test")

	di.InitUserService(stub.UserRepositoryStub{})
	di.InitTokenService(stub.TokenRepositoryStub{})
	di.InitRoleService(stub.RoleRepositoryStub{})

	code := m.Run()

	fmt.Println("Done domain package")

	os.Exit(code)
}