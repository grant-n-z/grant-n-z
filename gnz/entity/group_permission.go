package entity

import (
	"time"

	"github.com/google/uuid"
)

const (
	GroupPermissionTable GroupPermissionTableConfig = iota
	GroupPermissionId
	GroupPermissionPermissionUuid
	GroupPermissionGroupUuid
	GroupPermissionCreatedAt
	GroupPermissionUpdatedAt
)

// The table `group_permissions` struct
type GroupPermission struct {
	Id             int       `json:"id"`
	PermissionUuid uuid.UUID `validate:"required"json:"permission_uuid"`
	GroupUuid      uuid.UUID `validate:"required"json:"group_uuid"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

// GroupPermission table config
type GroupPermissionTableConfig int

func (gp GroupPermissionTableConfig) String() string {
	switch gp {
	case GroupPermissionTable:
		return "group_permissions"
	case GroupPermissionId:
		return "id"
	case GroupPermissionPermissionUuid:
		return "permission_uuid"
	case GroupPermissionGroupUuid:
		return "group_uuid"
	case GroupPermissionCreatedAt:
		return "created_at"
	case GroupPermissionUpdatedAt:
		return "updated_at"
	}
	panic("Unknown value")
}
