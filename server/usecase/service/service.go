package service

import (
	"github.com/tomoyane/grant-n-z/server/entity"
)

type Service interface {
	Get(queryParam string) (interface{}, *entity.ErrorResponse)

	GetServices() ([]*entity.Service, *entity.ErrorResponse)

	GetServiceById(id int) (*entity.Service, *entity.ErrorResponse)

	GetServiceByName(name string) (*entity.Service, *entity.ErrorResponse)

	InsertService(service *entity.Service) (*entity.Service, *entity.ErrorResponse)
}
