package repository

import (
	"strings"

	"github.com/jinzhu/gorm"

	"github.com/tomoyane/grant-n-z/server/entity"
	"github.com/tomoyane/grant-n-z/server/log"
	"github.com/tomoyane/grant-n-z/server/model"
)

type ServiceMemberRoleRepositoryImpl struct {
	Db *gorm.DB
}

func NewServiceMemberRoleRepository(db *gorm.DB) ServiceMemberRoleRepository {
	return ServiceMemberRoleRepositoryImpl{
		Db: db,
	}
}

func (smrri ServiceMemberRoleRepositoryImpl) FindAll() ([]*entity.ServiceMemberRole, *model.ErrorResponse) {
	var entities []*entity.ServiceMemberRole
	if err := smrri.Db.Find(&entities).Error; err != nil {
		if strings.Contains(err.Error(), "record not found") {
			return nil, nil
		}

		return nil, model.InternalServerError(err.Error())
	}

	return entities, nil
}

func (smrri ServiceMemberRoleRepositoryImpl) FindById(id int) ([]*entity.ServiceMemberRole, *model.ErrorResponse) {
	var entities []*entity.ServiceMemberRole
	if err := smrri.Db.Where("id = ?", id).Find(&entities).Error; err != nil {
		if strings.Contains(err.Error(), "record not found") {
			return nil, nil
		}

		return nil, model.InternalServerError(err.Error())
	}

	return entities, nil
}

func (smrri ServiceMemberRoleRepositoryImpl) FindByRoleId(roleId int) ([]*entity.ServiceMemberRole, *model.ErrorResponse) {
	var entities []*entity.ServiceMemberRole
	if err := smrri.Db.Where("role_id = ?", roleId).Find(&entities).Error; err != nil {
		if strings.Contains(err.Error(), "record not found") {
			return nil, nil
		}

		return nil, model.InternalServerError(err.Error())
	}

	return entities, nil
}

func (smrri ServiceMemberRoleRepositoryImpl) FindByUserServiceId(userServiceId int) ([]*entity.ServiceMemberRole, *model.ErrorResponse) {
	var entities []*entity.ServiceMemberRole
	if err := smrri.Db.Where("user_service_id = ?", userServiceId).Find(&entities).Error; err != nil {
		if strings.Contains(err.Error(), "record not found") {
			return nil, nil
		}

		return nil, model.InternalServerError(err.Error())
	}

	return entities, nil
}

func (smrri ServiceMemberRoleRepositoryImpl) Save(entity entity.ServiceMemberRole) (*entity.ServiceMemberRole, *model.ErrorResponse) {
	if err := smrri.Db.Create(&entity).Error; err != nil {
		log.Logger.Warn(err.Error())
		if strings.Contains(err.Error(), "1062") {
			return nil, model.Conflict("Already exit data.")
		} else if strings.Contains(err.Error(), "1452") {
			return nil, model.BadRequest("Not register relational id.")
		}

		return nil, model.InternalServerError()
	}

	return &entity, nil
}
