package entity

import (
	"github.com/google/uuid"
	"time"
)

const (
	UserServiceTable UserServiceTableConfig = iota
	UserServiceId
	UserServiceInternalId
	UserServiceUserUuid
	UserServiceServiceUuid
	UserServiceCreatedAt
	UserServiceUpdatedAt
)

// The table `user_services` struct
type UserService struct {
	Id          int       `json:"id"`
	InternalId  string    `json:"internal_id"`
	UserUuid    uuid.UUID `validate:"required"json:"user_uuid"`
	ServiceUuid uuid.UUID `validate:"required"json:"service_uuid"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// UserService table config struct
type UserServiceTableConfig int

func (usc UserServiceTableConfig) String() string {
	switch usc {
	case UserServiceTable:
		return "user_services"
	case UserServiceId:
		return "id"
	case UserServiceInternalId:
		return "internal_id"
	case UserServiceUserUuid:
		return "user_uuid"
	case UserServiceServiceUuid:
		return "service_uuid"
	case UserServiceCreatedAt:
		return "created_at"
	case UserServiceUpdatedAt:
		return "updated_at"
	}
	panic("Unknown value")
}
