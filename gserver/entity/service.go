package entity

import (
	"time"

	"github.com/satori/go.uuid"
)

const (
	ServiceId ServiceColumn = iota
	ServiceUuid
	ServiceName
	ServiceCreatedAt
	ServiceUpdatedAt
)

type Service struct {
	Id        int       `gorm:"primary_key"json:"id"`
	Uuid      uuid.UUID `gorm:"type:varchar(128)"json:"uuid"`
	Name      string    `gorm:"unique;type:varchar(128)"validate:"required"json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ServiceColumn int

func (sc ServiceColumn) String() string {
	switch sc {
	case ServiceId:
		return "id"
	case ServiceUuid:
		return "uuid"
	case ServiceName:
		return "name"
	case ServiceCreatedAt:
		return "created_at"
	case ServiceUpdatedAt:
		return "updated_at"
	}
	panic("Unknown value")
}

func (s Service) TableName() string {
	return "services"
}
