package entity

import (
	"time"

	"github.com/google/uuid"
)

const (
	PermissionTable PermissionTableConfig = iota
	PermissionId
	PermissionInternalId
	PermissionUuid
	PermissionName
	PermissionCreatedAt
	PermissionUpdatedAt
)

// The table `permissions` struct
type Permission struct {
	Id         int       `json:"id"`
	InternalId string    `json:"internal_id"`
	Uuid       uuid.UUID `json:"uuid"`
	Name       string    `validate:"required"json:"name"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

// Permission table config struct
type PermissionTableConfig int

func (pc PermissionTableConfig) String() string {
	switch pc {
	case PermissionTable:
		return "permissions"
	case PermissionId:
		return "id"
	case PermissionInternalId:
		return "internal_id"
	case PermissionName:
		return "name"
	case PermissionUuid:
		return "uuid"
	case PermissionCreatedAt:
		return "created_at"
	case PermissionUpdatedAt:
		return "updated_at"
	}
	panic("Unknown value")
}
