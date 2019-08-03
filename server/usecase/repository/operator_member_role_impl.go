package repository

import (
	"strings"

	"github.com/jinzhu/gorm"

	"github.com/tomoyane/grant-n-z/server/entity"
	"github.com/tomoyane/grant-n-z/server/log"
	"github.com/tomoyane/grant-n-z/server/model"
)

type OperatorMemberRoleRepositoryImpl struct {
	Db *gorm.DB
}

func NewOperatorMemberRoleRepository(db *gorm.DB) OperatorMemberRoleRepository {
	return OperatorMemberRoleRepositoryImpl{
		Db: db,
	}
}

func (omrri OperatorMemberRoleRepositoryImpl) FindAll() ([]*entity.OperatorMemberRole, *model.ErrorResponse) {
	var entities []*entity.OperatorMemberRole
	if err := omrri.Db.Find(&entities).Error; err != nil {
		if strings.Contains(err.Error(), "record not found") {
			return nil, nil
		}

		return nil, model.InternalServerError(err.Error())
	}

	return entities, nil
}

func (omrri OperatorMemberRoleRepositoryImpl) FindByUserId(userId int) ([]*entity.OperatorMemberRole, *model.ErrorResponse) {
	var entities []*entity.OperatorMemberRole
	if err := omrri.Db.Where("user_id = ?", userId).Find(&entities).Error; err != nil {
		if strings.Contains(err.Error(), "record not found") {
			return nil, nil
		}

		return nil, model.InternalServerError(err.Error())
	}

	return entities, nil
}

func (omrri OperatorMemberRoleRepositoryImpl) Save(entity entity.OperatorMemberRole) (*entity.OperatorMemberRole, *model.ErrorResponse) {
	if err := omrri.Db.Create(&entity).Error; err != nil {
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
