package entity

import (
	"time"
)

const (
	OperatorMemberRoleId OperatorPolicyColumn = iota
	OperatorMemberRoleRoleId
	OperatorMemberRoleUserId
	OperatorMemberRoleCreatedAt
	OperatorMemberRoleUpdatedAt
)

type OperatorPolicy struct {
	Id        int       `json:"id"`
	RoleId    int       `validate:"required"json:"role_id"`
	UserId    int       `validate:"required"json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type OperatorPolicyColumn int

func (opc OperatorPolicyColumn) String() string {
	switch opc {
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

func (opc OperatorPolicyColumn) TableName() string {
	return "operator_policies"
}
