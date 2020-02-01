package entity

import (
	"time"
)

const (
	GroupPermissionTable GroupPermissionTableConfig = iota
	GroupPermissionId
	GroupPermissionPermissionId
	GroupPermissionGroupId
	GroupPermissionCreatedAt
	GroupPermissionUpdatedAt
)

// The table `group_permissions` struct
type GroupPermission struct {
	Id           int       `json:"id"`
	PermissionId int       `validate:"required"json:"permission_id"`
	GroupId      int       `validate:"required"json:"group_id"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

// GroupPermission table config
type GroupPermissionTableConfig int

func (gp GroupPermissionTableConfig) String() string {
	switch gp {
	case GroupPermissionTable:
		return "group_permissions"
	case GroupPermissionId:
		return "id"
	case GroupPermissionPermissionId:
		return "permission_id"
	case GroupPermissionGroupId:
		return "group_id"
	case GroupPermissionCreatedAt:
		return "created_at"
	case GroupPermissionUpdatedAt:
		return "updated_at"
	}
	panic("Unknown value")
}
