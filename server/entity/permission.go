package entity

import (
	"time"

	"github.com/satori/go.uuid"
)

const (
	PERMISSION_ID PermissionColumn = iota
	PERMISSION_UUID
	PERMISSION_NAME
	PERMISSION_CREATED_AT
	PERMISSION_UPDATED_AT
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
	case PERMISSION_ID:
		return "id"
	case PERMISSION_NAME:
		return "name"
	case PERMISSION_UUID:
		return "uuid"
	case PERMISSION_CREATED_AT:
		return "created_at"
	case PERMISSION_UPDATED_AT:
		return "updated_at"
	}
	panic("Unknown value")
}

func (p Permission) TableName() string {
	return "permissions"
}
