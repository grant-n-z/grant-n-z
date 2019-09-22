package repository

import (
	"github.com/tomoyane/grant-n-z/gserver/entity"
	"github.com/tomoyane/grant-n-z/gserver/model"
)

type ServiceRepository interface {
	FindAll() ([]*entity.Service, *model.ErrorResBody)

	FindById(id int) (*entity.Service, *model.ErrorResBody)

	FindByName(name string) (*entity.Service, *model.ErrorResBody)

	FindByApiKey(apiKey string) (*entity.Service, *model.ErrorResBody)

	Save(service entity.Service) (*entity.Service, *model.ErrorResBody)

	Update(service entity.Service) *entity.Service
}

