package entity

import (
	"github.com/google/uuid"
	"time"
)

const (
	ServicePermissionTable ServicePermissionTableConfig = iota
	ServicePermissionId
	ServicePermissionPermissionUuid
	ServicePermissionServiceUuid
	ServicePermissionCreatedAt
	ServicePermissionUpdatedAt
)

// The table `service_permissions` struct
type ServicePermission struct {
	Id             int       `json:"id"`
	PermissionUuid uuid.UUID `validate:"required"json:"permission_uuid"`
	ServiceUuid    uuid.UUID `validate:"required"json:"service_uuid"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

// ServicePermission table config
type ServicePermissionTableConfig int

func (ur ServicePermissionTableConfig) String() string {
	switch ur {
	case ServicePermissionTable:
		return "service_permissions"
	case ServicePermissionId:
		return "id"
	case ServicePermissionPermissionUuid:
		return "permission_uuid"
	case ServicePermissionServiceUuid:
		return "service_uuid"
	case ServicePermissionCreatedAt:
		return "created_at"
	case ServicePermissionUpdatedAt:
		return "updated_at"
	}
	panic("Unknown value")
}
