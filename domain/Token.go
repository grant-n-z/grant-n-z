package domain

import "time"

type Tokens struct {
	Id           int       `gorm:"primary_key"json:"id"`
	TokenType    string    `json:"token_type"`
	Token        string    `json:"token"`
	RefreshToken string    `json:"refresh_token"`
	UserId       int       `json:"user_id"`
	ExpiresAt    time.Time `json:"expires_at"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

func (Tokens) TableName() string {
	return "tokens"
}
