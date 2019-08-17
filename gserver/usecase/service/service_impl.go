package service

import (
	"strings"

	"github.com/satori/go.uuid"

	"github.com/tomoyane/grant-n-z/gserver/common/driver"
	"github.com/tomoyane/grant-n-z/gserver/entity"
	"github.com/tomoyane/grant-n-z/gserver/log"
	"github.com/tomoyane/grant-n-z/gserver/model"
	"github.com/tomoyane/grant-n-z/gserver/usecase/repository"
)

type serviceImpl struct {
	serviceRepository repository.ServiceRepository
}

func NewServiceService() Service {
	log.Logger.Info("Inject `serviceRepository` to `Service`")
	return serviceImpl{serviceRepository: repository.ServiceRepositoryImpl{Db: driver.Db}}
}

func (ss serviceImpl) Get(queryParam string) (interface{}, *model.ErrorResponse)  {
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

func (ss serviceImpl) GetServices() ([]*entity.Service, *model.ErrorResponse)  {
	return ss.serviceRepository.FindAll()
}

func (ss serviceImpl) GetServiceById(id int) (*entity.Service, *model.ErrorResponse)  {
	return ss.serviceRepository.FindById(id)
}

func (ss serviceImpl) GetServiceByName(name string) (*entity.Service, *model.ErrorResponse)  {
	return ss.serviceRepository.FindByName(name)
}

func (ss serviceImpl) InsertService(service *entity.Service) (*entity.Service, *model.ErrorResponse) {
	service.Uuid, _ = uuid.NewV4()
	return ss.serviceRepository.Save(*service)
}
