package repository

import (
	"github.com/tomoyane/grant-n-z/domain/entity"
	"github.com/tomoyane/grant-n-z/infra"
)

type ServiceRepositoryImpl struct {
}

func (s ServiceRepositoryImpl) Save(service *entity.Service) *entity.Service {
	if err := infra.Db.Create(&service).Error; err != nil {
		return nil
	}

	return &service
}

func (s ServiceRepositoryImpl) FindAll() []*entity.Service {
	var services []*entity.Service

	if err := infra.Db.Find(&services).Error; err != nil {
		if err.Error() == "record not found" {
			return []*entity.Service{}
		}
		return nil
	}

	return services
}