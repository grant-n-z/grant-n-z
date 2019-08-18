package service

import (
	"github.com/tomoyane/grant-n-z/gserver/entity"
	"github.com/tomoyane/grant-n-z/gserver/model"
)

type OperatorMemberRoleService interface {
	Get(queryParam string) ([]*entity.OperatorMemberRole, *model.ErrorResponse)

	GetAll() ([]*entity.OperatorMemberRole, *model.ErrorResponse)

	GetByUserId(userId int) ([]*entity.OperatorMemberRole, *model.ErrorResponse)

	GetByUserIdAndRoleId(userId int, roleId int) (*entity.OperatorMemberRole, *model.ErrorResponse)

	GetRoleNameByUserId(userId int) ([]string, *model.ErrorResponse)

	Insert(roleMember *entity.OperatorMemberRole) (*entity.OperatorMemberRole, *model.ErrorResponse)
}
