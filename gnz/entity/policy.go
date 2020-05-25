package entity

import (
	"time"

	"github.com/google/uuid"
)

const (
	PolicyTable PolicyTableConfig = iota
	PolicyId
	PolicyName
	PolicyRoleUuid
	PolicyPermissionUuid
	PolicyServiceUuid
	PolicyUserGroupUuid
	PolicyCreatedAt
	PolicyUpdatedAt
)

// The table `policy` struct
type Policy struct {
	Id             int       `json:"id"`
	Name           string    `validate:"required"json:"name"`
	RoleUuid       uuid.UUID `validate:"required"json:"role_uuid"`
	PermissionUuid uuid.UUID `validate:"required"json:"permission_uuid"`
	ServiceUuid    uuid.UUID `validate:"required"json:"service_uuid"`
	UserGroupUuid  uuid.UUID `validate:"required"json:"user_group_uuid"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

// Policy table config struct
type PolicyTableConfig int

func (pc PolicyTableConfig) String() string {
	switch pc {
	case PolicyTable:
		return "policies"
	case PolicyId:
		return "id"
	case PolicyName:
		return "name"
	case PolicyRoleUuid:
		return "role_uuid"
	case PolicyPermissionUuid:
		return "permission_uuid"
	case PolicyServiceUuid:
		return "service_uuid"
	case PolicyUserGroupUuid:
		return "user_group_uuid"
	case PolicyCreatedAt:
		return "created_at"
	case PolicyUpdatedAt:
		return "updated_at"
	}
	panic("Unknown value")
}
