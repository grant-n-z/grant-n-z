package repository

import "github.com/tomoyane/grant-n-z/server/entity"

type PolicyRepository interface {
	FindAll() ([]*entity.Policy, *entity.ErrorResponse)

	FindByRoleId(roleId int) ([]*entity.Policy, *entity.ErrorResponse)

	Save(policy entity.Policy) (*entity.Policy, *entity.ErrorResponse)
}
