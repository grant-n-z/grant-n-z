package entity

import (
	"time"

	"github.com/satori/go.uuid"
)

type User struct {
	Id        int       `gorm:"primary_key"json:"id"`
	Uuid      uuid.UUID `gorm:"type:varchar(128)"json:"uuid"`
	Username  string    `gorm:"type:varchar(128)"validate:"required"json:"username"`
	Email     string    `gorm:"type:varchar(128);unique;index:email"validate:"required,email"json:"email"`
	Password  string    `gorm:"type:varchar(128)"validate:"min=8,required"json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (u User) TableName() string {
	return "users"
}
