package entity

import (
	"time"
)

const (
	UserServiceId UserServiceColumn = iota
	UserServiceUserId
	UserServiceServiceId
	UserServiceCreatedAt
	UserServiceUpdatedAt
)

type UserService struct {
	Id        int       `json:"id"`
	UserId    int       `validate:"required"json:"user_id"`
	ServiceId int       `validate:"required"json:"service_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserServiceColumn int

func (usc UserServiceColumn) String() string {
	switch usc {
	case UserServiceId:
		return "id"
	case UserServiceUserId:
		return "user_id"
	case UserServiceServiceId:
		return "service_id"
	case UserServiceCreatedAt:
		return "created_at"
	case UserServiceUpdatedAt:
		return "updated_at"
	}
	panic("Unknown value")
}

func (usc UserServiceColumn) TableName() string {
	return "user_services"
}
