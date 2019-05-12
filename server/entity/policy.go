package entity

import (
	"time"
)

const (
	POLICY_ID PolicyColumn = iota
	POLICY_NAME
	POLICY_PERMISSION_ID
	POLICY_ROLE_ID
	POLICY_CREATED_AT
	POLICY_UPDATED_AT
)

type Policy struct {
	Id           int       `json:"id"`
	Name         string    `validate:"required"json:"name"`
	PermissionId int       `validate:"required"json:"permission_id"`
	RoleId       int       `validate:"required"json:"role_id"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type PolicyColumn int

func (pc PolicyColumn) String() string {
	switch pc {
	case POLICY_ID:
		return "id"
	case POLICY_NAME:
		return "name"
	case POLICY_PERMISSION_ID:
		return "permission_id"
	case POLICY_ROLE_ID:
		return "role_id"
	case POLICY_CREATED_AT:
		return "created_at"
	case POLICY_UPDATED_AT:
		return "updated_at"
	}
	panic("Unknown value")
}

func (p Policy) TableName() string {
	return "policies"
}
