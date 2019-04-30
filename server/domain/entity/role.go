package entity

import (
	"github.com/satori/go.uuid"
	"time"
)

type Role struct {
	Id        int       `gorm:"primary_key"json:"id"`
	Uuid      uuid.UUID `gorm:"type:varchar(128)"json:"uuid"`
	Name      string    `gorm:"unique;type:varchar(128)"validate:"required"json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (t Role) TableName() string {
	return "roles"
}
