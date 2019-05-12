package repository

import (
	"strings"

	"github.com/jinzhu/gorm"

	"github.com/tomoyane/grant-n-z/server/entity"
	"github.com/tomoyane/grant-n-z/server/log"
)

type PermissionRepositoryImpl struct {
	Db *gorm.DB
}

func (pri PermissionRepositoryImpl) FindAll() ([]*entity.Permission, *entity.ErrorResponse) {
	var permissions []*entity.Permission
	if err := pri.Db.Find(&permissions).Error; err != nil {
		if strings.Contains(err.Error(), "record not found") {
			return nil, nil
		}

		return nil, entity.InternalServerError(err.Error())
	}

	return permissions, nil
}

func (pri PermissionRepositoryImpl) FindById(id int) (*entity.Permission, *entity.ErrorResponse) {
	permissions := entity.Permission{}
	if err := pri.Db.Where("id = ?", id).Find(&permissions).Error; err != nil {
		if strings.Contains(err.Error(), "record not found") {
			return nil, nil
		}

		return nil, entity.InternalServerError(err.Error())
	}

	return &permissions, nil
}

func (pri PermissionRepositoryImpl) Save(permission entity.Permission) (*entity.Permission, *entity.ErrorResponse) {
	if err := pri.Db.Create(&permission).Error; err != nil {
		errRes := entity.Conflict(err.Error())
		if strings.Contains(err.Error(), "Duplicate entry") {
			log.Logger.Warn(errRes.ToJson(), errRes.Detail)
			return nil, entity.Conflict(err.Error())
		}

		log.Logger.Error(errRes.ToJson(), errRes.Detail)
		return nil, entity.InternalServerError(err.Error())
	}

	return &permission, nil
}
