package domain

import (
	"testing"
	"github.com/tomoyane/grant-n-z/di"
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
	di.InitUserService(stub.UserRepositoryStub{})

	code := m.Run()

	fmt.Println("Done")

	os.Exit(code)
}

func TestEncryptPw(t *testing.T) {
	password := "test"
	passwordHash := di.ProviderUserService.EncryptPw(password)

	assert.Equal(t, true, di.ProviderUserService.ComparePw(passwordHash, password))
}

func TestGetUserByEmail(t *testing.T) {
	email := "test@gmail.com"
	userData := di.ProviderUserService.GetUserByEmail(email)

	assert.Equal(t, correctUserName, userData.Username)
}

func TestInsertUser(t *testing.T) {
	user := entity.User {
		Username: "test",
		Email: "test@gmail.com",
		Password: "21312abcdefg",
	}

	insertUser := di.ProviderUserService.InsertUser(user)
	isPassword := di.ProviderUserService.ComparePw(insertUser.Password, correctPassword)

	assert.Equal(t, correctUserName, insertUser.Username)
	assert.Equal(t, correctEmail, insertUser.Email)
	assert.Equal(t, true, isPassword)
}
