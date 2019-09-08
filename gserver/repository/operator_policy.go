package repository

import (
	"github.com/tomoyane/grant-n-z/gserver/entity"
	"github.com/tomoyane/grant-n-z/gserver/model"
)

type OperatorPolicyRepository interface {
	FindAll() ([]*entity.OperatorPolicy, *model.ErrorResponse)

	FindByUserId(userId int) ([]*entity.OperatorPolicy, *model.ErrorResponse)

	FindByUserIdAndRoleId(userId int, roleId int) (*entity.OperatorPolicy, *model.ErrorResponse)

	FindRoleNameByUserId(userId int) ([]string, *model.ErrorResponse)

	Save(role entity.OperatorPolicy) (*entity.OperatorPolicy, *model.ErrorResponse)
}
