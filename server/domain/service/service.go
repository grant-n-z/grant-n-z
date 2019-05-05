package service

import (
	"github.com/satori/go.uuid"
	"github.com/tomoyane/grant-n-z/server/domain/entity"
	"github.com/tomoyane/grant-n-z/server/domain/repository"
)

type Service struct {
	ServiceRepository repository.ServiceRepository
}

func NewServiceService() Service {
	return Service{ServiceRepository: repository.ServiceRepositoryImpl{}}
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
