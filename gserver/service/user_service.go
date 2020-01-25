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
	// Get all user_services
	GetUserServices() ([]*entity.UserService, *model.ErrorResBody)

	// Get user_services by user_id and service_id
	GetUserServiceByUserIdAndServiceId(userId int, serviceId int) (*entity.UserService, *model.ErrorResBody)

	// Insert user_service
	InsertUserService(userServiceEntity entity.UserService) (*entity.UserService, *model.ErrorResBody)
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

func (uss userServiceServiceImpl) InsertUserService(userServiceEntity entity.UserService) (*entity.UserService, *model.ErrorResBody) {
	userService, err := uss.userServiceRepository.FindByUserIdAndServiceId(userServiceEntity.UserId, userServiceEntity.ServiceId)
	if err != nil || userService != nil {
		return nil, model.Conflict("Already the user has this service account")
	}
	return uss.userServiceRepository.Save(userServiceEntity)
}
