package service

import (
	"strconv"
	"strings"

	"github.com/tomoyane/grant-n-z/gserver/common/driver"
	"github.com/tomoyane/grant-n-z/gserver/data"
	"github.com/tomoyane/grant-n-z/gserver/entity"
	"github.com/tomoyane/grant-n-z/gserver/log"
	"github.com/tomoyane/grant-n-z/gserver/model"
)

var ussInstance UserServiceService

type UserServiceService interface {
	Get(queryParam string) (interface{}, *model.ErrorResBody)

	GetUserServices() ([]*entity.UserService, *model.ErrorResBody)

	GetUserServicesByUserId(userId int) ([]*entity.UserService, *model.ErrorResBody)

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
	log.Logger.Info("Inject `UserServiceRepository`, `UserRepository`, `ServiceRepository` to `UserServiceService`")
	return userServiceServiceImpl{
		userServiceRepository: data.GetUserServiceRepositoryInstance(driver.Db),
		userRepository:        data.GetUserRepositoryInstance(driver.Db),
		serviceRepository:     data.GetServiceRepositoryInstance(driver.Db),
	}
}

func (uss userServiceServiceImpl) Get(queryParam string) (interface{}, *model.ErrorResBody) {
	var result interface{}

	if strings.EqualFold(queryParam, "") {
		return uss.GetUserServices()
	}

	i, castErr := strconv.Atoi(queryParam)
	if castErr != nil {
		log.Logger.Warn("The user_id is only integer")
		return nil, model.BadRequest(castErr.Error())
	}

	userServiceEntities, err := uss.GetUserServicesByUserId(i)
	if err != nil {
		return nil, err
	}

	if userServiceEntities == nil {
		result = new([]string)
	} else {
		result = userServiceEntities
	}

	return result, nil
}

func (uss userServiceServiceImpl) GetUserServices() ([]*entity.UserService, *model.ErrorResBody) {
	return uss.userServiceRepository.FindAll()
}

func (uss userServiceServiceImpl) GetUserServicesByUserId(userId int) ([]*entity.UserService, *model.ErrorResBody) {
	return uss.userServiceRepository.FindByUserId(userId)
}

func (uss userServiceServiceImpl) GetUserServiceByUserIdAndServiceId(userId int, serviceId int) (*entity.UserService, *model.ErrorResBody) {
	return uss.userServiceRepository.FindByUserIdAndServiceId(userId, serviceId)
}
