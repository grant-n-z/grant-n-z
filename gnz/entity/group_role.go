package entity

import (
	"time"
)

const (
	GroupRoleTable GroupRoleTableConfig = iota
	GroupRoleId
	GroupRoleRoleId
	GroupRoleGroupId
	GroupRoleCreatedAt
	GroupRoleUpdatedAt
)

// The table `group_roles` struct
type GroupRole struct {
	Id        int       `json:"id"`
	RoleId    int       `validate:"required"json:"role_id"`
	GroupId   int       `validate:"required"json:"group_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// GroupRole table config
type GroupRoleTableConfig int

func (gr GroupRoleTableConfig) String() string {
	switch gr {
	case GroupRoleTable:
		return "group_roles"
	case GroupRoleId:
		return "id"
	case GroupRoleRoleId:
		return "role_id"
	case GroupRoleGroupId:
		return "group_id"
	case GroupRoleCreatedAt:
		return "created_at"
	case GroupRoleUpdatedAt:
		return "updated_at"
	}
	panic("Unknown value")
}
