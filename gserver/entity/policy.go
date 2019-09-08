package entity

import (
	"time"
)

const (
	PolicyId PolicyColumn = iota
	PolicyName
	PolicyPermissionId
	PolicyServiceMemberRoleId
	PolicyCreatedAt
	PolicyUpdatedAt
)

type Policy struct {
	Id                  int       `json:"id"`
	Name                string    `validate:"required"json:"name"`
	PermissionId        int       `validate:"required"json:"permission_id"`
	ServiceMemberRoleId int       `validate:"required"json:"service_member_role_id"`
	CreatedAt           time.Time `json:"created_at"`
	UpdatedAt           time.Time `json:"updated_at"`
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
	case PolicyServiceMemberRoleId:
		return "service_member_role_id"
	case PolicyCreatedAt:
		return "created_at"
	case PolicyUpdatedAt:
		return "updated_at"
	}
	panic("Unknown value")
}

func (pc PolicyColumn) TableName() string {
	return "policies"
}
