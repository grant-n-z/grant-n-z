package entity

import (
	"time"
)

type Role struct {
	Id        int       `gorm:"primary_key"json:"id"`
	Name      string    `gorm:"type:varchar(128)"validate:"required"json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (t Role) TableName() string {
	return "roles"
}
