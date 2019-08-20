package service

import (
	"github.com/tomoyane/grant-n-z/gserver/entity"
	"github.com/tomoyane/grant-n-z/gserver/model"
)

type UserService interface {
	EncryptPw(password string) string

	ComparePw(passwordHash string, password string) bool

	GetUserById(id int) (*entity.User, *model.ErrorResponse)

	GetUserByEmail(email string) (*entity.User, *model.ErrorResponse)

	GetUserWithRoleByEmail(email string) (*model.UserOperatorMemberRole, *model.ErrorResponse)

	InsertUser(user *entity.User) (*entity.User, *model.ErrorResponse)

	UpdateUser(user *entity.User) (*entity.User, *model.ErrorResponse)

	GenerateJwt(user *entity.User, roleId int) *string

	ParseJwt(token string) (map[string]string, bool)
}
