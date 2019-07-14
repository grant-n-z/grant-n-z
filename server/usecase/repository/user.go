package repository

import (
	"github.com/tomoyane/grant-n-z/server/entity"
	"github.com/tomoyane/grant-n-z/server/model"
)

type UserRepository interface {
	FindById(id int) (*entity.User, *model.ErrorResponse)

	FindByEmail(email string) (*entity.User, *model.ErrorResponse)

	Save(user entity.User) (*entity.User, *model.ErrorResponse)

	Update(user entity.User) *entity.User
}
