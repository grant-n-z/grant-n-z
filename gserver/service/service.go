package service

import (
	"strings"

	"github.com/satori/go.uuid"

	"github.com/tomoyane/grant-n-z/gserver/common/driver"
	"github.com/tomoyane/grant-n-z/gserver/entity"
	"github.com/tomoyane/grant-n-z/gserver/log"
	"github.com/tomoyane/grant-n-z/gserver/model"
	"github.com/tomoyane/grant-n-z/gserver/repository"
)

var sInstance Service

type serviceImpl struct {
	serviceRepository repository.ServiceRepository
}

type Service interface {
	Get(queryParam string) (interface{}, *model.ErrorResBody)

	GetServices() ([]*entity.Service, *model.ErrorResBody)

	GetServiceById(id int) (*entity.Service, *model.ErrorResBody)

	GetServiceByName(name string) (*entity.Service, *model.ErrorResBody)

	GetServiceByApiKey(apiKey string) (*entity.Service, *model.ErrorResBody)

	InsertService(service *entity.Service) (*entity.Service, *model.ErrorResBody)
}

func GetServiceInstance() Service {
	if sInstance == nil {
		sInstance = NewServiceService()
	}
	return sInstance
}

func NewServiceService() Service {
	log.Logger.Info("New `Service` instance")
	log.Logger.Info("Inject `ServiceRepository` to `Service`")
	return serviceImpl{
		serviceRepository: repository.ServiceRepositoryImpl{Db: driver.Db},
	}
}

func (ss serviceImpl) Get(queryParam string) (interface{}, *model.ErrorResBody) {
	if !strings.EqualFold(queryParam, "") {
		serviceEntity, err := ss.GetServiceByName(queryParam)
		if err != nil {
			return nil, err
		}

		if serviceEntity == nil {
			return entity.Service{}, nil
		}

		return serviceEntity, nil

	} else {
		serviceEntities, err := ss.GetServices()
		if err != nil {
			return nil, err
		}

		if serviceEntities == nil {
			return []entity.Service{}, nil
		}

		return serviceEntities, nil
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

func (ss serviceImpl) GetServiceByApiKey(apiKey string) (*entity.Service, *model.ErrorResBody) {
	return ss.serviceRepository.FindByApiKey(apiKey)
}

func (ss serviceImpl) InsertService(service *entity.Service) (*entity.Service, *model.ErrorResBody) {
	service.Uuid, _ = uuid.NewV4()
	key, _ := uuid.NewV4()
	service.ApiKey = strings.Replace(key.String(), "-", "", -1)
	return ss.serviceRepository.Save(*service)
}
