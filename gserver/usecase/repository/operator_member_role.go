package repository

import (
	"github.com/tomoyane/grant-n-z/gserver/entity"
	"github.com/tomoyane/grant-n-z/gserver/model"
)

type OperatorMemberRoleRepository interface {
	FindAll() ([]*entity.OperatorMemberRole, *model.ErrorResponse)

	FindByUserId(userId int) ([]*entity.OperatorMemberRole, *model.ErrorResponse)

	FindByUserIdAndRoleId(userId int, roleId int) (*entity.OperatorMemberRole, *model.ErrorResponse)

	FindRoleNameByUserId(userId int) ([]string, *model.ErrorResponse)

	Save(role entity.OperatorMemberRole) (*entity.OperatorMemberRole, *model.ErrorResponse)
}
