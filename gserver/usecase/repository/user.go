package repository

import (
	"github.com/tomoyane/grant-n-z/gserver/entity"
	"github.com/tomoyane/grant-n-z/gserver/model"
)

type UserRepository interface {
	FindById(id int) (*entity.User, *model.ErrorResponse)

	FindByEmail(email string) (*entity.User, *model.ErrorResponse)

	FindUserWithRoleByEmail(email string) (*model.UserOperatorMemberRole, *model.ErrorResponse)

	Save(user entity.User) (*entity.User, *model.ErrorResponse)

	Update(user entity.User) (*entity.User, *model.ErrorResponse)
}
