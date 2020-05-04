package entity

import (
	"time"

	"github.com/google/uuid"
)

const (
	ServiceTable ServiceTableConfig = iota
	ServiceId
	ServiceUuid
	ServiceName
	ServiceSecret
	ServiceCreatedAt
	ServiceUpdatedAt
)

// The table `services` struct
type Service struct {
	Id        int       `gorm:"primary_key"json:"id"`
	Uuid      uuid.UUID `gorm:"type:varchar(128)"json:"uuid"`
	Name      string    `gorm:"unique;type:varchar(128)"validate:"required"json:"name"`
	Secret    string    `gorm:"type:varchar(128)"json:"secret"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Service table config struct
type ServiceTableConfig int

func (sc ServiceTableConfig) String() string {
	switch sc {
	case ServiceTable:
		return "services"
	case ServiceId:
		return "id"
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
