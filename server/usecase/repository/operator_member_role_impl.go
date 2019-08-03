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
		log.Logger.Warn(err.Error())
		if strings.Contains(err.Error(), "1062") {
			return nil, model.Conflict("Already exit data.")
		} else if strings.Contains(err.Error(), "1452") {
			return nil, model.BadRequest("Not register relational id.")
		}

		return nil, model.InternalServerError("Error internal processing.")
	}

	return &entity, nil
}
