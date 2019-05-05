package service

import (
	"github.com/satori/go.uuid"
	"github.com/tomoyane/grant-n-z/server/domain/entity"
	"github.com/tomoyane/grant-n-z/server/domain/repository"
	"github.com/tomoyane/grant-n-z/server/log"
	"strings"
)

type Service struct {
	ServiceRepository repository.ServiceRepository
}

func NewServiceService() Service {
	log.Logger.Info("inject `ServiceRepository` to `Service`")
	return Service{ServiceRepository: repository.ServiceRepositoryImpl{}}
}

func (ss Service) Get(queryParam string) (interface{}, *entity.ErrorResponse)  {
	var result interface{}

	if !strings.EqualFold(queryParam, "") {
		serviceEntity, err := ss.GetServiceByName(queryParam)
		if err != nil {
			return nil, err
		}

		if serviceEntity == nil {
			result = map[string]string{}
		} else {
			result = serviceEntity
		}

	} else {
		serviceEntities, err := ss.GetServices()
		if err != nil {
			return nil, err
		}

		if serviceEntities == nil {
			result = []string{}
		} else {
			result = serviceEntities
		}
	}

	return result, nil
}

func (ss Service) GetServices() ([]*entity.Service, *entity.ErrorResponse)  {
	return ss.ServiceRepository.FindAll()
}

func (ss Service) GetServiceById(id int) (*entity.Service, *entity.ErrorResponse)  {
	return ss.ServiceRepository.FindById(id)
}

func (ss Service) GetServiceByName(name string) (*entity.Service, *entity.ErrorResponse)  {
	return ss.ServiceRepository.FindByName(name)
}

func (ss Service) InsertService(service *entity.Service) (*entity.Service, *entity.ErrorResponse) {
	service.Uuid, _ = uuid.NewV4()
	return ss.ServiceRepository.Save(*service)
}
