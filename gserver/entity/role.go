package entity

import (
	"time"

	"github.com/satori/go.uuid"
)

const (
	RoleTable RoleTableConfig = iota
	RoleId
	RoleUuid
	RoleName
	RoleCreatedAt
	RoleUpdatedAt
)

type Role struct {
	Id        int       `json:"id"`
	Uuid      uuid.UUID `json:"uuid"`
	Name      string    `validate:"required"json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type RoleTableConfig int

func (rc RoleTableConfig) String() string {
	switch rc {
	case RoleTable:
		return "roles"
	case RoleId:
		return "id"
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
