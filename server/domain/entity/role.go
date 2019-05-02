package entity

import (
	"github.com/satori/go.uuid"
	"time"
)

type Role struct {
	Id        int       `json:"id"`
	Uuid      uuid.UUID `json:"uuid"`
	Name      string    `validate:"required"json:"name"`
	ServiceId int       `validate:"required"json:"service_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (t Role) TableName() string {
	return "roles"
}
