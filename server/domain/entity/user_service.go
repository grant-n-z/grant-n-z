package entity

import (
	"time"
)

type UserService struct {
	Id        int       `gorm:"primary_key"json:"id"`
	UserId    int       `validate:"required"json:"user_id"`
	ServiceId int       `validate:"required"json:"service_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	User      User      `gorm:"foreignkey:UserId"`
	Service   Service   `gorm:"foreignkey:ServiceId"`
}

func (u UserService) TableName() string {
	return "user_services"
}
