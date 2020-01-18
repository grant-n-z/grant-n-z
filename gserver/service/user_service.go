package service

import (
	"github.com/tomoyane/grant-n-z/gserver/common/driver"
	"github.com/tomoyane/grant-n-z/gserver/data"
	"github.com/tomoyane/grant-n-z/gserver/entity"
	"github.com/tomoyane/grant-n-z/gserver/log"
	"github.com/tomoyane/grant-n-z/gserver/model"
)

var ussInstance UserServiceService

type UserServiceService interface {
	GetUserServices() ([]*entity.UserService, *model.ErrorResBody)

	GetUserServiceByUserIdAndServiceId(userId int, serviceId int) (*entity.UserService, *model.ErrorResBody)
}

type userServiceServiceImpl struct {
	userServiceRepository data.UserServiceRepository
	userRepository        data.UserRepository
	serviceRepository     data.ServiceRepository
}

func GetUserServiceServiceInstance() UserServiceService {
	if ussInstance == nil {
		ussInstance = NewUserServiceService()
	}
	return ussInstance
}

func NewUserServiceService() UserServiceService {
	log.Logger.Info("New `UserServiceService` instance")
	return userServiceServiceImpl{
		userServiceRepository: data.GetUserServiceRepositoryInstance(driver.Db),
		userRepository:        data.GetUserRepositoryInstance(driver.Db),
		serviceRepository:     data.GetServiceRepositoryInstance(driver.Db),
	}
}

func (uss userServiceServiceImpl) GetUserServices() ([]*entity.UserService, *model.ErrorResBody) {
	return uss.userServiceRepository.FindAll()
}

func (uss userServiceServiceImpl) GetUserServiceByUserIdAndServiceId(userId int, serviceId int) (*entity.UserService, *model.ErrorResBody) {
	return uss.userServiceRepository.FindByUserIdAndServiceId(userId, serviceId)
}
