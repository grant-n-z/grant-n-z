package entity

import (
	"time"
	"github.com/satori/go.uuid"
)

type Member struct {
	Id          int       `json:"id"`
	Uuid        uuid.UUID `gorm:"type:varchar(128);not null"json:"uuid"`
	ServiceUuid uuid.UUID `gorm:"type:varchar(128);not null;index:service_uuid"json:"service_uuid"`
	UserUuid    uuid.UUID `gorm:"type:varchar(128);not null;index:user_uuid"json:"user_uuid"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (m Member) TableName() string {
	return "members"
}
