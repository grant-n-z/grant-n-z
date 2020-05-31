package entity

import (
	"time"

	"github.com/google/uuid"
)

const (
	RoleTable RoleTableConfig = iota
	RoleId
	RoleInternalId
	RoleUuid
	RoleName
	RoleCreatedAt
	RoleUpdatedAt
)

// The table `roles` struct
type Role struct {
	Id         int       `json:"id"`
	InternalId string    `json:"internal_id"`
	Uuid       uuid.UUID `json:"uuid"`
	Name       string    `validate:"required"json:"name"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

// Role table config struct
type RoleTableConfig int

func (rc RoleTableConfig) String() string {
	switch rc {
	case RoleTable:
		return "roles"
	case RoleId:
		return "id"
	case RoleInternalId:
		return "internal_id"
	case RoleUuid:
		return "uuid"
	case RoleName:
		return "name"
	case RoleCreatedAt:
		return "created_at"
	case RoleUpdatedAt:
		return "updated_at"
	}
	panic("Unknown value")
}
