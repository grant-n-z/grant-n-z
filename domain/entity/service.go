package entity

import (
	"time"
	"github.com/satori/go.uuid"
)

type Service struct {
	Id        int       `gorm:"primary_key"json:"id"`
	Uuid      uuid.UUID `gorm:"type:varchar(128);not null"json:"uuid"`
	Name      string    `gorm:"type:varchar(128);not null"json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (m Service) TableName() string {
	return "services"
}
