package entity

import (
	"time"
)

const (
	ServiceRoleTable ServiceRoleTableConfig = iota
	ServiceRoleId
	ServiceRoleRoleId
	ServiceRoleServiceId
	ServiceRoleCreatedAt
	ServiceRoleUpdatedAt
)

type ServiceRole struct {
	Id        int       `json:"id"`
	RoleId    int       `validate:"required"json:"role_id"`
	ServiceId int       `validate:"required"json:"service_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ServiceRoleTableConfig int

func (ur ServiceRoleTableConfig) String() string {
	switch ur {
	case ServiceRoleTable:
		return "service_groups"
	case ServiceRoleId:
		return "id"
	case ServiceRoleRoleId:
		return "role_id"
	case ServiceRoleServiceId:
		return "service_id"
	case ServiceRoleCreatedAt:
		return "created_at"
	case ServiceRoleUpdatedAt:
		return "updated_at"
	}
	panic("Unknown value")
}
