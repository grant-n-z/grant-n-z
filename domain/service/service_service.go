package service

import (
	"github.com/tomoyane/grant-n-z/domain/repository"
	"github.com/tomoyane/grant-n-z/domain/entity"
	"github.com/satori/go.uuid"
)

type ServiceService struct {
	ServiceRepository repository.ServiceRepository
}

func (s ServiceService) InsertService(service *entity.Service) *entity.Service {
	service.Uuid, _ = uuid.NewV4()
	return s.ServiceRepository.Save(service)
}

func (s ServiceService) GetAll() []*entity.Service {
	return s.ServiceRepository.FindAll()
}
