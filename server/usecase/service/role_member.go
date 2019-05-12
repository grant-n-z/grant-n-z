package service

import (
	"github.com/tomoyane/grant-n-z/server/entity"
)

type RoleMemberService interface {
	Get(queryParam string) ([]*entity.RoleMember, *entity.ErrorResponse)

	GetRoleMembers() ([]*entity.RoleMember, *entity.ErrorResponse)

	GetRoleMemberByUserId(userId int) ([]*entity.RoleMember, *entity.ErrorResponse)

	InsertRoleMember(roleMember *entity.RoleMember) (*entity.RoleMember, *entity.ErrorResponse)
}
