package repository

import (
	"github.com/tomoyane/grant-n-z/server/model"
	"strings"

	"github.com/jinzhu/gorm"

	"github.com/tomoyane/grant-n-z/server/entity"
	"github.com/tomoyane/grant-n-z/server/log"
)

type PermissionRepositoryImpl struct {
	Db *gorm.DB
}

func NewPermissionRepository(db *gorm.DB) PermissionRepository {
	return PermissionRepositoryImpl{
		Db: db,
	}
}

func (pri PermissionRepositoryImpl) FindAll() ([]*entity.Permission, *model.ErrorResponse) {
	var permissions []*entity.Permission
	if err := pri.Db.Find(&permissions).Error; err != nil {
		if strings.Contains(err.Error(), "record not found") {
			return nil, nil
		}

		return nil, model.InternalServerError(err.Error())
	}

	return permissions, nil
}

func (pri PermissionRepositoryImpl) FindById(id int) (*entity.Permission, *model.ErrorResponse) {
	permissions := entity.Permission{}
	if err := pri.Db.Where("id = ?", id).Find(&permissions).Error; err != nil {
		if strings.Contains(err.Error(), "record not found") {
			return nil, nil
		}

		return nil, model.InternalServerError(err.Error())
	}

	return &permissions, nil
}

func (pri PermissionRepositoryImpl) Save(permission entity.Permission) (*entity.Permission, *model.ErrorResponse) {
	if err := pri.Db.Create(&permission).Error; err != nil {
		errRes := model.Conflict(err.Error())
		if strings.Contains(err.Error(), "Duplicate entry") {
			log.Logger.Warn(errRes.ToJson(), errRes.Detail)
			return nil, model.Conflict(err.Error())
		}

		log.Logger.Error(errRes.ToJson(), errRes.Detail)
		return nil, model.InternalServerError(err.Error())
	}

	return &permission, nil
}
