package repository

import (
	"strings"

	"github.com/tomoyane/grant-n-z/server/config"
	"github.com/tomoyane/grant-n-z/server/domain/entity"
	"github.com/tomoyane/grant-n-z/server/log"
)

type RoleMemberRepositoryImpl struct {
}

func (rmri RoleMemberRepositoryImpl) FindAll() ([]*entity.RoleMember, *entity.ErrorResponse) {
	var roleMembers []*entity.RoleMember
	if err := config.Db.Find(&roleMembers).Error; err != nil {
		if strings.Contains(err.Error(), "record not found") {
			return nil, nil
		}

		return nil, entity.InternalServerError(err.Error())
	}

	return roleMembers, nil
}

func (rmri RoleMemberRepositoryImpl) FindByUserId(userId int) ([]*entity.RoleMember, *entity.ErrorResponse) {
	var roleMembers []*entity.RoleMember
	if err := config.Db.Where("user_id = ?", userId).Find(&roleMembers).Error; err != nil {
		if strings.Contains(err.Error(), "record not found") {
			return nil, nil
		}

		return nil, entity.InternalServerError(err.Error())
	}

	return roleMembers, nil
}

func (rmri RoleMemberRepositoryImpl) Save(roleMember entity.RoleMember) (*entity.RoleMember, *entity.ErrorResponse) {
	if err := config.Db.Create(&roleMember).Error; err != nil {
		errRes := entity.Conflict(err.Error())
		if strings.Contains(err.Error(), "Duplicate entry") {
			log.Logger.Warn(errRes.ToJson(), errRes.Detail)
			return nil, entity.Conflict(err.Error())
		}

		log.Logger.Error(errRes.ToJson(), errRes.Detail)
		return nil, entity.InternalServerError(err.Error())
	}

	return &roleMember, nil
}
