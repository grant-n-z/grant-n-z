package entity

import (
	"github.com/satori/go.uuid"
	"time"
)

type Member struct {
	Id          int       `json:"id"`
	Uuid        uuid.UUID `gorm:"type:varchar(128);not null"validate:"required"json:"uuid"`
	ServiceUuid uuid.UUID `gorm:"type:varchar(128);not null;index:service_uuid"validate:"required"json:"service_uuid"`
	UserUuid    uuid.UUID `gorm:"type:varchar(128);not null;index:user_uuid"validate:"required"json:"user_uuid"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (m Member) TableName() string {
	return "members"
}
