package service

import (
	"strings"

	"github.com/satori/go.uuid"

	"github.com/tomoyane/grant-n-z/server/config"
	"github.com/tomoyane/grant-n-z/server/entity"
	"github.com/tomoyane/grant-n-z/server/log"
	"github.com/tomoyane/grant-n-z/server/usecase/repository"
)

type serviceImpl struct {
	serviceRepository repository.ServiceRepository
}

func NewServiceService() Service {
	log.Logger.Info("Inject `serviceRepository` to `Service`")
	return serviceImpl{serviceRepository: repository.ServiceRepositoryImpl{Db: config.Db}}
}

func (ss serviceImpl) Get(queryParam string) (interface{}, *entity.ErrorResponse)  {
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

func (ss serviceImpl) GetServices() ([]*entity.Service, *entity.ErrorResponse)  {
	return ss.serviceRepository.FindAll()
}

func (ss serviceImpl) GetServiceById(id int) (*entity.Service, *entity.ErrorResponse)  {
	return ss.serviceRepository.FindById(id)
}

func (ss serviceImpl) GetServiceByName(name string) (*entity.Service, *entity.ErrorResponse)  {
	return ss.serviceRepository.FindByName(name)
}

func (ss serviceImpl) InsertService(service *entity.Service) (*entity.Service, *entity.ErrorResponse) {
	service.Uuid, _ = uuid.NewV4()
	return ss.serviceRepository.Save(*service)
}
