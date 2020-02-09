package data

import (
	"strings"

	"github.com/jinzhu/gorm"

	"github.com/tomoyane/grant-n-z/gnz/log"
	"github.com/tomoyane/grant-n-z/gnzserver/entity"
	"github.com/tomoyane/grant-n-z/gnzserver/model"
)

var plrInstance PolicyRepository

type PolicyRepository interface {
	FindAll() ([]*entity.Policy, *model.ErrorResBody)

	FindByRoleId(roleId int) ([]*entity.Policy, *model.ErrorResBody)

	FindById(id int) (entity.Policy, *model.ErrorResBody)

	Update(policy entity.Policy) (*entity.Policy, *model.ErrorResBody)
}

type PolicyRepositoryImpl struct {
	Db *gorm.DB
}

func GetPolicyRepositoryInstance(db *gorm.DB) PolicyRepository {
	if plrInstance == nil {
		plrInstance = NewPolicyRepository(db)
	}
	return plrInstance
}

func NewPolicyRepository(db *gorm.DB) PolicyRepository {
	log.Logger.Info("New `PolicyRepository` instance")
	return PolicyRepositoryImpl{
		Db: db,
	}
}

func (pri PolicyRepositoryImpl) FindAll() ([]*entity.Policy, *model.ErrorResBody) {
	var policies []*entity.Policy
	if err := pri.Db.Find(&policies).Error; err != nil {
		if strings.Contains(err.Error(), "record not found") {
			return nil, nil
		}

		return nil, model.InternalServerError(err.Error())
	}

	return policies, nil
}

func (pri PolicyRepositoryImpl) FindByRoleId(roleId int) ([]*entity.Policy, *model.ErrorResBody) {
	var policies []*entity.Policy
	if err := pri.Db.Where("role_id = ?", roleId).Find(&policies).Error; err != nil {
		if strings.Contains(err.Error(), "record not found") {
			return nil, nil
		}

		return nil, model.InternalServerError(err.Error())
	}

	return policies, nil
}

func (pri PolicyRepositoryImpl) FindById(id int) (entity.Policy, *model.ErrorResBody) {
	var policy entity.Policy
	if err := pri.Db.Where("id = ?", id).Find(&policy).Error; err != nil {
		if strings.Contains(err.Error(), "record not found") {
			return policy, nil
		}

		return policy, model.InternalServerError(err.Error())
	}

	return policy, nil
}

func (pri PolicyRepositoryImpl) Update(policy entity.Policy) (*entity.Policy, *model.ErrorResBody) {
	if err := pri.Db.Where("user_group_id = ?", policy.UserGroupId).Assign(policy).FirstOrCreate(&policy).Error; err != nil {
		log.Logger.Warn(err.Error())
		if strings.Contains(err.Error(), "1062") {
			return nil, model.Conflict("Already exit data.")
		} else if strings.Contains(err.Error(), "1452") {
			return nil, model.BadRequest("Not register relational id.")
		}

		return nil, model.InternalServerError()
	}

	return &policy, nil
}
