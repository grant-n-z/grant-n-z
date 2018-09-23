package entity

import (
	"time"
	"github.com/satori/go.uuid"
)

type Member struct {
	Id            int       `gorm:"primary_key"json:"id"`
	Uuid          uuid.UUID `gorm:"type:varchar(128);not null"json:"uuid"`
	Type          uuid.UUID `gorm:"type:varchar(128);not null"json:"type"`
	UserUuid      uuid.UUID `gorm:"type:varchar(128);not null;index:user_uuid"json:"user_uuid"`
	PrincipalUuid uuid.UUID `gorm:"type:varchar(128);not null;index:principal_uuid"json:"principal_uuid"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

func (m Member) TableName() string {
	return "members"
}
