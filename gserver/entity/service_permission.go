package entity

import (
	"time"
)

const (
	ServicePermissionTable ServicePermissionTableConfig = iota
	ServicePermissionId
	ServicePermissionPermissionId
	ServicePermissionServiceId
	ServicePermissionCreatedAt
	ServicePermissionUpdatedAt
)

// The table `service_permissions` struct
type ServicePermission struct {
	Id           int       `json:"id"`
	PermissionId int       `validate:"required"json:"permission_id"`
	ServiceId    int       `validate:"required"json:"service_id"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

// ServicePermission table config
type ServicePermissionTableConfig int

func (ur ServicePermissionTableConfig) String() string {
	switch ur {
	case ServicePermissionTable:
		return "service_permissions"
	case ServicePermissionId:
		return "id"
	case ServicePermissionPermissionId:
		return "permission_id"
	case ServicePermissionServiceId:
		return "service_id"
	case ServicePermissionCreatedAt:
		return "created_at"
	case ServicePermissionUpdatedAt:
		return "updated_at"
	}
	panic("Unknown value")
}
