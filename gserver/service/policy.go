package service

import (
	"github.com/tomoyane/grant-n-z/gserver/entity"
	"github.com/tomoyane/grant-n-z/gserver/model"
)

type PolicyService interface {
	Get(queryParam string) ([]*entity.Policy, *model.ErrorResBody)

	GetPolicies() ([]*entity.Policy, *model.ErrorResBody)

	GetPoliciesByRoleId(roleId int) ([]*entity.Policy, *model.ErrorResBody)

	InsertPolicy(policy *entity.Policy) (*entity.Policy, *model.ErrorResBody)

	EncryptData(data string) (*string, error)

	DecryptData(data string) (*string, error)
}
