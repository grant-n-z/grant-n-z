package entity

import (
	"time"
	"github.com/satori/go.uuid"
)

type Principal struct {
	Id         int       `gorm:"primary_key"json:"id"`
	MemberUuid uuid.UUID `gorm:"type:varchar(128);not null;index:member_uuid"json:"member_uuid"`
	RoleUuid   uuid.UUID `gorm:"type:varchar(128);not null;index:role_uuid"json:"role_uuid"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

func (m Principal) TableName() string {
	return "principals"
}
