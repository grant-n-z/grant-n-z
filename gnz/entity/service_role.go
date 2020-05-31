package entity

import (
	"github.com/google/uuid"
	"time"
)

const (
	ServiceRoleTable ServiceRoleTableConfig = iota
	ServiceRoleId
	ServiceRoleInternalId
	ServiceRoleRoleUuid
	ServiceRoleServiceUuid
	ServiceRoleCreatedAt
	ServiceRoleUpdatedAt
)

// The table `service_roles` struct
type ServiceRole struct {
	Id          int       `json:"id"`
	InternalId  string    `json:"internal_id"`
	RoleUuid    uuid.UUID `validate:"required"json:"role_uuid"`
	ServiceUuid uuid.UUID `validate:"required"json:"service_uuid"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// ServiceRole table config
type ServiceRoleTableConfig int

func (ur ServiceRoleTableConfig) String() string {
	switch ur {
	case ServiceRoleTable:
		return "service_roles"
	case ServiceRoleId:
		return "id"
	case ServiceRoleInternalId:
		return "internal_id"
	case ServiceRoleRoleUuid:
		return "role_uuid"
	case ServiceRoleServiceUuid:
		return "service_uuid"
	case ServiceRoleCreatedAt:
		return "created_at"
	case ServiceRoleUpdatedAt:
		return "updated_at"
	}
	panic("Unknown value")
}
