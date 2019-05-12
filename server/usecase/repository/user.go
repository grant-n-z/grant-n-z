package repository

import (
	"github.com/tomoyane/grant-n-z/server/entity"
)

type UserRepository interface {
	FindById(id int) (*entity.User, *entity.ErrorResponse)

	Save(user entity.User) (*entity.User, *entity.ErrorResponse)

	Update(user entity.User) *entity.User
}
