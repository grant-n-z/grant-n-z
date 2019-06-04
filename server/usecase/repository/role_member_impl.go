package repository

import (
	"strings"

	"github.com/jinzhu/gorm"

	"github.com/tomoyane/grant-n-z/server/entity"
	"github.com/tomoyane/grant-n-z/server/log"
	"github.com/tomoyane/grant-n-z/server/model"
)

type RoleMemberRepositoryImpl struct {
	Db *gorm.DB
}

func (rmri RoleMemberRepositoryImpl) FindAll() ([]*entity.RoleMember, *model.ErrorResponse) {
	var roleMembers []*entity.RoleMember
	if err := rmri.Db.Find(&roleMembers).Error; err != nil {
		if strings.Contains(err.Error(), "record not found") {
			return nil, nil
		}

		return nil, model.InternalServerError(err.Error())
	}

	return roleMembers, nil
}

func (rmri RoleMemberRepositoryImpl) FindByUserId(userId int) ([]*entity.RoleMember, *model.ErrorResponse) {
	var roleMembers []*entity.RoleMember
	if err := rmri.Db.Where("user_id = ?", userId).Find(&roleMembers).Error; err != nil {
		if strings.Contains(err.Error(), "record not found") {
			return nil, nil
		}

		return nil, model.InternalServerError(err.Error())
	}

	return roleMembers, nil
}

func (rmri RoleMemberRepositoryImpl) Save(roleMember entity.RoleMember) (*entity.RoleMember, *model.ErrorResponse) {
	if err := rmri.Db.Create(&roleMember).Error; err != nil {
		errRes := model.Conflict(err.Error())
		if strings.Contains(err.Error(), "Duplicate entry") {
			log.Logger.Warn(errRes.ToJson(), errRes.Detail)
			return nil, model.Conflict(err.Error())
		}

		log.Logger.Error(errRes.ToJson(), errRes.Detail)
		return nil, model.InternalServerError(err.Error())
	}

	return &roleMember, nil
}
