package entity

import (
	"github.com/google/uuid"
	"time"
)

const (
	ServiceGroupTable ServiceGroupTableConfig = iota
	ServiceGroupId
	ServiceGroupGroupUuid
	ServiceGroupServiceUuid
	ServiceGroupCreatedAt
	ServiceGroupUpdatedAt
)

// The table `service_groups` struct
type ServiceGroup struct {
	Id          int       `json:"id"`
	GroupUuid   uuid.UUID `validate:"required"json:"group_uuid"`
	ServiceUuid uuid.UUID `validate:"required"json:"service_uuid"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// ServiceGroup table config
type ServiceGroupTableConfig int

func (uc ServiceGroupTableConfig) String() string {
	switch uc {
	case ServiceGroupTable:
		return "service_groups"
	case ServiceGroupId:
		return "id"
	case ServiceGroupGroupUuid:
		return "group_uuid"
	case ServiceGroupServiceUuid:
		return "service_uuid"
	case ServiceGroupCreatedAt:
		return "created_at"
	case ServiceGroupUpdatedAt:
		return "updated_at"
	}
	panic("Unknown value")
}
