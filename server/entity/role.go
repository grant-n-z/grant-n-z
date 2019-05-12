package entity

import (
	"github.com/satori/go.uuid"
	"time"
)

const (
	ROLE_ID RoleColumn = iota
	ROLE_UUID
	ROLE_NAME
	ROLE_CREATED_AT
	ROLE_UPDATED_AT
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
	case ROLE_ID:
		return "id"
	case ROLE_UUID:
		return "uuid"
	case ROLE_NAME:
		return "name"
	case ROLE_CREATED_AT:
		return "created_at"
	case ROLE_UPDATED_AT:
		return "updated_at"
	}
	panic("Unknown value")
}

func (r Role) TableName() string {
	return "roles"
}
