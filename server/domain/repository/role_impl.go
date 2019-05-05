package repository

import (
	"strings"

	"github.com/tomoyane/grant-n-z/server/config"
	"github.com/tomoyane/grant-n-z/server/domain/entity"
	"github.com/tomoyane/grant-n-z/server/log"
)

type RoleRepositoryImpl struct {
}

func (rri RoleRepositoryImpl) FindAll() ([]*entity.Role, *entity.ErrorResponse) {
	var roles []*entity.Role
	if err := config.Db.Find(&roles).Error; err != nil {
		if strings.Contains(err.Error(), "record not found") {
			return nil, nil
		}

		return nil, entity.InternalServerError(err.Error())
	}

	return roles, nil
}

func (rri RoleRepositoryImpl) FindById(id int) (*entity.Role, *entity.ErrorResponse) {
	var role *entity.Role
	if err := config.Db.Where("id = ?", id).Find(&role).Error; err != nil {
		if strings.Contains(err.Error(), "record not found") {
			return nil, nil
		}

		return nil, entity.InternalServerError(err.Error())
	}

	return role, nil
}

func (rri RoleRepositoryImpl) Save(role entity.Role) (*entity.Role, *entity.ErrorResponse) {
	if err := config.Db.Create(&role).Error; err != nil {
		errRes := entity.Conflict(err.Error())
		if strings.Contains(err.Error(), "Duplicate entry") {
			log.Logger.Warn(errRes.ToJson(), errRes.Detail)
			return nil, entity.Conflict(err.Error())
		}

		log.Logger.Error(errRes.ToJson(), errRes.Detail)
		return nil, entity.InternalServerError(err.Error())
	}

	return &role, nil
}
