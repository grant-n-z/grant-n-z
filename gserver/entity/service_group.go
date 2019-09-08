package entity

import (
	"time"
)

const (
	ServiceGroupTAble ServiceGroupTableConfig = iota
	ServiceGroupId
	ServiceGroupGroupId
	ServiceGroupServiceId
	ServiceGroupCreatedAt
	ServiceGroupUpdatedAt
)

type ServiceGroup struct {
	Id        int       `json:"id"`
	GroupId   int       `validate:"required"json:"group_id"`
	ServiceId int       `validate:"required"json:"service_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ServiceGroupTableConfig int

func (uc ServiceGroupTableConfig) String() string {
	switch uc {
	case ServiceGroupTAble:
		return "service_groups"
	case ServiceGroupId:
		return "id"
	case ServiceGroupGroupId:
		return "group_id"
	case ServiceGroupServiceId:
		return "service_id"
	case ServiceGroupCreatedAt:
		return "created_at"
	case ServiceGroupUpdatedAt:
		return "updated_at"
	}
	panic("Unknown value")
}
