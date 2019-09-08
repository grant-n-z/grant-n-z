package service

import (
	"github.com/tomoyane/grant-n-z/gserver/entity"
	"github.com/tomoyane/grant-n-z/gserver/model"
)

type OperatorPolicyService interface {
	Get(queryParam string) ([]*entity.OperatorPolicy, *model.ErrorResponse)

	GetAll() ([]*entity.OperatorPolicy, *model.ErrorResponse)

	GetByUserId(userId int) ([]*entity.OperatorPolicy, *model.ErrorResponse)

	GetByUserIdAndRoleId(userId int, roleId int) (*entity.OperatorPolicy, *model.ErrorResponse)

	GetRoleNameByUserId(userId int) ([]string, *model.ErrorResponse)

	Insert(roleMember *entity.OperatorPolicy) (*entity.OperatorPolicy, *model.ErrorResponse)
}
