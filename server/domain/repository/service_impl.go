package repository

import (
	"strings"

	"github.com/tomoyane/grant-n-z/server/config"
	"github.com/tomoyane/grant-n-z/server/domain/entity"
	"github.com/tomoyane/grant-n-z/server/log"
)

type ServiceRepositoryImpl struct {
}

func (sri ServiceRepositoryImpl) FindById(id int) (*entity.Service, *entity.ErrorResponse) {
	service := entity.Service{}
	if err := config.Db.Where("id = ?", id).First(&service).Error; err != nil {
		return nil, entity.InternalServerError(err.Error())
	}

	return &service, nil
}

func (sri ServiceRepositoryImpl) Save(service entity.Service) (*entity.Service, *entity.ErrorResponse) {
	if err := config.Db.Create(&service).Error; err != nil {
		errRes := entity.Conflict(err.Error())
		if strings.Contains(err.Error(), "Duplicate entry") {
			log.Logger.Warn(errRes.ToJson(), errRes.Detail)
			return nil, entity.Conflict(err.Error())
		}

		log.Logger.Error(errRes.ToJson(), errRes.Detail)
		return nil, entity.InternalServerError(err.Error())
	}

	return &service, nil
}

func (sri ServiceRepositoryImpl) Update(service entity.Service) *entity.Service {
	if err := config.Db.Update(&service).Error; err != nil {
		return nil
	}

	return &service
}
