package entity

import (
	"github.com/satori/go.uuid"
	"time"
)

type Principal struct {
	Id         int       `gorm:"primary_key"json:"id"`
	Uuid       uuid.UUID `gorm:"type:varchar(128);not null"json:"uuid"`
	MemberUuid uuid.UUID `gorm:"type:varchar(128);not null;index:member_uuid"json:"member_uuid"`
	RoleUuid   uuid.UUID `gorm:"type:varchar(128);not null;index:role_uuid"json:"role_uuid"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type PrincipalRequest struct {
	UserName       string `validate:"required"json:"user_name"`
	ServiceName    string `validate:"required"json:"service_name"`
	RolePermission string `validate:"required"json:"role_permission"`
}

func (m Principal) TableName() string {
	return "principals"
}
