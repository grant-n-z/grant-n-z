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
)

func TestMain(m *testing.M) {
	di.InitUserService(stub.UserRepositoryStub{})

	code := m.Run()

	fmt.Println("Done domain package")

	os.Exit(code)
}