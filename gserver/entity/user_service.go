package entity

import (
	"time"
)

const (
	UserServiceTable UserServiceTableConfig = iota
	UserServiceId
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

type UserServiceTableConfig int

func (usc UserServiceTableConfig) String() string {
	switch usc {
	case UserServiceTable:
		return "user_services"
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
