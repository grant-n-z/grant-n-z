package entity

import (
	"time"

	"github.com/google/uuid"
)

const (
	ServiceTable ServiceTableConfig = iota
	ServiceId
	ServiceInternalId
	ServiceUuid
	ServiceName
	ServiceSecret
	ServiceCreatedAt
	ServiceUpdatedAt
)

// The table `services` struct
type Service struct {
	Id         int       `json:"id"`
	InternalId string    `json:"internal_id"`
	Uuid       uuid.UUID `json:"uuid"`
	Name       string    `validate:"required"json:"name"`
	Secret     string    `json:"secret"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

// Service table config struct
type ServiceTableConfig int

func (sc ServiceTableConfig) String() string {
	switch sc {
	case ServiceTable:
		return "services"
	case ServiceId:
		return "id"
	case ServiceInternalId:
		return "internal_id"
	case ServiceUuid:
		return "uuid"
	case ServiceName:
		return "name"
	case ServiceSecret:
		return "secret"
	case ServiceCreatedAt:
		return "created_at"
	case ServiceUpdatedAt:
		return "updated_at"
	}
	panic("Unknown value")
}
