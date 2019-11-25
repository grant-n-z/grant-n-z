package entity

import (
	"time"
)

const (
	OperatorPolicyTable OperatorPolicyTableConfig = iota
	OperatorPolicyId
	OperatorPolicyRoleId
	OperatorPolicyUserId
	OperatorPolicyCreatedAt
	OperatorPolicyUpdatedAt
)

// The table `operator_policies` struct
type OperatorPolicy struct {
	Id        int       `json:"id"`
	RoleId    int       `validate:"required"json:"role_id"`
	UserId    int       `validate:"required"json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	Role Role
	User User
}

// OperatorPolicy table config struct
type OperatorPolicyTableConfig int

func (opc OperatorPolicyTableConfig) String() string {
	switch opc {
	case OperatorPolicyTable:
		return "operator_policies"
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
