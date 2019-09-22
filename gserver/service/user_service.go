package service

import (
	"github.com/tomoyane/grant-n-z/gserver/entity"
	"github.com/tomoyane/grant-n-z/gserver/model"
)

type UserServiceService interface {
	Get(queryParam string) (interface{}, *model.ErrorResBody)

	GetUserServices() ([]*entity.UserService, *model.ErrorResBody)

	GetUserServicesByUserId(userId int) ([]*entity.UserService, *model.ErrorResBody)

	InsertUserService(userService *entity.UserService) (*entity.UserService, *model.ErrorResBody)
}
