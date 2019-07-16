package entity

import (
	"time"

	"github.com/satori/go.uuid"
)

const (
	PermissionId PermissionColumn = iota
	PermissionUuid
	PermissionName
	PermissionCreatedAt
	PermissionUpdatedAt
)

type Permission struct {
	Id           int       `json:"id"`
	Uuid         uuid.UUID `json:"uuid"`
	Name         string    `validate:"required"json:"name"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type PermissionColumn int

func (pc PermissionColumn) String() string {
	switch pc {
	case PermissionId:
		return "id"
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

func (p Permission) TableName() string {
	return "permissions"
}
