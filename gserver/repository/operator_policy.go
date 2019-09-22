package repository

import (
	"github.com/tomoyane/grant-n-z/gserver/entity"
	"github.com/tomoyane/grant-n-z/gserver/model"
)

type OperatorPolicyRepository interface {
	FindAll() ([]*entity.OperatorPolicy, *model.ErrorResBody)

	FindByUserId(userId int) ([]*entity.OperatorPolicy, *model.ErrorResBody)

	FindByUserIdAndRoleId(userId int, roleId int) (*entity.OperatorPolicy, *model.ErrorResBody)

	FindRoleNameByUserId(userId int) ([]string, *model.ErrorResBody)

	Save(role entity.OperatorPolicy) (*entity.OperatorPolicy, *model.ErrorResBody)
}
