package entity

import (
	"time"
)

const (
	PolicyId PolicyColumn = iota
	PolicyName
	PolicyPermissionId
	PolicyRoleId
	PolicyCreatedAt
	PolicyUpdatedAt
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
	case PolicyId:
		return "id"
	case PolicyName:
		return "name"
	case PolicyPermissionId:
		return "permission_id"
	case PolicyRoleId:
		return "role_id"
	case PolicyCreatedAt:
		return "created_at"
	case PolicyUpdatedAt:
		return "updated_at"
	}
	panic("Unknown value")
}

func (p Policy) TableName() string {
	return "policies"
}
