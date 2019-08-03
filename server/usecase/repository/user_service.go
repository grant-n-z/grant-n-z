package repository

import (
	"github.com/tomoyane/grant-n-z/server/entity"
	"github.com/tomoyane/grant-n-z/server/model"
)

type UserServiceRepository interface {
	FindAll() ([]*entity.UserService, *model.ErrorResponse)

	FindById(id int) ([]*entity.UserService, *model.ErrorResponse)

	FindByUserId(userId int) ([]*entity.UserService, *model.ErrorResponse)

	Save(userService entity.UserService) (*entity.UserService, *model.ErrorResponse)
}
