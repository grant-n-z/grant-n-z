package repository

import (
	"github.com/tomoyane/grant-n-z/server/entity"
	"github.com/tomoyane/grant-n-z/server/model"
)

type OperatorMemberRoleRepository interface {
	FindAll() ([]*entity.OperatorMemberRole, *model.ErrorResponse)

	FindByUserId(userId int) ([]*entity.OperatorMemberRole, *model.ErrorResponse)

	Save(role entity.OperatorMemberRole) (*entity.OperatorMemberRole, *model.ErrorResponse)
}
