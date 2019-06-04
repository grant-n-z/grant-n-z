package service

import (
	"github.com/tomoyane/grant-n-z/server/entity"
	"github.com/tomoyane/grant-n-z/server/model"
)

type RoleMemberService interface {
	Get(queryParam string) ([]*entity.RoleMember, *model.ErrorResponse)

	GetRoleMembers() ([]*entity.RoleMember, *model.ErrorResponse)

	GetRoleMemberByUserId(userId int) ([]*entity.RoleMember, *model.ErrorResponse)

	InsertRoleMember(roleMember *entity.RoleMember) (*entity.RoleMember, *model.ErrorResponse)
}
