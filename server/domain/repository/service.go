package repository

import "github.com/tomoyane/grant-n-z/server/domain/entity"

type ServiceRepository interface {
	Save(service entity.Service) (*entity.Service, *entity.ErrorResponse)

	Update(service entity.Service) *entity.Service
}

