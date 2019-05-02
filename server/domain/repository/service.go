package repository

import "github.com/tomoyane/grant-n-z/server/domain/entity"

type ServiceRepository interface {
	FindById(id int) (*entity.Service, *entity.ErrorResponse)

	Save(service entity.Service) (*entity.Service, *entity.ErrorResponse)

	Update(service entity.Service) *entity.Service
}

