package repository

import (
	"github.com/tomoyane/grant-n-z/gserver/entity"
	"github.com/tomoyane/grant-n-z/gserver/model"
)

type PolicyRepository interface {
	FindAll() ([]*entity.Policy, *model.ErrorResponse)

	FindByRoleId(roleId int) ([]*entity.Policy, *model.ErrorResponse)

	Save(policy entity.Policy) (*entity.Policy, *model.ErrorResponse)
}
