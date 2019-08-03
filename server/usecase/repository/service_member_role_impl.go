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
		errRes := model.Conflict(err.Error())
		if strings.Contains(err.Error(), "Duplicate entry") {
			log.Logger.Warn(errRes.ToJson(), errRes.Detail)
			return nil, model.Conflict(err.Error())
		}

		log.Logger.Error(errRes.ToJson(), errRes.Detail)
		return nil, model.InternalServerError(err.Error())
	}

	return &entity, nil
}
