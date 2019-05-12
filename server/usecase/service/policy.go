package service

import (
	"github.com/tomoyane/grant-n-z/server/entity"
)

type PolicyService interface {
	Get(queryParam string) ([]*entity.Policy, *entity.ErrorResponse)

	GetPolicies() ([]*entity.Policy, *entity.ErrorResponse)

	GetPoliciesByRoleId(roleId int) ([]*entity.Policy, *entity.ErrorResponse)

	InsertPolicy(policy *entity.Policy) (*entity.Policy, *entity.ErrorResponse)
}
