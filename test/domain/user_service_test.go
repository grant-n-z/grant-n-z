package domain

import (
	"testing"
	"github.com/tomoyane/grant-n-z/common"
	"github.com/tomoyane/grant-n-z/test/stub"
	"fmt"
	"os"
	"src/github.com/stretchr/testify/assert"
	"github.com/tomoyane/grant-n-z/domain/entity"
)

var(
	correctUserName = "test"
	correctEmail = "test@gmail.com"
	correctPassword = "21312abcdefg"
)

func TestMain(m *testing.M) {
	common.InitUserService(stub.UserRepositoryStub{})

	code := m.Run()

	fmt.Println("Done")

	os.Exit(code)
}

func TestEncryptPw(t *testing.T) {
	password := "test"
	passwordHash := common.ProvideUserService.EncryptPw(password)

	assert.Equal(t, common.ProvideUserService.ComparePw(passwordHash, password), true)
}

func TestGetUserByEmail(t *testing.T) {
	email := "test@gmail.com"
	userData := common.ProvideUserService.GetUserByEmail(email)

	assert.Equal(t, userData.Username, correctUserName)
}

func TestInsertUser(t *testing.T) {
	user := entity.User {
		Username: "test",
		Email: "test@gmail.com",
		Password: "21312abcdefg",
	}

	insertUser := common.ProvideUserService.InsertUser(user)

	assert.Equal(t, insertUser.Username, correctUserName)
	assert.Equal(t, insertUser.Email, correctEmail)
	assert.Equal(t, insertUser.Password, correctPassword)
}
