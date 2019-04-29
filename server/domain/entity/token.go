package entity

import (
	"time"
	"github.com/satori/go.uuid"
)

type Token struct {
	Id           int       `gorm:"primary_key"json:"id"`
	TokenType    string    `gorm:"type:varchar(128)"json:"token_type"`
	Token        string    `gorm:"type:varchar(512)"json:"token"`
	RefreshToken string    `gorm:"type:varchar(512)"json:"refresh_token"`
	UserUuid     uuid.UUID `gorm:"type:varchar(128);index:user_uuid"json:"user_uuid"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

func (t Token) TableName() string {
	return "tokens"
}
