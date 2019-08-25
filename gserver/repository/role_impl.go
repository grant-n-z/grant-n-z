package repository

import (
	"strings"

	"github.com/jinzhu/gorm"

	"github.com/tomoyane/grant-n-z/gserver/entity"
	"github.com/tomoyane/grant-n-z/gserver/log"
	"github.com/tomoyane/grant-n-z/gserver/model"
)

var rrInstance RoleRepository

type RoleRepositoryImpl struct {
	Db *gorm.DB
}

func GetRoleRepositoryInstance(db *gorm.DB) RoleRepository {
	if rrInstance == nil {
		rrInstance = NewRoleRepository(db)
	}
	return rrInstance
}

func NewRoleRepository(db *gorm.DB) RoleRepository {
	log.Logger.Info("New `RoleRepository` instance")
	log.Logger.Info("Inject `gorm.DB` to `RoleRepository`")
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
	var role entity.Role
	if err := rri.Db.Where("id = ?", id).Find(&role).Error; err != nil {
		if strings.Contains(err.Error(), "record not found") {
			return nil, nil
		}

		return nil, model.InternalServerError(err.Error())
	}

	return &role, nil
}

func (rri RoleRepositoryImpl) Save(role entity.Role) (*entity.Role, *model.ErrorResponse) {
	if err := rri.Db.Create(&role).Error; err != nil {
		log.Logger.Warn(err.Error())
		if strings.Contains(err.Error(), "1062") {
			return nil, model.Conflict("Already exit data.")
		}

		return nil, model.InternalServerError(err.Error())
	}

	return &role, nil
}
