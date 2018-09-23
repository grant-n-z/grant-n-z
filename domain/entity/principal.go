package entity

import (
	"time"
	"github.com/satori/go.uuid"
)

type Principal struct {
	Id        int       `json:"id"`
	Name      string    `gorm:"type:varchar(128);not null"validate:"required"json:"name"`
	Uuid      uuid.UUID `gorm:"type:varchar(128);not null"json:"uuid"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (m Principal) TableName() string {
	return "principals"
}
