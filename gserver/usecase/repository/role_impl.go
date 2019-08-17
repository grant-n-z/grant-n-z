package repository

import (
	"strings"

	"github.com/jinzhu/gorm"

	"github.com/tomoyane/grant-n-z/gserver/entity"
	"github.com/tomoyane/grant-n-z/gserver/log"
	"github.com/tomoyane/grant-n-z/gserver/model"
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
		log.Logger.Warn(err.Error.Error())
		if strings.Contains(err.Error.Error(), "1062") {
			return nil, model.Conflict("Already exit data.")
		}

		return nil, model.InternalServerError(err.Error.Error())
	}

	return &role, nil
}
