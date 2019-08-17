package service

import (
	"github.com/tomoyane/grant-n-z/gserver/entity"
	"github.com/tomoyane/grant-n-z/gserver/model"
)

type UserServiceService interface {
	Get(queryParam string) (interface{}, *model.ErrorResponse)

	GetUserServices() ([]*entity.UserService, *model.ErrorResponse)

	GetUserServicesByUserId(userId int) ([]*entity.UserService, *model.ErrorResponse)

	InsertUserService(userService *entity.UserService) (*entity.UserService, *model.ErrorResponse)
}
