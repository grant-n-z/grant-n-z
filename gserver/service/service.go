package service

import (
	"strings"

	"github.com/google/uuid"

	"github.com/tomoyane/grant-n-z/gserver/common/ctx"
	"github.com/tomoyane/grant-n-z/gserver/common/driver"
	"github.com/tomoyane/grant-n-z/gserver/data"
	"github.com/tomoyane/grant-n-z/gserver/entity"
	"github.com/tomoyane/grant-n-z/gserver/log"
	"github.com/tomoyane/grant-n-z/gserver/model"
)

var sInstance Service

type serviceImpl struct {
	serviceRepository     data.ServiceRepository
	userServiceRepository data.UserServiceRepository
}

type Service interface {
	// Get service
	GetServices() ([]*entity.Service, *model.ErrorResBody)

	// Get service by service id
	GetServiceById(id int) (*entity.Service, *model.ErrorResBody)

	// Get service by service name
	GetServiceByName(name string) (*entity.Service, *model.ErrorResBody)

	// Get service by service api key
	GetServiceOfApiKey() (*entity.Service, *model.ErrorResBody)

	// Get service of user
	GetServiceOfUser() ([]*entity.Service, *model.ErrorResBody)

	// Insert service
	InsertService(service *entity.Service) (*entity.Service, *model.ErrorResBody)
}

// Get Policy instance.
// If use singleton pattern, call this instance method
func GetServiceInstance() Service {
	if sInstance == nil {
		sInstance = NewServiceService()
	}
	return sInstance
}

// Constructor
func NewServiceService() Service {
	log.Logger.Info("New `Service` instance")
	log.Logger.Info("Inject `ServiceRepository`, `UserServiceRepository` to `Service`")
	return serviceImpl{
		serviceRepository:     data.ServiceRepositoryImpl{Db: driver.Db},
		userServiceRepository: data.UserServiceRepositoryImpl{Db: driver.Db},
	}
}

func (ss serviceImpl) GetServices() ([]*entity.Service, *model.ErrorResBody) {
	return ss.serviceRepository.FindAll()
}

func (ss serviceImpl) GetServiceById(id int) (*entity.Service, *model.ErrorResBody) {
	return ss.serviceRepository.FindById(id)
}

func (ss serviceImpl) GetServiceByName(name string) (*entity.Service, *model.ErrorResBody) {
	return ss.serviceRepository.FindByName(name)
}

func (ss serviceImpl) GetServiceOfApiKey() (*entity.Service, *model.ErrorResBody) {
	service, err := ss.serviceRepository.FindByApiKey(ctx.GetApiKey().(string))
	if service == nil || err != nil {
		err := model.BadRequest("Api-Key is invalid")
		return nil, err
	}

	return service, nil
}

func (ss serviceImpl) GetServiceOfUser() ([]*entity.Service, *model.ErrorResBody) {
	return ss.userServiceRepository.FindByUserId(ctx.GetUserId().(int))
}

func (ss serviceImpl) InsertService(service *entity.Service) (*entity.Service, *model.ErrorResBody) {
	service.Uuid = uuid.New()
	key := uuid.New()
	service.ApiKey = strings.Replace(key.String(), "-", "", -1)
	return ss.serviceRepository.Save(*service)
}
