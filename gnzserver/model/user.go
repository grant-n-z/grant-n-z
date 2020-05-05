package model

import (
	"github.com/google/uuid"

	"github.com/tomoyane/grant-n-z/gnz/entity"
)

// User data in jwt
type AuthUser struct {
	UserUuid  uuid.UUID `json:"user_uuid"`
	UserId    int       `json:"user_id"`
	UserName  string    `json:"user_name"`
	ServiceId int       `json:"service_id"`
	Expires   string    `json:"expires"`
	IssueDate string    `json:"issue_date"`
	RoleId    int       `json:"role_id"`
	PolicyId  int       `json:"policy_id"`
}

// The table `users` and `operator_policies` struct
type UserWithOperatorPolicy struct {
	entity.User
	entity.OperatorPolicy
}

// The table `users` and `user_services` and `services` struct
type UserWithUserServiceWithService struct {
	entity.User
	entity.UserService
	entity.Service
}

// Add user id
type AddUser struct {
	UserEmail string `json:"user_email"`
}
