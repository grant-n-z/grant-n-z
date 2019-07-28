package service

import (
	"github.com/tomoyane/grant-n-z/server/entity"
	"github.com/tomoyane/grant-n-z/server/model"
)

type OperatorMemberRoleService interface {
	Get(queryParam string) ([]*entity.OperatorMemberRole, *model.ErrorResponse)

	GetAll() ([]*entity.OperatorMemberRole, *model.ErrorResponse)

	GetByUserId(userId int) ([]*entity.OperatorMemberRole, *model.ErrorResponse)

	Insert(roleMember *entity.OperatorMemberRole) (*entity.OperatorMemberRole, *model.ErrorResponse)
}
