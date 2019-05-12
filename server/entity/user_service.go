package entity

import (
	"time"
)

const (
	USER_SERVICE_ID UserServiceColumn = iota
	USER_SERVICE_USER_ID
	USER_SERVICE_SERVICE_ID
	USER_SERVICE_CREATED_AT
	USER_SERVICE_UPDATED_AT
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
	case USER_SERVICE_ID:
		return "id"
	case USER_SERVICE_USER_ID:
		return "user_id"
	case USER_SERVICE_SERVICE_ID:
		return "service_id"
	case USER_SERVICE_CREATED_AT:
		return "created_at"
	case USER_SERVICE_UPDATED_AT:
		return "updated_at"
	}
	panic("Unknown value")
}

func (u UserService) TableName() string {
	return "user_services"
}
