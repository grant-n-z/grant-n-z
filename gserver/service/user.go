package service

import (
	"github.com/tomoyane/grant-n-z/gserver/entity"
	"github.com/tomoyane/grant-n-z/gserver/model"
)

type UserService interface {
	EncryptPw(password string) string

	ComparePw(passwordHash string, password string) bool

	GetUserById(id int) (*entity.User, *model.ErrorResBody)

	GetUserByEmail(email string) (*entity.User, *model.ErrorResBody)

	GetUserWithRoleByEmail(email string) (*model.UserOperatorMemberRole, *model.ErrorResBody)

	InsertUser(user *entity.User) (*entity.User, *model.ErrorResBody)

	InsertUserWithService(user *entity.User, userService *entity.UserService) (*entity.User, *model.ErrorResBody)

	UpdateUser(user *entity.User) (*entity.User, *model.ErrorResBody)

	GenerateJwt(user *entity.User, roleId int) *string

	ParseJwt(token string) (map[string]string, bool)
}
