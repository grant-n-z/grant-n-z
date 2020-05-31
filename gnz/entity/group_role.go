package entity

import (
	"time"

	"github.com/google/uuid"
)

const (
	GroupRoleTable GroupRoleTableConfig = iota
	GroupRoleId
	GroupRoleInternalId
	GroupRoleRoleUuid
	GroupRoleGroupUuid
	GroupRoleCreatedAt
	GroupRoleUpdatedAt
)

// The table `group_roles` struct
type GroupRole struct {
	Id         int       `json:"id"`
	InternalId string    `json:"internal_id"`
	RoleUuid   uuid.UUID `validate:"required"json:"role_uuid"`
	GroupUuid  uuid.UUID `validate:"required"json:"group_uuid"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

// GroupRole table config
type GroupRoleTableConfig int

func (gr GroupRoleTableConfig) String() string {
	switch gr {
	case GroupRoleTable:
		return "group_roles"
	case GroupRoleId:
		return "id"
	case GroupRoleInternalId:
		return "internal_id"
	case GroupRoleRoleUuid:
		return "role_uuid"
	case GroupRoleGroupUuid:
		return "group_uuid"
	case GroupRoleCreatedAt:
		return "created_at"
	case GroupRoleUpdatedAt:
		return "updated_at"
	}
	panic("Unknown value")
}
