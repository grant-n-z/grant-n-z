package entity

import (
	"time"

	"github.com/satori/go.uuid"
)

const (
	RoleId RoleColumn = iota
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

type RoleColumn int

func (rc RoleColumn) String() string {
	switch rc {
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

func (r Role) TableName() string {
	return "roles"
}
