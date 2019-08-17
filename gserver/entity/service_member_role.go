package entity

import (
	"time"
)

const (
	ServiceMemberRoleId ServiceMemberRoleColumn = iota
	ServiceMemberRoleRoleId
	ServiceMemberRoleUserServiceId
	ServiceMemberRoleCreatedAt
	ServiceMemberRoleUpdatedAt
)

type ServiceMemberRole struct {
	Id            int       `json:"id"`
	RoleId        int       `validate:"required"json:"role_id"`
	UserServiceId int       `validate:"required"json:"user_service_id"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type ServiceMemberRoleColumn int

func (smr ServiceMemberRoleColumn) String() string {
	switch smr {
	case ServiceMemberRoleId:
		return "id"
	case ServiceMemberRoleRoleId:
		return "role_id"
	case ServiceMemberRoleUserServiceId:
		return "user_service_id"
	case ServiceMemberRoleCreatedAt:
		return "created_at"
	case ServiceMemberRoleUpdatedAt:
		return "updated_at"
	}
	panic("Unknown value")
}

func (smr ServiceMemberRole) TableName() string {
	return "service_member_roles"
}
