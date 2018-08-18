package test

import (
	"testing"
	"os"
	"github.com/tomoyane/grant-n-z/common"
	"github.com/tomoyane/grant-n-z/test/stub"
)

func TestMain(m *testing.M) {
	common.InitUserService(stub.UserRepositoryStub{})
	code := m.Run()
	os.Exit(code)
}
