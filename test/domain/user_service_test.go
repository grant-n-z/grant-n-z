package domain

import (
	"testing"
	"github.com/tomoyane/grant-n-z/di"
	"github.com/tomoyane/grant-n-z/domain/entity"
	"github.com/stretchr/testify/assert"
	"github.com/satori/go.uuid"
)

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

func TestGetUserByUuid(t *testing.T) {
	username := "test"
	userUuidStr := "52F6228E-9169-4563-ADE2-07ED697B67BA"
	userData := di.ProviderUserService.GetUserByUuid(username, userUuidStr)

	correctUserUuid, _ := uuid.FromString(correctUserUuid)
	assert.Equal(t, correctUserName, userData.Username)
	assert.Equal(t, correctUserUuid, userData.Uuid)
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
