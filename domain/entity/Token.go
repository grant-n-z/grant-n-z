package entity

import (
	"github.com/satori/go.uuid"
	"time"
)

type Token struct {
	Id           int       `gorm:"primary_key"json:"id"`
	TokenType    string    `gorm:"type:varchar(128);not null"validate:"required"json:"token_type"`
	Token        string    `gorm:"type:varchar(512);not null"validate:"required"json:"token"`
	RefreshToken string    `gorm:"type:varchar(512);not null"validate:"required"json:"refresh_token"`
	UserUuid     uuid.UUID `gorm:"type:varchar(128);not null;index:user_uuid"validate:"required"json:"user_uuid"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

func (m Token) TableName() string {
	return "tokens"
}
