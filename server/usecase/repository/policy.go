package repository

import (
	"github.com/tomoyane/grant-n-z/server/entity"
	"github.com/tomoyane/grant-n-z/server/model"
)

type PolicyRepository interface {
	FindAll() ([]*entity.Policy, *model.ErrorResponse)

	FindByRoleId(roleId int) ([]*entity.Policy, *model.ErrorResponse)

	Save(policy entity.Policy) (*entity.Policy, *model.ErrorResponse)
}
