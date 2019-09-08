package entity

import (
	"time"
)

const (
	OperatorPolicyId OperatorPolicyColumn = iota
	OperatorPolicyRoleId
	OperatorPolicyUserId
	OperatorPolicyCreatedAt
	OperatorPolicyUpdatedAt
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
	case OperatorPolicyId:
		return "id"
	case OperatorPolicyRoleId:
		return "role_id"
	case OperatorPolicyUserId:
		return "user_id"
	case OperatorPolicyCreatedAt:
		return "created_at"
	case OperatorPolicyUpdatedAt:
		return "updated_at"
	}
	panic("Unknown value")
}

func (opc OperatorPolicyColumn) TableName() string {
	return "operator_policies"
}
