package entity

import (
	"time"
	"github.com/satori/go.uuid"
)

type Token struct {
	Id           int       `gorm:"primary_key"json:"id"`
	TokenType    string    `json:"token_type"`
	Token        string    `json:"token"`
	RefreshToken string    `json:"refresh_token"`
	UserUuid     uuid.UUID `json:"user_uuid"`
	ExpiresAt    time.Time `json:"expires_at"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

func (t Token) TableName() string {
	return "tokens"
}
