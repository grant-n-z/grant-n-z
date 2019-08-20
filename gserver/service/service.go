package service

import (
	"github.com/tomoyane/grant-n-z/gserver/entity"
	"github.com/tomoyane/grant-n-z/gserver/model"
)

type Service interface {
	Get(queryParam string) (interface{}, *model.ErrorResponse)

	GetServices() ([]*entity.Service, *model.ErrorResponse)

	GetServiceById(id int) (*entity.Service, *model.ErrorResponse)

	GetServiceByName(name string) (*entity.Service, *model.ErrorResponse)

	InsertService(service *entity.Service) (*entity.Service, *model.ErrorResponse)
}
