package repository

import "github.com/tomoyane/grant-n-z/domain/entity"

type ServiceRepository interface {
	Save(service *entity.Service) *entity.Service

	FindAll() []*entity.Service

	FindByName(name string) *entity.Service
}
