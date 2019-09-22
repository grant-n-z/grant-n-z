package service

import (
	"github.com/tomoyane/grant-n-z/gserver/entity"
	"github.com/tomoyane/grant-n-z/gserver/model"
)

type OperatorPolicyService interface {
	Get(queryParam string) ([]*entity.OperatorPolicy, *model.ErrorResBody)

	GetAll() ([]*entity.OperatorPolicy, *model.ErrorResBody)

	GetByUserId(userId int) ([]*entity.OperatorPolicy, *model.ErrorResBody)

	GetByUserIdAndRoleId(userId int, roleId int) (*entity.OperatorPolicy, *model.ErrorResBody)

	GetRoleNameByUserId(userId int) ([]string, *model.ErrorResBody)

	Insert(roleMember *entity.OperatorPolicy) (*entity.OperatorPolicy, *model.ErrorResBody)
}
