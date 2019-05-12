package repository

import "github.com/tomoyane/grant-n-z/server/entity"

type UserServiceRepository interface {
	FindAll() ([]*entity.UserService, *entity.ErrorResponse)

	FindByUserId(userId int) ([]*entity.UserService, *entity.ErrorResponse)

	Save(userService entity.UserService) (*entity.UserService, *entity.ErrorResponse)
}
