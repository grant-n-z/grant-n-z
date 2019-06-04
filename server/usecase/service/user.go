package service

import (
	"github.com/tomoyane/grant-n-z/server/entity"
	"github.com/tomoyane/grant-n-z/server/model"
)

type UserService interface {
	EncryptPw(password string) string

	ComparePw(passwordHash string, password string) bool

	GetUserById(id int) (*entity.User, *model.ErrorResponse)

	InsertUser(user *entity.User) (*entity.User, *model.ErrorResponse)
}
