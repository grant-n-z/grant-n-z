package entity

import (
	"github.com/satori/go.uuid"
	"time"
)

type Role struct {
	Id         int       `gorm:"primary_key"json:"id"`
	Uuid       uuid.UUID `gorm:"type:varchar(128);not null"validate:"required"json:"uuid"`
	Permission string    `gorm:"type:varchar(128);not null"validate:"required"json:"permission"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

func (m Role) TableName() string {
	return "roles"
}
