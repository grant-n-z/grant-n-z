package entity

import (
	"github.com/google/uuid"
	"time"
)

const (
	OperatorPolicyTable OperatorPolicyTableConfig = iota
	OperatorPolicyId
	OperatorPolicyRoleUuid
	OperatorPolicyUserUuid
	OperatorPolicyCreatedAt
	OperatorPolicyUpdatedAt
)

// The table `operator_policies` struct
type OperatorPolicy struct {
	Id        int       `json:"id"`
	RoleUuid  uuid.UUID `validate:"required"json:"role_uuid"`
	UserUuid  uuid.UUID `validate:"required"json:"user_uuid"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// OperatorPolicy table config struct
type OperatorPolicyTableConfig int

func (opc OperatorPolicyTableConfig) String() string {
	switch opc {
	case OperatorPolicyTable:
		return "operator_policies"
	case OperatorPolicyId:
		return "id"
	case OperatorPolicyRoleUuid:
		return "role_uuid"
	case OperatorPolicyUserUuid:
		return "user_uuid"
	case OperatorPolicyCreatedAt:
		return "created_at"
	case OperatorPolicyUpdatedAt:
		return "updated_at"
	}
	panic("Unknown value")
}
