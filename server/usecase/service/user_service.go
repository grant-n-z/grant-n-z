package service

import (
	"github.com/tomoyane/grant-n-z/server/entity"
)

type UserServiceService interface {
	Get(queryParam string) (interface{}, *entity.ErrorResponse)

	GetUserServices() ([]*entity.UserService, *entity.ErrorResponse)

	GetUserServicesByUserId(userId int) ([]*entity.UserService, *entity.ErrorResponse)

	InsertUserService(userService *entity.UserService) (*entity.UserService, *entity.ErrorResponse)
}
