package repository

import (
	"github.com/tomoyane/grant-n-z/server/model"
	"strings"

	"github.com/jinzhu/gorm"

	"github.com/tomoyane/grant-n-z/server/entity"
	"github.com/tomoyane/grant-n-z/server/log"
)

type PolicyRepositoryImpl struct {
	Db *gorm.DB
}

func NewPolicyRepository(db *gorm.DB) PolicyRepository {
	return PolicyRepositoryImpl{
		Db: db,
	}
}

func (pri PolicyRepositoryImpl) FindAll() ([]*entity.Policy, *model.ErrorResponse) {
	var policies []*entity.Policy
	if err := pri.Db.Find(&policies).Error; err != nil {
		if strings.Contains(err.Error(), "record not found") {
			return nil, nil
		}

		return nil, model.InternalServerError(err.Error())
	}

	return policies, nil
}

func (pri PolicyRepositoryImpl) FindByRoleId(roleId int) ([]*entity.Policy, *model.ErrorResponse) {
	var policies []*entity.Policy
	if err := pri.Db.Where("role_id = ?", roleId).Find(&policies).Error; err != nil {
		if strings.Contains(err.Error(), "record not found") {
			return nil, nil
		}

		return nil, model.InternalServerError(err.Error())
	}

	return policies, nil
}

func (pri PolicyRepositoryImpl) Save(policy entity.Policy) (*entity.Policy, *model.ErrorResponse) {
	if err := pri.Db.Create(&policy).Error; err != nil {
		errRes := model.Conflict(err.Error())
		if strings.Contains(err.Error(), "Duplicate entry") {
			log.Logger.Warn(errRes.ToJson(), errRes.Detail)
			return nil, model.Conflict(err.Error())
		}

		log.Logger.Error(errRes.ToJson(), errRes.Detail)
		return nil, model.InternalServerError(err.Error())
	}

	return &policy, nil
}
