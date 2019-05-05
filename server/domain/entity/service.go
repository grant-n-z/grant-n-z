package entity

import (
	"time"

	"github.com/satori/go.uuid"
)

const (
	SERVICE_ID ServiceColumn = iota
	SERVICE_UUID
	SERVICE_NAME
	SERVICE_CREATED_AT
	SERVICE_UPDATED_AT
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
	case SERVICE_ID:
		return "id"
	case SERVICE_UUID:
		return "uuid"
	case SERVICE_NAME:
		return "name"
	case SERVICE_CREATED_AT:
		return "created_at"
	case SERVICE_UPDATED_AT:
		return "updated_at"
	}
	panic("Unknown value")
}

func (s Service) TableName() string {
	return "services"
}
