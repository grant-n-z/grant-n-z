package repository

import (
	"strings"

	"github.com/jinzhu/gorm"

	"github.com/tomoyane/grant-n-z/server/entity"
	"github.com/tomoyane/grant-n-z/server/log"
	"github.com/tomoyane/grant-n-z/server/model"
)

type ServiceRepositoryImpl struct {
	Db *gorm.DB
}

func (sri ServiceRepositoryImpl) FindAll() ([]*entity.Service, *model.ErrorResponse) {
	var services []*entity.Service
	if err := sri.Db.Find(&services).Error; err != nil {
		if strings.Contains(err.Error(), "record not found") {
			return nil, nil
		}

		return nil, model.InternalServerError(err.Error())
	}

	return services, nil
}

func (sri ServiceRepositoryImpl) FindById(id int) (*entity.Service, *model.ErrorResponse) {
	service := entity.Service{}
	if err := sri.Db.Where("id = ?", id).First(&service).Error; err != nil {
		if strings.Contains(err.Error(), "record not found") {
			return nil, nil
		}

		return nil, model.InternalServerError(err.Error())
	}

	return &service, nil
}

func (sri ServiceRepositoryImpl) FindByName(name string) (*entity.Service, *model.ErrorResponse) {
	service := entity.Service{}
	if err := sri.Db.Where("name = ?", name).First(&service).Error; err != nil {
		if strings.Contains(err.Error(), "record not found") {
			return nil, nil
		}

		return nil, model.InternalServerError(err.Error())
	}

	return &service, nil
}

func (sri ServiceRepositoryImpl) Save(service entity.Service) (*entity.Service, *model.ErrorResponse) {
	if err := sri.Db.Create(&service).Error; err != nil {
		errRes := model.Conflict(err.Error())
		if strings.Contains(err.Error(), "Duplicate entry") {
			log.Logger.Warn(errRes.ToJson(), errRes.Detail)
			return nil, model.Conflict(err.Error())
		}

		log.Logger.Error(errRes.ToJson(), errRes.Detail)
		return nil, model.InternalServerError(err.Error())
	}

	return &service, nil
}

func (sri ServiceRepositoryImpl) Update(service entity.Service) *entity.Service {
	if err := sri.Db.Update(&service).Error; err != nil {
		return nil
	}

	return &service
}
