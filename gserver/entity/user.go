package entity

import (
	"time"

	"github.com/google/uuid"
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

// User data in jwt
type AuthUser struct {
	UserUuid  uuid.UUID `json:"user_uuid"`
	UserId    int       `json:"user_id"`
	ServiceId int       `json:"service_id"`
	Expires   string    `json:"expires"`
	RoleId    int       `json:"role_id"`
	PolicyId  int       `json:"policy_id"`
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

// Add user id
type AddUser struct {
	UserEmail string `json:"user_email"`
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
