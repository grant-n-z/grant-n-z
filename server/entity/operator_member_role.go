package entity

import (
	"time"
)

const (
	OperatorMemberRoleId OperatorMemberRoleColumn = iota
	OperatorMemberRoleRoleId
	OperatorMemberRoleUserId
	OperatorMemberRoleCreatedAt
	OperatorMemberRoleUpdatedAt
)

type OperatorMemberRole struct {
	Id        int       `json:"id"`
	RoleId    int       `validate:"required"json:"role_id"`
	UserId    int       `validate:"required"json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type OperatorMemberRoleColumn int

func (omrc OperatorMemberRoleColumn) String() string {
	switch omrc {
	case OperatorMemberRoleId:
		return "id"
	case OperatorMemberRoleRoleId:
		return "role_id"
	case OperatorMemberRoleUserId:
		return "user_id"
	case OperatorMemberRoleCreatedAt:
		return "created_at"
	case OperatorMemberRoleUpdatedAt:
		return "updated_at"
	}
	panic("Unknown value")
}

func (omr OperatorMemberRole) TableName() string {
	return "operator_member_roles"
}
