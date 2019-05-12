package repository

import "github.com/tomoyane/grant-n-z/server/entity"

type ServiceRepository interface {
	FindAll() ([]*entity.Service, *entity.ErrorResponse)

	FindById(id int) (*entity.Service, *entity.ErrorResponse)

	FindByName(name string) (*entity.Service, *entity.ErrorResponse)

	Save(service entity.Service) (*entity.Service, *entity.ErrorResponse)

	Update(service entity.Service) *entity.Service
}

