package repository

import (
	"github.com/tomoyane/grant-n-z/gserver/entity"
	"github.com/tomoyane/grant-n-z/gserver/model"
)

type UserServiceRepository interface {
	FindAll() ([]*entity.UserService, *model.ErrorResBody)

	FindById(id int) ([]*entity.UserService, *model.ErrorResBody)

	FindByUserId(userId int) ([]*entity.UserService, *model.ErrorResBody)

	Save(userService entity.UserService) (*entity.UserService, *model.ErrorResBody)
}
