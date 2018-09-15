package domain

import (
	"github.com/tomoyane/grant-n-z/di"
	"github.com/tomoyane/grant-n-z/test/stub"
	"fmt"
	"os"
	"testing"
)

var(
	correctUserName = "test"
	correctEmail = "test@gmail.com"
	correctPassword = "21312abcdefg"
	correctUserUuid = "52F6228E-9169-4563-ADE2-07ED697B67BA"
)

func TestMain(m *testing.M) {
	di.InitUserService(stub.UserRepositoryStub{})

	code := m.Run()

	fmt.Println("Done domain package")

	os.Exit(code)
}