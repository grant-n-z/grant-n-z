package repository

import (
	"strings"

	"github.com/jinzhu/gorm"

	"github.com/tomoyane/grant-n-z/server/entity"
	"github.com/tomoyane/grant-n-z/server/log"
	"github.com/tomoyane/grant-n-z/server/model"
)

type RoleRepositoryImpl struct {
	Db *gorm.DB
}

func NewRoleRepository(db *gorm.DB) RoleRepository {
	return RoleRepositoryImpl{
		Db: db,
	}
}

func (rri RoleRepositoryImpl) FindAll() ([]*entity.Role, *model.ErrorResponse) {
	var roles []*entity.Role
	if err := rri.Db.Find(&roles).Error; err != nil {
		if strings.Contains(err.Error(), "record not found") {
			return nil, nil
		}

		return nil, model.InternalServerError(err.Error())
	}

	return roles, nil
}

func (rri RoleRepositoryImpl) FindById(id int) (*entity.Role, *model.ErrorResponse) {
	role := entity.Role{}
	if err := rri.Db.Where("id = ?", id).Find(&role).Error; err != nil {
		if strings.Contains(err.Error(), "record not found") {
			return nil, nil
		}

		return nil, model.InternalServerError(err.Error())
	}

	return &role, nil
}

func (rri RoleRepositoryImpl) Save(role entity.Role) (*entity.Role, *model.ErrorResponse) {
	if err := rri.Db.Create(&role); err != nil {
		errRes := model.Conflict(err.Error.Error())
		if strings.Contains(err.Error.Error(), "Duplicate entry") {
			log.Logger.Warn(errRes.ToJson(), errRes.Detail)
			return nil, model.Conflict(err.Error.Error())
		}

		log.Logger.Error(errRes.ToJson(), errRes.Detail)
		return nil, model.InternalServerError(err.Error.Error())
	}

	return &role, nil
}
