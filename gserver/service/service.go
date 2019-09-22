package service

import (
	"github.com/tomoyane/grant-n-z/gserver/entity"
	"github.com/tomoyane/grant-n-z/gserver/model"
)

type Service interface {
	Get(queryParam string) (interface{}, *model.ErrorResBody)

	GetServices() ([]*entity.Service, *model.ErrorResBody)

	GetServiceById(id int) (*entity.Service, *model.ErrorResBody)

	GetServiceByName(name string) (*entity.Service, *model.ErrorResBody)

	GetServiceByApiKey(apiKey string) (*entity.Service, *model.ErrorResBody)

	InsertService(service *entity.Service) (*entity.Service, *model.ErrorResBody)
}
