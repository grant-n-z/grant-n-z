package entity

import (
	"time"
	"github.com/satori/go.uuid"
)

type Role struct {
	Id        int       `gorm:"primary_key"json:"id"`
	Type      string    `gorm:"type:varchar(128)"json:"type"`
	UserUuid  uuid.UUID `gorm:"type:varchar(128);index:user_uuid"json:"user_uuid"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (t Role) TableName() string {
	return "roles"
}
