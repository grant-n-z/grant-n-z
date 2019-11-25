package entity

import (
	"time"

	"github.com/satori/go.uuid"
)

const (
	UserTable UserTableConfig = iota
	UserId
	UserUuid
	UserUsername
	UserEmail
	UserPassword
	UserCreatedAt
	UserUpdatedAt
)

// The table `users` struct
type User struct {
	Id        int       `json:"id"`
	Uuid      uuid.UUID `json:"uuid"`
	Username  string    `validate:"required"json:"username"`
	Email     string    `validate:"required,email"json:"email"`
	Password  string    `validate:"min=8,required"json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// The table `users` and `operator_policies` struct
type UserWithOperatorPolicy struct {
	User
	OperatorPolicy
}

// The table `users` and `user_services` and `services` struct
type UserWithUserServiceWithService struct {
	User
	UserService
	Service
}

// User table config struct
type UserTableConfig int

func (uc UserTableConfig) String() string {
	switch uc {
	case UserTable:
		return "users"
	case UserId:
		return "id"
	case UserUuid:
		return "uuid"
	case UserUsername:
		return "username"
	case UserEmail:
		return "email"
	case UserPassword:
		return "password"
	case UserCreatedAt:
		return "created_at"
	case UserUpdatedAt:
		return "updated_at"
	}
	panic("Unknown value")
}
